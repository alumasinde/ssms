package repositories

import (
	models "school-ms/internal/Modules/Exams/Models"
	"github.com/jmoiron/sqlx"
)

type ExamRepository struct{ db *sqlx.DB }

func NewExamRepository(db *sqlx.DB) *ExamRepository { return &ExamRepository{db: db} }

func (r *ExamRepository) CreateExam(e *models.Exam) error {
	res, err := r.db.Exec(
		`INSERT INTO exams (school_id,term_id,class_id,name,type,start_date,end_date) VALUES (?,?,?,?,?,?,?)`,
		e.SchoolID, e.TermID, e.ClassID, e.Name, e.Type, e.StartDate, e.EndDate)
	if err != nil { return err }
	id, _ := res.LastInsertId(); e.ID = id; return nil
}

func (r *ExamRepository) ListBySchool(schoolID int64) ([]models.Exam, error) {
	var list []models.Exam
	return list, r.db.Select(&list,
		`SELECT * FROM exams WHERE school_id=? ORDER BY start_date DESC`, schoolID)
}

func (r *ExamRepository) ListByTerm(termID int64) ([]models.Exam, error) {
	var list []models.Exam
	return list, r.db.Select(&list,
		`SELECT * FROM exams WHERE term_id=? ORDER BY start_date`, termID)
}

func (r *ExamRepository) ListByClass(classID int64) ([]models.Exam, error) {
	var list []models.Exam
	return list, r.db.Select(&list,
		`SELECT * FROM exams WHERE class_id=? ORDER BY start_date DESC`, classID)
}

func (r *ExamRepository) FindByID(id int64) (*models.Exam, error) {
	var e models.Exam; return &e, r.db.Get(&e, `SELECT * FROM exams WHERE id=?`, id)
}

// LoadGradeScales loads all scales for a school into a map for in-memory resolution.
func (r *ExamRepository) LoadGradeScales(schoolID int64) ([]models.GradeScale, error) {
	var list []models.GradeScale
	return list, r.db.Select(&list,
		`SELECT * FROM grade_scales WHERE school_id=? ORDER BY min_score DESC`, schoolID)
}

func resolveGradeFromScales(scales []models.GradeScale, score float64) string {
	for _, gs := range scales {
		if score >= gs.MinScore && score <= gs.MaxScore {
			return gs.Grade
		}
	}
	return "N/A"
}

func (r *ExamRepository) BulkUpsertResults(results []models.ExamResult, scales []models.GradeScale) error {
	tx, err := r.db.Beginx()
	if err != nil { return err }
	
	q := `INSERT INTO exam_results (exam_id,student_id,subject_id,class_id,graded_by,marks,max_marks,grade,remarks)
	      VALUES (?,?,?,?,?,?,?,?,?)
	      ON DUPLICATE KEY UPDATE marks=VALUES(marks),grade=VALUES(grade),
	      remarks=VALUES(remarks),graded_by=VALUES(graded_by)`
	for _, res := range results {
		pct := 0.0
		if res.MaxMarks > 0 { pct = res.Marks / res.MaxMarks * 100 }
		res.Grade = resolveGradeFromScales(scales, pct)
		if _, err := tx.Exec(q, res.ExamID, res.StudentID, res.SubjectID, res.ClassID,
			res.GradedBy, res.Marks, res.MaxMarks, res.Grade, res.Remarks); err != nil {
			tx.Rollback(); return err
		}
	}
	return tx.Commit()
}

func (r *ExamRepository) GetResultsByExam(examID int64) ([]models.ExamResult, error) {
	var list []models.ExamResult
	return list, r.db.Select(&list, `SELECT * FROM exam_results WHERE exam_id=?`, examID)
}

func (r *ExamRepository) GetStudentResults(studentID, examID int64) ([]models.ExamResult, error) {
	var list []models.ExamResult
	return list, r.db.Select(&list,
		`SELECT * FROM exam_results WHERE student_id=? AND exam_id=?`, studentID, examID)
}

func (r *ExamRepository) GetGradeScales(schoolID int64) ([]models.GradeScale, error) {
	var list []models.GradeScale
	return list, r.db.Select(&list,
		`SELECT * FROM grade_scales WHERE school_id=? ORDER BY min_score DESC`, schoolID)
}

func (r *ExamRepository) FindGradeScaleByID(id int64) (*models.GradeScale, error) {
	var gs models.GradeScale
	return &gs, r.db.Get(&gs, `SELECT * FROM grade_scales WHERE id=?`, id)
}

func (r *ExamRepository) UpdateGradeScale(id int64, gs *models.GradeScale) error {
	_, err := r.db.Exec(
		`UPDATE grade_scales SET grade=?, min_score=?, max_score=?, remark=? WHERE id=?`,
		gs.Grade, gs.MinScore, gs.MaxScore, gs.Remark, id)
	return err
}

func (r *ExamRepository) DeleteGradeScale(id int64) error {
	_, err := r.db.Exec(`DELETE FROM grade_scales WHERE id=?`, id)
	return err
}

func (r *ExamRepository) CreateGradeScale(gs *models.GradeScale) error {
	res, err := r.db.Exec(
		`INSERT INTO grade_scales (school_id,grade,min_score,max_score,remark) VALUES (?,?,?,?,?)`,
		gs.SchoolID, gs.Grade, gs.MinScore, gs.MaxScore, gs.Remark)
	if err != nil { return err }
	id, _ := res.LastInsertId(); gs.ID = id; return nil
}

func (r *ExamRepository) GetResultsByExamEnriched(examID int64) ([]models.ExamResultDetail, error) {
	var list []models.ExamResultDetail
	err := r.db.Select(&list, `
		SELECT er.id, er.exam_id, er.student_id, er.subject_id, er.class_id, er.graded_by,
		       er.marks, er.max_marks, er.grade, er.remarks,
		       CONCAT(s.first_name,' ',s.last_name) AS student_name,
		       s.admission_no,
		       sub.name AS subject_name, sub.code AS subject_code
		FROM exam_results er
		JOIN students s   ON s.id   = er.student_id
		JOIN subjects sub ON sub.id = er.subject_id
		WHERE er.exam_id = ?
		ORDER BY s.last_name, s.first_name, sub.name`, examID)
	return list, err
}

func (r *ExamRepository) GetResultsByExamAndClass(examID, classID int64) ([]models.ExamResultDetail, error) {
	var list []models.ExamResultDetail
	err := r.db.Select(&list, `
		SELECT er.id, er.exam_id, er.student_id, er.subject_id, er.class_id, er.graded_by,
		       er.marks, er.max_marks, er.grade, er.remarks,
		       CONCAT(s.first_name,' ',s.last_name) AS student_name,
		       s.admission_no,
		       sub.name AS subject_name, sub.code AS subject_code
		FROM exam_results er
		JOIN students s   ON s.id   = er.student_id
		JOIN subjects sub ON sub.id = er.subject_id
		WHERE er.exam_id=? AND er.class_id=?
		ORDER BY s.last_name, s.first_name, sub.name`, examID, classID)
	return list, err
}
