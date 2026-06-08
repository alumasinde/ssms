package auth

import (
	"school-ms/internal/Modules/Auth/Handlers"
	"school-ms/internal/Modules/Auth/Repositories"
	"school-ms/internal/Modules/Auth/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repositories.NewAuthRepository(db)
	svc := services.NewAuthService(repo)
	h := handlers.NewAuthHandler(svc)

	// Public – tenant resolved from Host header by ResolveTenantFromDB middleware
	r.Post("/auth/login", h.Login)

	// Protected
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/auth/me", h.Whoami)
	})
}
