package schools

import (
	handlers "school-ms/internal/Modules/Schools/Handlers"
	repos "school-ms/internal/Modules/Schools/Repositories"
	services "school-ms/internal/Modules/Schools/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewSchoolRepository(db)
	svc := services.NewSchoolService(repo)
	h := handlers.NewSchoolHandler(svc)

	r.Route("/schools", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "schools.view")).Get("/", h.ListSchools)
		r.With(middleware.RequirePermission(db, "schools.create")).Post("/", h.CreateSchool)
		r.With(middleware.RequirePermission(db, "schools.view")).Get("/{id}", h.GetSchool)
		r.With(middleware.RequirePermission(db, "schools.edit")).Put("/{id}", h.UpdateSchool)

		// Terms (nested under schools for now)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/{id}/terms", h.ListTerms)
		r.With(middleware.RequirePermission(db, "academic_years.create")).Post("/{id}/terms", h.CreateTerm)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/{id}/terms/current", h.GetCurrentTerm)
	})
}
