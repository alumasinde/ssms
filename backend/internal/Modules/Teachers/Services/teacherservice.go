package services

import (
dtos   "school-ms/internal/Modules/Teachers/DTOs"
models "school-ms/internal/Modules/Teachers/Models"
repos  "school-ms/internal/Modules/Teachers/Repositories"
)

type TeacherService struct{ repo *repos.TeacherRepository }

func NewTeacherService(repo *repos.TeacherRepository) *TeacherService {
return &TeacherService{repo: repo}
}

func (s *TeacherService) Create(dto dtos.CreateTeacherDTO) (*models.Teacher, error) {
t := &models.Teacher{
UserID: dto.UserID, SchoolID: dto.SchoolID, EmployeeNo: dto.EmployeeNo,
Phone: dto.Phone, Gender: dto.Gender, DOB: dto.DOB, TSCNo: dto.TSCNo,
HireDate: dto.HireDate, Qualification: dto.Qualification,
Specialization: dto.Specialization, EmploymentType: dto.EmploymentType,
NationalID: dto.NationalID, Address: dto.Address, PhotoURL: dto.PhotoURL,
}
return t, s.repo.Create(t)
}

func (s *TeacherService) GetByID(id int64) (*models.TeacherDetail, error) { return s.repo.FindByID(id) }
func (s *TeacherService) List(schoolID int64) ([]models.TeacherDetail, error) { return s.repo.ListBySchool(schoolID) }

func (s *TeacherService) Update(id int64, dto dtos.UpdateTeacherDTO) (*models.TeacherDetail, error) {
t, err := s.repo.FindByID(id)
if err != nil { return nil, err }
// Patch — only overwrite non-nil fields
if dto.Phone != nil          { t.Phone = dto.Phone }
if dto.Gender != nil         { t.Gender = dto.Gender }
if dto.DOB != nil            { t.DOB = dto.DOB }
if dto.TSCNo != nil          { t.TSCNo = dto.TSCNo }
if dto.HireDate != nil       { t.HireDate = dto.HireDate }
if dto.Qualification != nil  { t.Qualification = dto.Qualification }
if dto.Specialization != nil { t.Specialization = dto.Specialization }
if dto.NationalID != nil     { t.NationalID = dto.NationalID }
if dto.Address != nil        { t.Address = dto.Address }
if dto.PhotoURL != nil       { t.PhotoURL = dto.PhotoURL }
if dto.IsClassTeacher != nil { t.IsClassTeacher = *dto.IsClassTeacher }
if dto.EmploymentType != nil { t.EmploymentType = *dto.EmploymentType }
if err := s.repo.Update(&t.Teacher); err != nil { return nil, err }
return t, nil
}

func (s *TeacherService) Delete(id, actorID int64) error { return s.repo.SoftDelete(id, actorID) }

func (s *TeacherService) AssignSubject(teacherID, schoolID int64, dto dtos.AssignSubjectDTO) error {
return s.repo.AssignSubject(&models.TeacherSubject{
TeacherID: teacherID, SubjectID: dto.SubjectID,
ClassID: dto.ClassID, SchoolID: schoolID,
})
}

func (s *TeacherService) GetSubjects(teacherID int64) ([]models.TeacherSubjectDetail, error) {
return s.repo.GetSubjectsByTeacher(teacherID)
}

func (s *TeacherService) RemoveSubject(teacherID, subjectID, classID int64) error {
return s.repo.RemoveSubject(teacherID, subjectID, classID)
}