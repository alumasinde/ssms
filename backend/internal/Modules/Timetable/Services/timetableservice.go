package services

import (
	dtos   "school-ms/internal/Modules/Timetable/DTOs"
	models "school-ms/internal/Modules/Timetable/Models"
	repos  "school-ms/internal/Modules/Timetable/Repositories"
)

type TimetableService struct{ repo *repos.TimetableRepository }
func NewTimetableService(r *repos.TimetableRepository) *TimetableService { return &TimetableService{repo: r} }

func (s *TimetableService) Create(dto dtos.CreateSlotDTO, schoolID int64) (*models.TimetableSlot, error) {
	slot := &models.TimetableSlot{SchoolID:schoolID,ClassID:dto.ClassID,SubjectID:dto.SubjectID,
		TeacherID:dto.TeacherID,TermID:dto.TermID,DayOfWeek:dto.DayOfWeek,
		StartTime:dto.StartTime,EndTime:dto.EndTime,Room:dto.Room}
	return slot, s.repo.Create(slot)
}
func (s *TimetableService) Update(id int64, dto dtos.CreateSlotDTO) error {
	return s.repo.Update(&models.TimetableSlot{ID:id,SubjectID:dto.SubjectID,TeacherID:dto.TeacherID,
		DayOfWeek:dto.DayOfWeek,StartTime:dto.StartTime,EndTime:dto.EndTime,Room:dto.Room})
}
func (s *TimetableService) Delete(id int64) error { return s.repo.Delete(id) }
func (s *TimetableService) ListByClass(classID,termID int64) ([]models.TimetableSlotDetail,error) { return s.repo.ListByClass(classID,termID) }
func (s *TimetableService) ListByTeacher(teacherID,termID int64) ([]models.TimetableSlotDetail,error) { return s.repo.ListByTeacher(teacherID,termID) }
func (s *TimetableService) ListBySchool(schoolID,termID int64) ([]models.TimetableSlotDetail,error) { return s.repo.ListBySchool(schoolID,termID) }
