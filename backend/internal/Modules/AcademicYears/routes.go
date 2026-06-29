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

	r.Route("/academic-years", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/current", h.GetCurrent)
		r.With(middleware.RequirePermission(db, "academic_years.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db, "academic_years.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "academic_years.edit")).Put("/{id}", h.Update)
		r.With(middleware.RequirePermission(db, "academic_years.edit")).Post("/{id}/set-current", h.SetCurrent)
		r.With(middleware.RequirePermission(db, "academic_years.edit")).Delete("/{id}", h.Delete)
	})
}
