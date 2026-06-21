package routes

import (
	"net/http"

	"school-ms/internal/middleware"
	"school-ms/internal/pkg/permcache"

	auth          "school-ms/internal/Modules/Auth"
	academicyears "school-ms/internal/Modules/AcademicYears"
	attendance    "school-ms/internal/Modules/Attendance"
	classes       "school-ms/internal/Modules/Classes"
	exams         "school-ms/internal/Modules/Exams"
	finance       "school-ms/internal/Modules/Finance"
	notices       "school-ms/internal/Modules/Notices"
	parents       "school-ms/internal/Modules/Parents"
	reporthandlers "school-ms/internal/Modules/Reports/Handlers"
	reportservices "school-ms/internal/Modules/Reports/Services"
	schools       "school-ms/internal/Modules/Schools"
	students      "school-ms/internal/Modules/Students"
	subjects      "school-ms/internal/Modules/Subjects"
	teachers      "school-ms/internal/Modules/Teachers"
	tenants       "school-ms/internal/Modules/Tenants"
	terms         "school-ms/internal/Modules/Terms"
	users         "school-ms/internal/Modules/Users"
	assignments    "school-ms/internal/Modules/Assignments"
   	discipline     "school-ms/internal/Modules/Discipline"
   	staffattendance "school-ms/internal/Modules/StaffAttendance"
   	timetable      "school-ms/internal/Modules/Timetable"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

func Setup(db *sqlx.DB, allowedOrigins []string) http.Handler {
	r := chi.NewRouter()

	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)

	r.Use(middleware.ResolveTenantFromDB(db))

	pc := permcache.New(db)
	r.Use(middleware.InjectPermCache(pc))

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

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	r.Route("/api/v1", func(r chi.Router) {
		auth.RegisterRoutes(r, db)
		users.RegisterRoutes(r, db)
		tenants.RegisterRoutes(r, db)

		schools.RegisterRoutes(r, db)
		academicyears.RegisterRoutes(r, db)
		terms.RegisterRoutes(r, db)
		classes.RegisterRoutes(r, db)
		subjects.RegisterRoutes(r, db)

		teachers.RegisterRoutes(r, db)
		students.RegisterRoutes(r, db)
		parents.RegisterRoutes(r, db)

		attendance.RegisterRoutes(r, db)
		exams.RegisterRoutes(r, db)
		finance.RegisterRoutes(r, db)
		notices.RegisterRoutes(r, db)
		assignments.RegisterRoutes(r, db)
   		discipline.RegisterRoutes(r, db)
   		staffattendance.RegisterRoutes(r, db)
   		timetable.RegisterRoutes(r, db)

		// Reports
		rptSvc := reportservices.NewReportService(db)
		rptH   := reporthandlers.NewReportHandler(rptSvc)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Authenticate)
			r.Route("/reports", func(r chi.Router) {
				r.With(middleware.RequirePermission(db, "reports.view")).
					Get("/report-card/{studentId}", rptH.ReportCard)
				r.With(middleware.RequirePermission(db, "reports.view")).
					Get("/class-results", rptH.ClassResults)
				r.With(middleware.RequirePermission(db, "reports.view")).
					Get("/fee-collection", rptH.FeeCollection)
				r.With(middleware.RequirePermission(db, "reports.view")).
					Get("/attendance-summary", rptH.AttendanceSummary)
				r.With(middleware.RequirePermission(db, "reports.view")).
					Get("/subject-performance", rptH.SubjectPerformance)
			})
		})
	})

	return r
}
