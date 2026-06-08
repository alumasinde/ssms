package parents

import (
	handlers "school-ms/internal/Modules/Parents/Handlers"
	repos "school-ms/internal/Modules/Parents/Repositories"
	services "school-ms/internal/Modules/Parents/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewParentRepository(db)
	svc := services.NewParentService(repo)
	h := handlers.NewParentHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/parents", func(r chi.Router) {
			r.Get("/", h.List)
			r.Post("/", h.Create)
			r.Get("/{id}", h.Get)
			r.Post("/link-student", h.LinkStudent)
			r.Get("/student/{studentId}", h.GetStudentParents)
		})
	})
}
