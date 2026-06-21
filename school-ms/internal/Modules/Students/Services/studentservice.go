package services

import (
dtos   "school-ms/internal/Modules/Students/DTOs"
models "school-ms/internal/Modules/Students/Models"
repos  "school-ms/internal/Modules/Students/Repositories"
)

type StudentService struct{ repo *repos.StudentRepository }

func NewStudentService(repo *repos.StudentRepository) *StudentService {
return &StudentService{repo: repo}
}

func (s *StudentService) Create(dto dtos.CreateStudentDTO) (*models.Student, error) {
st := &models.Student{
SchoolID: dto.SchoolID, ClassID: dto.ClassID, AdmissionNo: dto.AdmissionNo,
FirstName: dto.FirstName, MiddleName: dto.MiddleName, LastName: dto.LastName,
Gender: dto.Gender, DOB: dto.DOB, Nationality: dto.Nationality,
NationalID: dto.NationalID, Religion: dto.Religion,
BloodGroup: dto.BloodGroup, Address: dto.Address,
MedicalNotes: dto.MedicalNotes, PhotoURL: dto.PhotoURL,
}
return st, s.repo.Create(st)
}

func (s *StudentService) GetByID(id int64) (*models.Student, error) { return s.repo.FindByID(id) }

func (s *StudentService) List(schoolID int64, page, perPage int) ([]models.Student, int64, error) {
return s.repo.ListBySchool(schoolID, page, perPage)
}

func (s *StudentService) ListByClass(classID int64) ([]models.Student, error) {
return s.repo.ListByClass(classID)
}

func (s *StudentService) Search(schoolID int64, q string) ([]models.Student, error) {
return s.repo.Search(schoolID, q)
}

func (s *StudentService) ListByParentUser(userID int64) ([]models.Student, error) {
return s.repo.FindByParentUser(userID)
}

func (s *StudentService) ListByTeacherUser(userID int64) ([]models.Student, error) {
return s.repo.FindByTeacherUser(userID)
}

func (s *StudentService) IsParentOfStudent(userID, studentID int64) (bool, error) {
return s.repo.IsParentOfStudent(userID, studentID)
}

func (s *StudentService) IsTeacherOfStudent(userID, studentID int64) (bool, error) {
return s.repo.IsTeacherOfStudent(userID, studentID)
}

func (s *StudentService) Update(id int64, dto dtos.UpdateStudentDTO) error {
st := &models.Student{
ID: id, ClassID: dto.ClassID,
FirstName: dto.FirstName, MiddleName: dto.MiddleName, LastName: dto.LastName,
Gender: dto.Gender, DOB: dto.DOB, Nationality: dto.Nationality,
NationalID: dto.NationalID, Religion: dto.Religion,
BloodGroup: dto.BloodGroup, Address: dto.Address,
MedicalNotes: dto.MedicalNotes, PhotoURL: dto.PhotoURL,
}
return s.repo.Update(st)
}

func (s *StudentService) Deactivate(id, actorID int64) error {
return s.repo.SoftDelete(id, actorID)
}

func (s *StudentService) Promote(dto dtos.PromoteStudentsDTO, actorID int64) (int, error) {
return s.repo.PromoteToClass(
dto.FromClassID, dto.ToClassID, dto.AcademicYearID, actorID, dto.StudentIDs)
}

func (s *StudentService) GetParents(studentID int64) ([]repos.ParentSummary, error) {
return s.repo.GetParentsByStudent(studentID)
}