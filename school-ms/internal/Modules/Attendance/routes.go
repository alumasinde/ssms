package attendance

import (
	handlers "school-ms/internal/Modules/Attendance/Handlers"
	repos "school-ms/internal/Modules/Attendance/Repositories"
	services "school-ms/internal/Modules/Attendance/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewAttendanceRepository(db)
	svc := services.NewAttendanceService(repo)
	h := handlers.NewAttendanceHandler(svc)

	r.Route("/attendance", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "attendance.mark")).Post("/", h.Mark)
		r.With(middleware.RequirePermission(db, "attendance.view")).Get("/class/{classId}", h.GetByClassDate)
		r.With(middleware.RequirePermission(db, "attendance.view")).Get("/student/{studentId}", h.GetByStudent)
		r.With(middleware.RequirePermission(db, "attendance.view")).Get("/summary/class/{classId}", h.Summary)
	})
}
