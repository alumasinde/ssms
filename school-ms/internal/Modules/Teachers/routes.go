package teachers

import (
	handlers "school-ms/internal/Modules/Teachers/Handlers"
	repos "school-ms/internal/Modules/Teachers/Repositories"
	services "school-ms/internal/Modules/Teachers/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewTeacherRepository(db)
	svc := services.NewTeacherService(repo)
	h := handlers.NewTeacherHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/teachers", func(r chi.Router) {
			r.Get("/", h.List)
			r.Post("/", h.Create)
			r.Get("/{id}", h.Get)
			r.Get("/{id}/subjects", h.GetSubjects)
			r.Post("/assign-subject", h.AssignSubject)
		})
	})
}
