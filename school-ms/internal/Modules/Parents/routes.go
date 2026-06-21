package parents

import (
	handlers "school-ms/internal/Modules/Parents/Handlers"
	repos "school-ms/internal/Modules/Parents/Repositories"
	services "school-ms/internal/Modules/Parents/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewParentRepository(db)
	svc := services.NewParentService(repo)
	h := handlers.NewParentHandler(svc)

	r.Route("/parents", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "parents.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db, "parents.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db, "parents.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "parents.view")).Get("/{id}/students", h.GetStudentParents)
		r.With(middleware.RequirePermission(db, "parents.create")).Post("/link-student", h.LinkStudent)
	})
}
