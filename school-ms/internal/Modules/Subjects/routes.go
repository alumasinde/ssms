package subjects

import (
	handlers "school-ms/internal/Modules/Subjects/Handlers"
	repos "school-ms/internal/Modules/Subjects/Repositories"
	services "school-ms/internal/Modules/Subjects/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewSubjectRepository(db)
	svc := services.NewSubjectService(repo)
	h := handlers.NewSubjectHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/subjects", func(r chi.Router) {
			r.Get("/", h.List)
			r.Post("/", h.Create)
			r.Get("/{id}", h.Get)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}
