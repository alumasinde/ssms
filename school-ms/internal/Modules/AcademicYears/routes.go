package academicyears

import (
	handlers "school-ms/internal/Modules/AcademicYears/Handlers"
	repos "school-ms/internal/Modules/AcademicYears/Repositories"
	services "school-ms/internal/Modules/AcademicYears/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewAcademicYearRepository(db)
	svc := services.NewAcademicYearService(repo)
	h := handlers.NewAcademicYearHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/academic-years", func(r chi.Router) {
			r.Get("/", h.List)
			r.Post("/", h.Create)
			r.Get("/current", h.GetCurrent)
			r.Get("/{id}", h.Get)
			r.Put("/{id}", h.Update)
			r.Put("/{id}/set-current", h.SetCurrent)
			r.Delete("/{id}", h.Delete)
		})
	})
}
