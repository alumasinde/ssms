package notices

import (
	handlers "school-ms/internal/Modules/Notices/Handlers"
	repos "school-ms/internal/Modules/Notices/Repositories"
	services "school-ms/internal/Modules/Notices/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewNoticeRepository(db)
	svc := services.NewNoticeService(repo)
	h := handlers.NewNoticeHandler(svc)

	r.Route("/notices", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "notices.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db, "notices.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db, "notices.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "notices.delete")).Delete("/{id}", h.Delete)
	})
}
