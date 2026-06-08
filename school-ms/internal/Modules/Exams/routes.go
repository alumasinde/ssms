package exams

import (
	handlers "school-ms/internal/Modules/Exams/Handlers"
	repos "school-ms/internal/Modules/Exams/Repositories"
	services "school-ms/internal/Modules/Exams/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewExamRepository(db)
	svc := services.NewExamService(repo)
	h := handlers.NewExamHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/exams", func(r chi.Router) {
			r.Get("/", h.ListExams)
			r.Post("/", h.CreateExam)
			r.Get("/{id}", h.GetExam)
			r.Get("/{id}/results", h.GetResults)
			r.Post("/results", h.SubmitResults)
			r.Get("/student/{studentId}/results", h.GetStudentResults)
		})
		r.Route("/grade-scales", func(r chi.Router) {
			r.Get("/", h.GetGradeScales)
			r.Post("/", h.CreateGradeScale)
		})
	})
}
