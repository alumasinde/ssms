package services

import (
	models "school-ms/internal/Modules/Classes/Models"
	repos  "school-ms/internal/Modules/Classes/Repositories"
)

type ClassSubjectService struct{ repo *repos.ClassSubjectRepository }

func NewClassSubjectService(repo *repos.ClassSubjectRepository) *ClassSubjectService {
	return &ClassSubjectService{repo: repo}
}

// Assign adds or updates a subject assignment for a class.
// schoolID is required because class_subjects.school_id is NOT NULL.
func (s *ClassSubjectService) Assign(
	classID, subjectID, schoolID int64,
	compulsory bool,
) error {
	return s.repo.AssignSubject(classID, subjectID, schoolID, compulsory)
}

func (s *ClassSubjectService) Remove(classID, subjectID int64) error {
	return s.repo.RemoveSubject(classID, subjectID)
}

func (s *ClassSubjectService) List(classID int64) ([]models.ClassSubject, error) {
	return s.repo.ListSubjects(classID)
}

func (s *ClassSubjectService) ListUnassigned(
	classID, schoolID int64,
) ([]models.ClassSubject, error) {
	return s.repo.ListUnassignedSubjects(classID, schoolID)
}
