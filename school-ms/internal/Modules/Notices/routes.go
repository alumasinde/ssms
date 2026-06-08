package notices

import (
	handlers "school-ms/internal/Modules/Notices/Handlers"
	repos "school-ms/internal/Modules/Notices/Repositories"
	services "school-ms/internal/Modules/Notices/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewNoticeRepository(db)
	svc := services.NewNoticeService(repo)
	h := handlers.NewNoticeHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/notices", func(r chi.Router) {
			r.Get("/", h.List)
			r.Post("/", h.Create)
			r.Get("/{id}", h.Get)
			r.Delete("/{id}", h.Delete)
		})
	})
}
