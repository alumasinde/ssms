package users
import (
	handlers "school-ms/internal/Modules/Users/Handlers"
	repos    "school-ms/internal/Modules/Users/Repositories"
	services "school-ms/internal/Modules/Users/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewUserRepository(db)
	svc  := services.NewUserService(repo)
	h    := handlers.NewUserHandler(svc)

	// Any authenticated user — self-service endpoints
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/users/me", h.Me)
		r.Put("/users/me/password", h.ChangePassword)
	})

	// Admin / superadmin — user management
	// IMPORTANT: static paths (/school) MUST be registered before wildcard (/{id})
	// to prevent chi matching the literal string "school" as an ID param.
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Use(middleware.RequireRole("superadmin", "admin"))

		r.Post("/users", h.Register)
		r.Get("/users", h.List)
		r.Get("/users/school", h.ListBySchool)   // static — before /{id}

		r.Get("/users/{id}", h.Get)
		r.Put("/users/{id}/role", h.UpdateRole)
		r.Post("/users/{id}/activate", h.Activate)
		r.Delete("/users/{id}", h.Deactivate)
	})
}