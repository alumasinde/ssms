package auth

import (
	"school-ms/internal/Modules/Auth/Handlers"
	"school-ms/internal/Modules/Auth/Repositories"
	"school-ms/internal/Modules/Auth/Services"
	"school-ms/internal/middleware"
	"school-ms/internal/pkg/ratelimit"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// RegisterRoutes wires auth endpoints.
// limiter is applied to the public login/refresh endpoints to prevent
// brute-force attacks. Pass ratelimit.New(10, time.Minute) from routes/web.go.
func RegisterRoutes(r chi.Router, db *sqlx.DB, limiter *ratelimit.Limiter) {
	repo := repositories.NewAuthRepository(db)
	svc  := services.NewAuthService(repo)
	h    := handlers.NewAuthHandler(svc)

	// ── Public (rate-limited) ─────────────────────────────────────────────────
	// ResolveTenantFromDB must be mounted globally before these routes run.
	r.Group(func(r chi.Router) {
		r.Use(limiter.Middleware)
		r.Post("/auth/login",   h.Login)
		r.Post("/auth/refresh", h.Refresh)
	})

	// ── Protected ─────────────────────────────────────────────────────────────
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Get("/auth/me",      h.Me)
		r.Post("/auth/logout", h.Logout)
	})
}
