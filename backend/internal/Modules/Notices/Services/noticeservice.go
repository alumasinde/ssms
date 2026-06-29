package services

import (
	dtos "school-ms/internal/Modules/Notices/DTOs"
	models "school-ms/internal/Modules/Notices/Models"
	repos "school-ms/internal/Modules/Notices/Repositories"
)

type NoticeService struct{ repo *repos.NoticeRepository }

func NewNoticeService(repo *repos.NoticeRepository) *NoticeService { return &NoticeService{repo: repo} }

func (s *NoticeService) Create(dto dtos.CreateNoticeDTO, authorID int64) (*models.Notice, error) {
	n := &models.Notice{SchoolID: dto.SchoolID, AuthorID: authorID, Title: dto.Title, Body: dto.Body, Audience: dto.Audience}
	return n, s.repo.Create(n)
}

func (s *NoticeService) List(schoolID int64, audience string) ([]models.Notice, error) {
	return s.repo.List(schoolID, audience)
}

func (s *NoticeService) GetByID(id int64) (*models.Notice, error) { return s.repo.FindByID(id) }
func (s *NoticeService) Delete(id int64) error                     { return s.repo.Delete(id) }
