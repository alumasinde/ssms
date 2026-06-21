package finance

import (
	handlers "school-ms/internal/Modules/Finance/Handlers"
	repos    "school-ms/internal/Modules/Finance/Repositories"
	services "school-ms/internal/Modules/Finance/Services"
	"school-ms/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	repo := repos.NewFinanceRepository(db)
	svc  := services.NewFinanceService(repo)
	h    := handlers.NewFinanceHandler(svc)

	r.Route("/finance", func(r chi.Router) {
		r.Use(middleware.Authenticate)

		r.With(middleware.RequirePermission(db, "finance.view")).Get("/fee-types", h.ListFeeTypes)
		r.With(middleware.RequirePermission(db, "finance.create")).Post("/fee-types", h.CreateFeeType)

		r.With(middleware.RequirePermission(db, "finance.create")).Post("/invoices/generate", h.GenerateInvoices)
		r.With(middleware.RequirePermission(db, "finance.view")).Get("/invoices/{invoiceId}/payments", h.InvoicePayments)

		r.With(middleware.RequirePermission(db, "finance.create")).Post("/payments", h.RecordPayment)
		r.With(middleware.RequirePermission(db, "finance.view")).Get("/statement/student/{studentId}", h.StudentStatement)

		r.With(middleware.RequirePermission(db, "finance.create")).Post("/discounts", h.CreateDiscount)
		r.With(middleware.RequirePermission(db, "finance.view")).Get("/discounts/student/{studentId}", h.ListDiscounts)

		r.With(middleware.RequirePermission(db, "finance.create")).Post("/mpesa/stk-push", h.MpesaStkPush)
	})

	// Public M-Pesa callback — no auth (Safaricom posts to this)
	r.Post("/finance/mpesa/callback", h.MpesaCallback)
}
