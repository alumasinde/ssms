package repositories

import (
	models "school-ms/internal/Modules/Notices/Models"
	"github.com/jmoiron/sqlx"
)

type NoticeRepository struct{ db *sqlx.DB }

func NewNoticeRepository(db *sqlx.DB) *NoticeRepository { return &NoticeRepository{db: db} }

func (r *NoticeRepository) Create(n *models.Notice) error {
	res, err := r.db.Exec(
		`INSERT INTO notices (school_id,author_id,title,body,audience,published_at) VALUES (?,?,?,?,?,NOW())`,
		n.SchoolID, n.AuthorID, n.Title, n.Body, n.Audience)
	if err != nil { return err }
	id, _ := res.LastInsertId(); n.ID = id; return nil
}

func (r *NoticeRepository) List(schoolID int64, audience string) ([]models.Notice, error) {
	var list []models.Notice
	if audience != "" && audience != "all" {
		return list, r.db.Select(&list,
			`SELECT * FROM notices WHERE school_id=? AND (audience='all' OR audience=?) ORDER BY published_at DESC`,
			schoolID, audience)
	}
	return list, r.db.Select(&list,
		`SELECT * FROM notices WHERE school_id=? ORDER BY published_at DESC`, schoolID)
}

func (r *NoticeRepository) FindByID(id int64) (*models.Notice, error) {
	var n models.Notice; return &n, r.db.Get(&n, `SELECT * FROM notices WHERE id=?`, id)
}

func (r *NoticeRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM notices WHERE id=?`, id); return err
}
