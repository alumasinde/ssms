package classes

import (
	handlers "school-ms/internal/Modules/Classes/Handlers"
	repos    "school-ms/internal/Modules/Classes/Repositories"
	services "school-ms/internal/Modules/Classes/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	// Class CRUD
	classRepo := repos.NewClassRepository(db)
	classSvc  := services.NewClassService(classRepo)
	classH    := handlers.NewClassHandler(classSvc)

	// Class ↔ Subject assignment
	csRepo := repos.NewClassSubjectRepository(db)
	csSvc  := services.NewClassSubjectService(csRepo)
	csH    := handlers.NewClassSubjectHandler(csSvc)

	r.Route("/classes", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		// ── Class CRUD ────────────────────────────────────────────────────────
		r.With(middleware.RequirePermission(db, "classes.view")).Get("/", classH.List)
		r.With(middleware.RequirePermission(db, "classes.create")).Post("/", classH.Create)
		r.With(middleware.RequirePermission(db, "classes.view")).Get("/{id}", classH.Get)
		r.With(middleware.RequirePermission(db, "classes.edit")).Put("/{id}", classH.Update)
		r.With(middleware.RequirePermission(db, "classes.delete")).Delete("/{id}", classH.Delete)

		// ── Class ↔ Subject ───────────────────────────────────────────────────
		// NOTE: static path /unassigned MUST be registered before the
		// wildcard /{subjectId} route to avoid chi matching "unassigned" as an ID.
		r.With(middleware.RequirePermission(db, "classes.view")).
			Get("/{id}/subjects", csH.List)

		r.With(middleware.RequirePermission(db, "classes.view")).
			Get("/{id}/subjects/unassigned", csH.ListUnassigned)

		r.With(middleware.RequirePermission(db, "classes.edit")).
			Post("/{id}/subjects", csH.Assign)

		r.With(middleware.RequirePermission(db, "classes.edit")).
			Delete("/{id}/subjects/{subjectId}", csH.Remove)
	})
}
