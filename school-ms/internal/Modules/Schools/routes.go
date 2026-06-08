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

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/schools", func(r chi.Router) {
			r.Get("/", h.ListSchools)
			r.Post("/", h.CreateSchool)
			r.Get("/{id}", h.GetSchool)
			r.Put("/{id}", h.UpdateSchool)
		})
	})
}
