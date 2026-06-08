package students

import (
	handlers "school-ms/internal/Modules/Students/Handlers"
	repos "school-ms/internal/Modules/Students/Repositories"
	services "school-ms/internal/Modules/Students/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewStudentRepository(db)
	svc := services.NewStudentService(repo)
	h := handlers.NewStudentHandler(svc)

	r.Route("/students", func(r chi.Router) {
        r.Use(middleware.Authenticate)

        r.Group(func(r chi.Router) {
            r.Use(middleware.RequirePermission(db, "students.view"))
            r.Get("/", h.List)
            r.Get("/class/{classId}", h.ListByClass)
            r.Get("/{id}", h.Get)
        })

        r.Group(func(r chi.Router) {
            r.Use(middleware.RequirePermission(db, "students.create"))
            r.Post("/", h.Create)
        })

        r.Group(func(r chi.Router) {
            r.Use(middleware.RequirePermission(db, "students.edit"))
            r.Put("/{id}", h.Update)
            r.Post("/{id}", h.Update)
        })

        r.Group(func(r chi.Router) {
            r.Use(middleware.RequirePermission(db, "students.delete"))
            r.Delete("/{id}", h.Deactivate)
        })
    })
}
