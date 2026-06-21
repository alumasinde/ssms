package terms

import (
	handlers "school-ms/internal/Modules/Terms/Handlers"
	repos    "school-ms/internal/Modules/Terms/Repositories"
	services "school-ms/internal/Modules/Terms/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewTermRepository(db)
	svc  := services.NewTermService(repo)
	h    := handlers.NewTermHandler(svc)

	r.Route("/terms", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/", h.ListBySchool)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/current", h.GetCurrent)
		r.With(middleware.RequirePermission(db, "academic_years.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "academic_years.edit")).Put("/{id}", h.Update)
		r.With(middleware.RequirePermission(db, "academic_years.edit")).Post("/{id}/set-current", h.SetCurrent)
		r.With(middleware.RequirePermission(db, "academic_years.edit")).Delete("/{id}", h.Delete)
	})

	r.Route("/academic-years/{yearId}/terms", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/", h.ListByYear)
	})
}
