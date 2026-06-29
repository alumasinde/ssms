package tenants

import (
	handlers "school-ms/internal/Modules/Tenants/Handlers"
	repos "school-ms/internal/Modules/Tenants/Repositories"
	services "school-ms/internal/Modules/Tenants/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewTenantRepository(db)
	svc := services.NewTenantService(repo)
	h := handlers.NewTenantHandler(svc)

	r.Route("/tenants", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Use(middleware.RequireRole("superadmin"))
		r.Get("/", h.List)
		r.Post("/", h.Create)
		r.Get("/{id}", h.Get)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}
