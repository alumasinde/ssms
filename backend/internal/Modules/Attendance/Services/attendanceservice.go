package services

import (
	dtos "school-ms/internal/Modules/Attendance/DTOs"
	models "school-ms/internal/Modules/Attendance/Models"
	repos "school-ms/internal/Modules/Attendance/Repositories"
)

type AttendanceService struct{ repo *repos.AttendanceRepository }

func NewAttendanceService(repo *repos.AttendanceRepository) *AttendanceService { return &AttendanceService{repo: repo} }

func (s *AttendanceService) Mark(dto dtos.MarkAttendanceDTO, recordedBy int64) error {
	var records []models.Attendance
	for _, rec := range dto.Records {
		records = append(records, models.Attendance{
			StudentID: rec.StudentID, ClassID: dto.ClassID, TermID: dto.TermID,
			RecordedBy: recordedBy, Date: dto.Date, Status: rec.Status, Remark: rec.Remark,
		})
	}
	return s.repo.BulkUpsert(records)
}

func (s *AttendanceService) GetByClassDate(classID int64, date string) ([]models.Attendance, error) {
	return s.repo.ListByClassDate(classID, date)
}

func (s *AttendanceService) GetByStudent(studentID, termID int64) ([]models.Attendance, error) {
	return s.repo.ListByStudent(studentID, termID)
}

func (s *AttendanceService) GetSummary(classID, termID int64) ([]models.AttendanceSummary, error) {
	return s.repo.SummaryByClass(classID, termID)
}
