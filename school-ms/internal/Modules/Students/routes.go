package students

import (
	handlers "school-ms/internal/Modules/Students/Handlers"
	repos    "school-ms/internal/Modules/Students/Repositories"
	services "school-ms/internal/Modules/Students/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewStudentRepository(db)
	svc  := services.NewStudentService(repo)
	h    := handlers.NewStudentHandler(svc, db)

	r.Route("/students", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		// ── Static paths MUST come before /{id} wildcard ─────────────────
		// chi is most-specific-first but only within the same method;
		// registering static routes before wildcard is the safest pattern.
		r.With(middleware.RequirePermission(db, "students.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db, "students.edit")).Post("/promote", h.Promote)
		r.With(middleware.RequirePermission(db, "students.view")).Get("/search", h.Search)
		r.With(middleware.RequirePermission(db, "students.view")).Get("/class/{classId}", h.ListByClass)

		// ── Paginated / role-filtered list ────────────────────────────────
		r.With(middleware.RequirePermission(db, "students.view")).Get("/", h.List)

		// ── Single student — wildcard LAST ────────────────────────────────
		r.With(middleware.RequirePermission(db, "students.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "students.view")).Get("/{id}/parents", h.GetParents)
		r.With(middleware.RequirePermission(db, "students.edit")).Put("/{id}", h.Update)
		r.With(middleware.RequirePermission(db, "students.delete")).Delete("/{id}", h.Deactivate)
	})
}