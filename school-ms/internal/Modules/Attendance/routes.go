package attendance

import (
	handlers "school-ms/internal/Modules/Attendance/Handlers"
	repos "school-ms/internal/Modules/Attendance/Repositories"
	services "school-ms/internal/Modules/Attendance/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewAttendanceRepository(db)
	svc := services.NewAttendanceService(repo)
	h := handlers.NewAttendanceHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/attendance", func(r chi.Router) {
			r.Post("/", h.Mark)
			r.Get("/class/{classId}", h.GetByClassDate)
			r.Get("/student/{studentId}", h.GetByStudent)
			r.Get("/summary/class/{classId}", h.Summary)
		})
	})
}
