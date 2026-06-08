package main

import (
	"fmt"
	"net/http"
	"strings"

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

	// Database
	db, err := sqlx.Connect("mysql", cfg.DBDSN)
	if err != nil {
		logger.Error.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	logger.Info.Println("database connected")

	// Router
	origins := strings.Split(cfg.CORSAllowedOrigins, ",")
	handler := routes.Setup(db, origins)

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	logger.Info.Printf("%s listening on %s [%s]", cfg.AppName, addr, cfg.AppEnv)

	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Error.Fatalf("server error: %v", err)
	}
}
