package discipline

import (
	handlers "school-ms/internal/Modules/Discipline/Handlers"
	repos    "school-ms/internal/Modules/Discipline/Repositories"
	services "school-ms/internal/Modules/Discipline/Services"
	"school-ms/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewDisciplineRepository(db)
	svc  := services.NewDisciplineService(repo)
	h    := handlers.NewDisciplineHandler(svc)
	r.Route("/discipline", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db,"discipline.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db,"discipline.view")).Get("/student/{studentId}", h.ListByStudent)
		r.With(middleware.RequirePermission(db,"discipline.create")).Post("/", h.Create)
		r.With(middleware.RequirePermission(db,"discipline.edit")).Delete("/{id}", h.Delete)
	})
}
