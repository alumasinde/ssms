package services

import (
	dtos   "school-ms/internal/Modules/StaffAttendance/DTOs"
	models "school-ms/internal/Modules/StaffAttendance/Models"
	repos  "school-ms/internal/Modules/StaffAttendance/Repositories"
)

type StaffAttendanceService struct{ repo *repos.StaffAttendanceRepository }
func NewStaffAttendanceService(r *repos.StaffAttendanceRepository) *StaffAttendanceService { return &StaffAttendanceService{repo:r} }

func (s *StaffAttendanceService) Mark(dto dtos.MarkStaffDTO, schoolID, recordedBy int64) error {
	var records []models.StaffAttendance
	for _,e := range dto.Records {
		records = append(records, models.StaffAttendance{
			TeacherID:e.TeacherID,SchoolID:schoolID,Date:dto.Date,
			Status:e.Status,CheckIn:e.CheckIn,CheckOut:e.CheckOut,
			RecordedBy:recordedBy,Remark:e.Remark})
	}
	return s.repo.BulkUpsert(records)
}
func (s *StaffAttendanceService) ListByDate(schoolID int64, date string) ([]models.StaffAttendanceDetail,error) {
	return s.repo.ListByDate(schoolID,date)
}
func (s *StaffAttendanceService) Summary(schoolID int64, yearMonth string) ([]models.StaffAttendanceDetail,error) {
	return s.repo.SummaryByMonth(schoolID,yearMonth)
}
