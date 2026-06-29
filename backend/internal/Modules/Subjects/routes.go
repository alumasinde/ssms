package subjects

import (
	handlers "school-ms/internal/Modules/Subjects/Handlers"
	repos "school-ms/internal/Modules/Subjects/Repositories"
	services "school-ms/internal/Modules/Subjects/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewSubjectRepository(db)
	svc := services.NewSubjectService(repo)
	h := handlers.NewSubjectHandler(svc)

	r.Route("/subjects", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "subjects.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db, "subjects.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db, "subjects.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "subjects.edit")).Put("/{id}", h.Update)
		r.With(middleware.RequirePermission(db, "subjects.delete")).Delete("/{id}", h.Delete)
	})
}
