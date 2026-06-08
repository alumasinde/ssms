package repositories

import (
	models "school-ms/internal/Modules/Exams/Models"
	"github.com/jmoiron/sqlx"
)

type ExamRepository struct{ db *sqlx.DB }

func NewExamRepository(db *sqlx.DB) *ExamRepository { return &ExamRepository{db: db} }

func (r *ExamRepository) CreateExam(e *models.Exam) error {
	res, err := r.db.Exec(`INSERT INTO exams (school_id,term_id,name,type,start_date,end_date) VALUES (?,?,?,?,?,?)`,
		e.SchoolID, e.TermID, e.Name, e.Type, e.StartDate, e.EndDate)
	if err != nil { return err }
	id, _ := res.LastInsertId(); e.ID = id; return nil
}

func (r *ExamRepository) ListBySchool(schoolID int64) ([]models.Exam, error) {
	var list []models.Exam
	return list, r.db.Select(&list, `SELECT * FROM exams WHERE school_id=? ORDER BY start_date DESC`, schoolID)
}

func (r *ExamRepository) FindByID(id int64) (*models.Exam, error) {
	var e models.Exam; return &e, r.db.Get(&e, `SELECT * FROM exams WHERE id=?`, id)
}

func (r *ExamRepository) BulkUpsertResults(results []models.ExamResult) error {
	tx, err := r.db.Beginx()
	if err != nil { return err }
	q := `INSERT INTO exam_results (exam_id,student_id,subject_id,graded_by,marks,max_marks,grade,remarks)
	      VALUES (?,?,?,?,?,?,?,?)
	      ON DUPLICATE KEY UPDATE marks=VALUES(marks),grade=VALUES(grade),remarks=VALUES(remarks),graded_by=VALUES(graded_by)`
	for _, res := range results {
		if _, err := tx.Exec(q, res.ExamID, res.StudentID, res.SubjectID, res.GradedBy, res.Marks, res.MaxMarks, res.Grade, res.Remarks); err != nil {
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
	return list, r.db.Select(&list, `SELECT * FROM exam_results WHERE student_id=? AND exam_id=?`, studentID, examID)
}

func (r *ExamRepository) GetGradeScales(schoolID int64) ([]models.GradeScale, error) {
	var list []models.GradeScale
	return list, r.db.Select(&list, `SELECT * FROM grade_scales WHERE school_id=? ORDER BY min_score DESC`, schoolID)
}

func (r *ExamRepository) CreateGradeScale(gs *models.GradeScale) error {
	res, err := r.db.Exec(`INSERT INTO grade_scales (school_id,grade,min_score,max_score,remark) VALUES (?,?,?,?,?)`,
		gs.SchoolID, gs.Grade, gs.MinScore, gs.MaxScore, gs.Remark)
	if err != nil { return err }
	id, _ := res.LastInsertId(); gs.ID = id; return nil
}

func (r *ExamRepository) ResolveGrade(schoolID int64, score float64) string {
	var grade string
	r.db.Get(&grade, `SELECT grade FROM grade_scales WHERE school_id=? AND ? BETWEEN min_score AND max_score LIMIT 1`, schoolID, score)
	if grade == "" { return "N/A" }
	return grade
}
