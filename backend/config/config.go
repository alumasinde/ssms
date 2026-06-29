package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
	AppName string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBDSN  string

	JWTSecret     string
	JWTAccessTTL  time.Duration
	JWTRefreshTTL time.Duration

	CORSAllowedOrigins string
}

var App *Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("[config] .env not found, reading from environment")
	}

	accessTTL, err := time.ParseDuration(getEnv("JWT_ACCESS_TTL", "15m"))
	if err != nil {
		accessTTL = 15 * time.Minute
	}
	refreshTTL, err := time.ParseDuration(getEnv("JWT_REFRESH_TTL", "168h"))
	if err != nil {
		refreshTTL = 7 * 24 * time.Hour
	}

	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "")
	name := getEnv("DB_NAME", "school_ms")

	App = &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
		AppName: getEnv("APP_NAME", "SchoolMS"),

		DBHost: host,
		DBPort: port,
		DBUser: user,
		DBPass: pass,
		DBName: name,
		DBDSN:  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pass, host, port, name),

		JWTSecret:     getEnv("JWT_SECRET", "changeme"),
		JWTAccessTTL:  accessTTL,
		JWTRefreshTTL: refreshTTL,

		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
