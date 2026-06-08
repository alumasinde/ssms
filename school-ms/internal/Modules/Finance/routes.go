package finance

import (
	handlers "school-ms/internal/Modules/Finance/Handlers"
	repos "school-ms/internal/Modules/Finance/Repositories"
	services "school-ms/internal/Modules/Finance/Services"
	middleware "school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewFinanceRepository(db)
	svc := services.NewFinanceService(repo)
	h := handlers.NewFinanceHandler(svc)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Route("/finance", func(r chi.Router) {
			r.Get("/fee-types", h.ListFeeTypes)
			r.Post("/fee-types", h.CreateFeeType)
			r.Post("/invoices/generate", h.GenerateInvoices)
			r.Get("/invoices/{invoiceId}/payments", h.InvoicePayments)
			r.Post("/payments", h.RecordPayment)
			r.Get("/statement/student/{studentId}", h.StudentStatement)
		})
	})
}
