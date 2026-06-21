package teachers
import (
	handlers "school-ms/internal/Modules/Teachers/Handlers"
	repos    "school-ms/internal/Modules/Teachers/Repositories"
	services "school-ms/internal/Modules/Teachers/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewTeacherRepository(db)
	svc  := services.NewTeacherService(repo)
	h    := handlers.NewTeacherHandler(svc)

	r.Route("/teachers", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "teachers.view")).Get("/", h.List)
		r.With(middleware.RequirePermission(db, "teachers.create")).Post("/", h.Create)

		r.With(middleware.RequirePermission(db, "teachers.view")).Get("/{id}", h.Get)
		r.With(middleware.RequirePermission(db, "teachers.edit")).Put("/{id}", h.Update)
		r.With(middleware.RequirePermission(db, "teachers.delete")).Delete("/{id}", h.Delete)

		r.With(middleware.RequirePermission(db, "teachers.view")).Get("/{id}/subjects", h.GetSubjects)
		r.With(middleware.RequirePermission(db, "teachers.edit")).Post("/{id}/subjects", h.AssignSubject)
		r.With(middleware.RequirePermission(db, "teachers.edit")).
			Delete("/{id}/subjects/{subjectId}/{classId}", h.RemoveSubject)
	})
}