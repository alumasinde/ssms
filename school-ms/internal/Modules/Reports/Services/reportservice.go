package services

import (
	"github.com/jmoiron/sqlx"
)

// ReportService aggregates data from multiple modules for reporting
type ReportService struct {
	db *sqlx.DB
}

func NewReportService(db *sqlx.DB) *ReportService {
	return &ReportService{db: db}
}

// --- Report Card ---

type SubjectResult struct {
	SubjectName string  `db:"subject_name" json:"subject_name"`
	Marks       float64 `db:"marks" json:"marks"`
	MaxMarks    float64 `db:"max_marks" json:"max_marks"`
	Percentage  float64 `json:"percentage"`
	Grade       string  `db:"grade" json:"grade"`
	Remarks     string  `db:"remarks" json:"remarks"`
}

type ReportCard struct {
	StudentID      int64           `json:"student_id"`
	StudentName    string          `json:"student_name"`
	AdmissionNo    string          `json:"admission_no"`
	ClassName      string          `json:"class_name"`
	ExamName       string          `json:"exam_name"`
	TermName       string          `json:"term_name"`
	Results        []SubjectResult `json:"results"`
	TotalMarks     float64         `json:"total_marks"`
	TotalMax       float64         `json:"total_max"`
	Average        float64         `json:"average"`
	OverallGrade   string          `json:"overall_grade"`
	Position       int             `json:"position"`
	ClassSize      int             `json:"class_size"`
	AttendancePct  float64         `json:"attendance_pct"`
}

func (s *ReportService) GetReportCard(studentID, examID int64) (*ReportCard, error) {
	card := &ReportCard{}

	// Student + class info
	err := s.db.QueryRow(`
		SELECT st.id, st.name, st.admission_no, c.name AS class_name
		FROM students st JOIN classes c ON c.id=st.class_id
		WHERE st.id=?`, studentID).Scan(&card.StudentID, &card.StudentName, &card.AdmissionNo, &card.ClassName)
	if err != nil {
		return nil, err
	}

	// Exam + term name
	s.db.QueryRow(`
		SELECT e.name, t.name FROM exams e JOIN terms t ON t.id=e.term_id WHERE e.id=?`, examID).
		Scan(&card.ExamName, &card.TermName)

	// Subject results
	rows, err := s.db.Queryx(`
		SELECT sub.name AS subject_name, er.marks, er.max_marks, er.grade, er.remarks
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

	// Class position
	s.db.QueryRow(`
		SELECT COUNT(*)+1 FROM (
			SELECT student_id, SUM(marks)/SUM(max_marks)*100 AS avg
			FROM exam_results WHERE exam_id=?
			GROUP BY student_id
			HAVING avg > (SELECT SUM(marks)/SUM(max_marks)*100 FROM exam_results WHERE exam_id=? AND student_id=?)
		) ranked`, examID, examID, studentID).Scan(&card.Position)

	s.db.QueryRow(`SELECT COUNT(DISTINCT student_id) FROM exam_results WHERE exam_id=?`, examID).Scan(&card.ClassSize)

	// Attendance %
	s.db.QueryRow(`
		SELECT ROUND(SUM(status='present')/COUNT(*)*100,2)
		FROM attendance a
		JOIN exams e ON e.term_id=a.term_id
		WHERE a.student_id=? AND e.id=?`, studentID, examID).Scan(&card.AttendancePct)

	return card, nil
}

// --- Class Result Summary ---

type ClassResultRow struct {
	StudentID   int64   `db:"student_id" json:"student_id"`
	StudentName string  `db:"student_name" json:"student_name"`
	AdmissionNo string  `db:"admission_no" json:"admission_no"`
	TotalMarks  float64 `db:"total_marks" json:"total_marks"`
	TotalMax    float64 `db:"total_max" json:"total_max"`
	Average     float64 `db:"average" json:"average"`
	Position    int     `json:"position"`
}

func (s *ReportService) GetClassResults(examID int64) ([]ClassResultRow, error) {
	var rows []ClassResultRow
	err := s.db.Select(&rows, `
		SELECT er.student_id, st.name AS student_name, st.admission_no,
		       SUM(er.marks) AS total_marks,
		       SUM(er.max_marks) AS total_max,
		       ROUND(SUM(er.marks)/SUM(er.max_marks)*100,2) AS average
		FROM exam_results er
		JOIN students st ON st.id=er.student_id
		WHERE er.exam_id=?
		GROUP BY er.student_id, st.name, st.admission_no
		ORDER BY average DESC`, examID)
	if err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].Position = i + 1
	}
	return rows, nil
}

// --- Fee Collection Report ---

type FeeCollectionReport struct {
	SchoolID    int64   `json:"school_id"`
	TermID      int64   `json:"term_id"`
	TotalBilled float64 `json:"total_billed"`
	TotalPaid   float64 `json:"total_paid"`
	Balance     float64 `json:"balance"`
	PaidCount   int     `json:"paid_count"`
	UnpaidCount int     `json:"unpaid_count"`
	PartialCount int    `json:"partial_count"`
}

func (s *ReportService) GetFeeCollection(schoolID, termID int64) (*FeeCollectionReport, error) {
	report := &FeeCollectionReport{SchoolID: schoolID, TermID: termID}
	s.db.QueryRow(`
		SELECT COALESCE(SUM(fi.amount),0),
		       COALESCE(SUM(fp_sum.total_paid),0),
		       COUNT(CASE WHEN fi.status='paid' THEN 1 END),
		       COUNT(CASE WHEN fi.status='unpaid' THEN 1 END),
		       COUNT(CASE WHEN fi.status='partial' THEN 1 END)
		FROM fee_invoices fi
		JOIN students st ON st.school_id=? AND st.id=fi.student_id
		LEFT JOIN (SELECT invoice_id, SUM(amount_paid) AS total_paid FROM fee_payments GROUP BY invoice_id) fp_sum
		          ON fp_sum.invoice_id=fi.id
		WHERE fi.term_id=?`, schoolID, termID).
		Scan(&report.TotalBilled, &report.TotalPaid, &report.PaidCount, &report.UnpaidCount, &report.PartialCount)
	report.Balance = report.TotalBilled - report.TotalPaid
	return report, nil
}

// --- Attendance Report (School-wide) ---

type AttendanceReport struct {
	TermID      int64   `json:"term_id"`
	ClassID     int64   `json:"class_id"`
	ClassName   string  `json:"class_name"`
	TotalDays   int     `json:"total_days"`
	AvgPresent  float64 `json:"avg_present_pct"`
}

func (s *ReportService) GetAttendanceSummary(schoolID, termID int64) ([]AttendanceReport, error) {
	var reports []AttendanceReport
	err := s.db.Select(&reports, `
		SELECT t.id AS term_id, c.id AS class_id, c.name AS class_name,
		       COUNT(DISTINCT a.date) AS total_days,
		       ROUND(AVG(CASE WHEN a.status='present' THEN 100.0 ELSE 0 END),2) AS avg_present
		FROM attendance a
		JOIN classes c ON c.id=a.class_id
		JOIN terms t ON t.id=a.term_id
		WHERE c.school_id=? AND a.term_id=?
		GROUP BY t.id, c.id, c.name
		ORDER BY c.name`, schoolID, termID)
	return reports, err
}
