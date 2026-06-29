package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	//"log/slog"


	"school-ms/config"
	"school-ms/internal/pkg/logger"
	"school-ms/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func main() {
	// Load config
	config.Load()
	cfg := config.App

/*
slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})))
	*/

	// ── Database ──────────────────────────────────────────────────────────────
	db, err := sqlx.Connect("mysql", cfg.DBDSN)
	if err != nil {
		logger.Error.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Connection pool tuning
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	// FIX: without these, stale connections cause "broken pipe" errors in prod
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	logger.Info.Println("database connected")

	// ── Router ────────────────────────────────────────────────────────────────
	origins := strings.Split(cfg.CORSAllowedOrigins, ",")
	handler := routes.Setup(db, origins)

	// ── HTTP Server ───────────────────────────────────────────────────────────
	addr := fmt.Sprintf(":%s", cfg.AppPort)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
		// FIX: without timeouts, slow/malicious clients (Slowloris) hold
		// connections open indefinitely, exhausting the goroutine pool.
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start server in a goroutine so we can listen for OS signals below
	go func() {
		logger.Info.Printf("%s listening on %s [%s]", cfg.AppName, addr, cfg.AppEnv)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("server error: %v", err)
		}
	}()

	// ── Graceful Shutdown ─────────────────────────────────────────────────────
	// FIX: previous code used http.ListenAndServe directly — a SIGTERM from
	// systemd would kill the process immediately, mid-write, risking data
	// corruption on active DB transactions.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info.Println("shutdown signal received — draining connections...")

	// Give active requests up to 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error.Printf("forced shutdown: %v", err)
	}

	logger.Info.Println("server stopped cleanly")
}
