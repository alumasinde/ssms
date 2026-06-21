package assignments

import (
	handlers "school-ms/internal/Modules/Assignments/Handlers"
	repos    "school-ms/internal/Modules/Assignments/Repositories"
	services "school-ms/internal/Modules/Assignments/Services"
	"school-ms/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewAssignmentRepository(db)
	svc  := services.NewAssignmentService(repo)
	h    := handlers.NewAssignmentHandler(svc)
	r.Route("/assignments", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db,"assignments.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db,"assignments.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db,"assignments.edit")).Delete("/{id}", h.Delete)
	})
}
