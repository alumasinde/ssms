package exams

import (
	handlers "school-ms/internal/Modules/Exams/Handlers"
	repos "school-ms/internal/Modules/Exams/Repositories"
	services "school-ms/internal/Modules/Exams/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewExamRepository(db)
	svc := services.NewExamService(repo)
	h := handlers.NewExamHandler(svc)

	r.Route("/exams", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "exams.view")).Get("/", h.ListExams)
		r.With(middleware.RequirePermission(db, "exams.create")).Post("/", h.CreateExam)
		r.With(middleware.RequirePermission(db, "exams.view")).Get("/{id}", h.GetExam)
		r.With(middleware.RequirePermission(db, "exams.view")).Get("/{id}/results", h.GetResults)
		r.With(middleware.RequirePermission(db, "exams.grade")).Post("/results", h.SubmitResults)
		r.With(middleware.RequirePermission(db, "exams.view")).Get("/student/{studentId}/results", h.GetStudentResults)
	})

	r.Route("/grade-scales", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.With(middleware.RequirePermission(db, "exams.view")).Get("/", h.GetGradeScales)
		r.With(middleware.RequirePermission(db, "exams.create")).Post("/", h.CreateGradeScale)
		r.With(middleware.RequirePermission(db, "exams.update")).Put("/{id}", h.UpdateGradeScale)
		r.With(middleware.RequirePermission(db, "exams.delete")).Delete("/{id}", h.DeleteGradeScale)
	})
}
