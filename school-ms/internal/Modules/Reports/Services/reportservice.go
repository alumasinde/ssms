package services

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type ReportService struct {
	db *sqlx.DB
}

func NewReportService(db *sqlx.DB) *ReportService {
	return &ReportService{db: db}
}

// ── Report Card ───────────────────────────────────────────────────────────────

type SubjectResult struct {
	SubjectName string  `db:"subject_name" json:"subject_name"`
	Marks       float64 `db:"marks"        json:"marks"`
	MaxMarks    float64 `db:"max_marks"    json:"max_marks"`
	Percentage  float64 `json:"percentage"`
	Grade       string  `db:"grade"        json:"grade"`
	Remarks     string  `db:"remarks"      json:"remarks"`
}

type ReportCard struct {
	StudentID     int64           `json:"student_id"`
	StudentName   string          `json:"student_name"`
	AdmissionNo   string          `json:"admission_no"`
	ClassName     string          `json:"class_name"`
	ExamName      string          `json:"exam_name"`
	TermName      string          `json:"term_name"`
	Results       []SubjectResult `json:"results"`
	TotalMarks    float64         `json:"total_marks"`
	TotalMax      float64         `json:"total_max"`
	Average       float64         `json:"average"`
	OverallGrade  string          `json:"overall_grade"`
	Position      int             `json:"position"`
	ClassSize     int             `json:"class_size"`
	AttendancePct float64         `json:"attendance_pct"`
}

func (s *ReportService) GetReportCard(studentID, examID int64) (*ReportCard, error) {
	if studentID == 0 {
		return nil, errors.New("student_id is required")
	}
	if examID == 0 {
		return nil, errors.New("exam_id is required")
	}
	card := &ReportCard{}

	// Student + class info — uses split name columns
	err := s.db.QueryRow(`
		SELECT st.id,
		       CONCAT(st.first_name,' ',COALESCE(NULLIF(st.middle_name,''),''),' ',st.last_name) AS student_name,
		       st.admission_no,
		       c.name AS class_name
		FROM students st JOIN classes c ON c.id=st.class_id
		WHERE st.id=?`, studentID).
		Scan(&card.StudentID, &card.StudentName, &card.AdmissionNo, &card.ClassName)
	if err != nil {
		return nil, err
	}

	s.db.QueryRow(`SELECT e.name, t.name FROM exams e JOIN terms t ON t.id=e.term_id WHERE e.id=?`, examID).
		Scan(&card.ExamName, &card.TermName)

	rows, err := s.db.Queryx(`
		SELECT sub.name AS subject_name, er.marks, er.max_marks,
		       COALESCE(er.grade,'')   AS grade,
		       COALESCE(er.remarks,'') AS remarks
		FROM exam_results er
		JOIN subjects sub ON sub.id=er.subject_id
		WHERE er.student_id=? AND er.exam_id=?
		ORDER BY sub.name`, studentID, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r SubjectResult
		if err := rows.StructScan(&r); err != nil {
			continue
		}
		if r.MaxMarks > 0 {
			r.Percentage = r.Marks / r.MaxMarks * 100
		}
		card.Results = append(card.Results, r)
		card.TotalMarks += r.Marks
		card.TotalMax += r.MaxMarks
	}
	if card.TotalMax > 0 {
		card.Average = card.TotalMarks / card.TotalMax * 100
	}

	s.db.QueryRow(`
		SELECT COUNT(*)+1 FROM (
			SELECT student_id, SUM(marks)/SUM(max_marks)*100 AS avg
			FROM exam_results WHERE exam_id=?
			GROUP BY student_id
			HAVING avg > (SELECT SUM(marks)/SUM(max_marks)*100 FROM exam_results WHERE exam_id=? AND student_id=?)
		) ranked`, examID, examID, studentID).Scan(&card.Position)

	s.db.QueryRow(`SELECT COUNT(DISTINCT student_id) FROM exam_results WHERE exam_id=?`, examID).Scan(&card.ClassSize)

	s.db.QueryRow(`
		SELECT COALESCE(ROUND(SUM(status='present')/COUNT(*)*100,2), 0)
		FROM attendance a JOIN exams e ON e.term_id=a.term_id
		WHERE a.student_id=? AND e.id=?`, studentID, examID).Scan(&card.AttendancePct)

	return card, nil
}

// ── Class Result Summary ──────────────────────────────────────────────────────

type ClassResultRow struct {
	StudentID   int64   `db:"student_id"   json:"student_id"`
	StudentName string  `db:"student_name" json:"student_name"`
	AdmissionNo string  `db:"admission_no" json:"admission_no"`
	TotalMarks  float64 `db:"total_marks"  json:"total_marks"`
	TotalMax    float64 `db:"total_max"    json:"total_max"`
	Average     float64 `db:"average"      json:"average"`
	Position    int     `json:"position"`
}

func (s *ReportService) GetClassResults(examID int64) ([]ClassResultRow, error) {
	if examID == 0 {
		return []ClassResultRow{}, nil // no exam selected yet — return empty, not error
	}
	var rows []ClassResultRow
	err := s.db.Select(&rows, `
		SELECT er.student_id,
		       CONCAT(st.first_name,' ',st.last_name) AS student_name,
		       st.admission_no,
		       SUM(er.marks)                                        AS total_marks,
		       SUM(er.max_marks)                                    AS total_max,
		       ROUND(SUM(er.marks)/NULLIF(SUM(er.max_marks),0)*100,2) AS average
		FROM exam_results er
		JOIN students st ON st.id=er.student_id
		WHERE er.exam_id=?
		GROUP BY er.student_id, st.first_name, st.last_name, st.admission_no
		ORDER BY average DESC`, examID)
	if err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].Position = i + 1
	}
	return rows, nil
}

// ── Fee Collection Report ─────────────────────────────────────────────────────

type FeeCollectionReport struct {
	SchoolID     int64   `json:"school_id"`
	TermID       int64   `json:"term_id"`
	TotalBilled  float64 `json:"total_billed"`
	TotalPaid    float64 `json:"total_paid"`
	Balance      float64 `json:"balance"`
	PaidCount    int     `json:"paid_count"`
	UnpaidCount  int     `json:"unpaid_count"`
	PartialCount int     `json:"partial_count"`
}

func (s *ReportService) GetFeeCollection(schoolID, termID int64) (*FeeCollectionReport, error) {
	report := &FeeCollectionReport{SchoolID: schoolID, TermID: termID}
	if termID == 0 {
		return report, nil
	}
	s.db.QueryRow(`
		SELECT COALESCE(SUM(fi.amount),0),
		       COALESCE(SUM(fp_sum.total_paid),0),
		       COUNT(CASE WHEN fi.status='paid'    THEN 1 END),
		       COUNT(CASE WHEN fi.status='unpaid'  THEN 1 END),
		       COUNT(CASE WHEN fi.status='partial' THEN 1 END)
		FROM fee_invoices fi
		JOIN students st ON st.school_id=? AND st.id=fi.student_id
		LEFT JOIN (
			SELECT invoice_id, SUM(amount_paid) AS total_paid
			FROM fee_payments GROUP BY invoice_id
		) fp_sum ON fp_sum.invoice_id=fi.id
		WHERE fi.term_id=?`, schoolID, termID).
		Scan(&report.TotalBilled, &report.TotalPaid,
			&report.PaidCount, &report.UnpaidCount, &report.PartialCount)
	report.Balance = report.TotalBilled - report.TotalPaid
	return report, nil
}

// ── Attendance Summary ────────────────────────────────────────────────────────

type AttendanceReport struct {
	TermID     int64   `db:"term_id"    json:"term_id"`
	ClassID    int64   `db:"class_id"   json:"class_id"`
	ClassName  string  `db:"class_name" json:"class_name"`
	TotalDays  int     `db:"total_days" json:"total_days"`
	AvgPresent float64 `db:"avg_present" json:"avg_present_pct"`
}

func (s *ReportService) GetAttendanceSummary(schoolID, termID int64) ([]AttendanceReport, error) {
	if termID == 0 {
		return []AttendanceReport{}, nil
	}
	var reports []AttendanceReport
	err := s.db.Select(&reports, `
		SELECT t.id AS term_id, c.id AS class_id, c.name AS class_name,
		       COUNT(DISTINCT a.date) AS total_days,
		       COALESCE(ROUND(AVG(CASE WHEN a.status='present' THEN 100.0 ELSE 0 END),2),0) AS avg_present
		FROM attendance a
		JOIN classes c ON c.id=a.class_id
		JOIN terms   t ON t.id=a.term_id
		WHERE c.school_id=? AND a.term_id=?
		GROUP BY t.id, c.id, c.name
		ORDER BY c.name`, schoolID, termID)
	return reports, err
}

// ── Subject Performance Report ────────────────────────────────────────────────

type SubjectPerformanceRow struct {
	SubjectID   int64   `db:"subject_id"   json:"subject_id"`
	SubjectName string  `db:"subject_name" json:"subject_name"`
	SubjectCode string  `db:"subject_code" json:"subject_code"`
	EntryCount  int     `db:"entry_count"  json:"entry_count"`
	AvgScore    float64 `db:"avg_score"    json:"avg_score"`
	MinScore    float64 `db:"min_score"    json:"min_score"`
	MaxScore    float64 `db:"max_score"    json:"max_score"`
	PassCount   int     `db:"pass_count"   json:"pass_count"`
}

func (s *ReportService) GetSubjectPerformance(examID int64) ([]SubjectPerformanceRow, error) {
	if examID == 0 {
		return []SubjectPerformanceRow{}, nil
	}
	var rows []SubjectPerformanceRow
	err := s.db.Select(&rows, `
		SELECT er.subject_id,
		       sub.name AS subject_name,
		       sub.code AS subject_code,
		       COUNT(*)                                           AS entry_count,
		       ROUND(AVG(er.marks/NULLIF(er.max_marks,0)*100),2) AS avg_score,
		       ROUND(MIN(er.marks/NULLIF(er.max_marks,0)*100),2) AS min_score,
		       ROUND(MAX(er.marks/NULLIF(er.max_marks,0)*100),2) AS max_score,
		       SUM(er.marks/NULLIF(er.max_marks,0)*100 >= 50)    AS pass_count
		FROM exam_results er
		JOIN subjects sub ON sub.id = er.subject_id
		WHERE er.exam_id=?
		GROUP BY er.subject_id, sub.name, sub.code
		ORDER BY avg_score DESC`, examID)
	return rows, err
}
