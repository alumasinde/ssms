package classes

import (
	handlers "school-ms/internal/Modules/Classes/Handlers"
	repos "school-ms/internal/Modules/Classes/Repositories"
	services "school-ms/internal/Modules/Classes/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewClassRepository(db)
	svc := services.NewClassService(repo)
	h := handlers.NewClassHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/classes", func(r chi.Router) {
			r.Get("/", h.List)
			r.Post("/", h.Create)
			r.Get("/{id}", h.Get)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}
