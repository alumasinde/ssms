package timetable

import (
	handlers "school-ms/internal/Modules/Timetable/Handlers"
	repos    "school-ms/internal/Modules/Timetable/Repositories"
	services "school-ms/internal/Modules/Timetable/Services"
	"school-ms/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewTimetableRepository(db)
	svc  := services.NewTimetableService(repo)
	h    := handlers.NewTimetableHandler(svc)
	r.Route("/timetable", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db,"timetable.view")).Get("/", h.ListBySchool)
		r.With(middleware.RequirePermission(db,"timetable.view")).Get("/class/{classId}", h.ListByClass)
		r.With(middleware.RequirePermission(db,"timetable.view")).Get("/teacher/{teacherId}", h.ListByTeacher)
		r.With(middleware.RequirePermission(db,"timetable.edit")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db,"timetable.edit")).Put("/{id}", h.Update)
		r.With(middleware.RequirePermission(db,"timetable.edit")).Delete("/{id}", h.Delete)
	})
}
