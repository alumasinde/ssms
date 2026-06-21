package services

import (
	dtos   "school-ms/internal/Modules/Discipline/DTOs"
	models "school-ms/internal/Modules/Discipline/Models"
	repos  "school-ms/internal/Modules/Discipline/Repositories"
)

type DisciplineService struct{ repo *repos.DisciplineRepository }
func NewDisciplineService(r *repos.DisciplineRepository) *DisciplineService { return &DisciplineService{repo:r} }

func (s *DisciplineService) Create(dto dtos.CreateDisciplineDTO, schoolID, recordedBy int64) (*models.DisciplineRecord,error) {
	d := &models.DisciplineRecord{SchoolID:schoolID,StudentID:dto.StudentID,TermID:dto.TermID,
		IncidentDate:dto.IncidentDate,Type:dto.Type,Description:dto.Description,
		ActionTaken:dto.ActionTaken,RecordedBy:recordedBy}
	return d, s.repo.Create(d)
}
func (s *DisciplineService) ListBySchool(schoolID,termID int64) ([]models.DisciplineDetail,error) { return s.repo.ListBySchool(schoolID,termID) }
func (s *DisciplineService) ListByStudent(studentID int64) ([]models.DisciplineDetail,error) { return s.repo.ListByStudent(studentID) }
func (s *DisciplineService) Delete(id int64) error { return s.repo.Delete(id) }
