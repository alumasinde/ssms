package dtos

type CreateNoticeDTO struct {
	SchoolID int64  `json:"school_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Audience string `json:"audience"`
}
