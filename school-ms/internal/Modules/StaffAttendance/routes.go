package staffattendance

import (
	handlers "school-ms/internal/Modules/StaffAttendance/Handlers"
	repos    "school-ms/internal/Modules/StaffAttendance/Repositories"
	services "school-ms/internal/Modules/StaffAttendance/Services"
	"school-ms/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewStaffAttendanceRepository(db)
	svc  := services.NewStaffAttendanceService(repo)
	h    := handlers.NewStaffAttendanceHandler(svc)
	r.Route("/staff-attendance", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db,"staff_attendance.view")).Get("/", h.ListByDate)
		r.With(middleware.RequirePermission(db,"staff_attendance.view")).Get("/summary", h.Summary)
		r.With(middleware.RequirePermission(db,"staff_attendance.mark")).Post("/", h.Mark)
	})
}
