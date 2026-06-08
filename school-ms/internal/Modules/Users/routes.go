package users

import (
	handlers "school-ms/internal/Modules/Users/Handlers"
	repos "school-ms/internal/Modules/Users/Repositories"
	services "school-ms/internal/Modules/Users/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewUserRepository(db)
	svc := services.NewUserService(repo)
	h := handlers.NewUserHandler(svc)

	// Superadmin: create users for a tenant
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Use(middleware.RequireRole("superadmin", "admin"))
		r.Post("/users", h.Register)
		r.Get("/users", h.List)
		r.Delete("/users/{id}", h.Deactivate)
	})

	// Any authenticated user
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/users/me", h.Me)
		r.Put("/users/me/password", h.ChangePassword)
	})
}
