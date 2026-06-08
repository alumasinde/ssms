package routes

import (
	"net/http"

	"school-ms/internal/middleware"

	auth "school-ms/internal/Modules/Auth"
	academicyears "school-ms/internal/Modules/AcademicYears"
	attendance "school-ms/internal/Modules/Attendance"
	classes "school-ms/internal/Modules/Classes"
	exams "school-ms/internal/Modules/Exams"
	finance "school-ms/internal/Modules/Finance"
	notices "school-ms/internal/Modules/Notices"
	parents "school-ms/internal/Modules/Parents"
	reporthandlers "school-ms/internal/Modules/Reports/Handlers"
	reportservices "school-ms/internal/Modules/Reports/Services"
	schools "school-ms/internal/Modules/Schools"
	students "school-ms/internal/Modules/Students"
	subjects "school-ms/internal/Modules/Subjects"
	teachers "school-ms/internal/Modules/Teachers"
	tenants "school-ms/internal/Modules/Tenants"
	users "school-ms/internal/Modules/Users"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

func Setup(db *sqlx.DB, allowedOrigins []string) http.Handler {
	r := chi.NewRouter()

	// ── Global middleware ────────────────────────────────────────────────────
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)

	// Domain → tenant resolution (reads DB). Must come before auth routes.
	r.Use(middleware.ResolveTenantFromDB(db))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type",
			"X-Tenant-Slug", "X-Request-ID",
		},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// ── Health ───────────────────────────────────────────────────────────────
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// ── API v1 ───────────────────────────────────────────────────────────────
	r.Route("/api/v1", func(r chi.Router) {
		// Auth – login, whoami (tenant resolved from Host header)
		auth.RegisterRoutes(r, db)

		// Users – register, profile, password change
		users.RegisterRoutes(r, db)

		// Tenant management (superadmin only)
		tenants.RegisterRoutes(r, db)

		// Academic structure
		schools.RegisterRoutes(r, db)
		academicyears.RegisterRoutes(r, db)
		classes.RegisterRoutes(r, db)
		subjects.RegisterRoutes(r, db)

		// People
		teachers.RegisterRoutes(r, db)
		students.RegisterRoutes(r, db)
		parents.RegisterRoutes(r, db)

		// Operations
		attendance.RegisterRoutes(r, db)
		exams.RegisterRoutes(r, db)
		finance.RegisterRoutes(r, db)
		notices.RegisterRoutes(r, db)

		// Reports – inline (no routes.go in Reports module)
		rptSvc := reportservices.NewReportService(db)
		rptH := reporthandlers.NewReportHandler(rptSvc)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Authenticate)
			r.Route("/reports", func(r chi.Router) {
				r.Get("/report-card/{studentId}", rptH.ReportCard)
				r.Get("/class-results", rptH.ClassResults)
				r.Get("/fee-collection", rptH.FeeCollection)
				r.Get("/attendance-summary", rptH.AttendanceSummary)
			})
		})
	})

	// ── Path-prefix support: /ssms/api/v1/... ───────────────────────────────
	// Allows highwayhighschool.ac.ke/ssms to work alongside subdomain access.
	// We just re-mount the same router under /ssms.

	return r
}
