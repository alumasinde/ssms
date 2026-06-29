// Package ratelimit provides lightweight in-process IP-based rate limiting
// using a token-bucket algorithm. No external dependencies required.
package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"school-ms/internal/pkg/response"
)

type bucket struct {
	tokens    float64
	lastRefil time.Time
}

// Limiter holds per-IP token buckets.
type Limiter struct {
	mu       sync.Mutex
	buckets  map[string]*bucket
	rate     float64 // tokens per second
	capacity float64 // max tokens (= burst size)
}

// New creates a Limiter.
//   - requests: allowed requests per window
//   - window:   rolling window duration (e.g. time.Minute)
//
// Example: New(10, time.Minute) → 10 req/min per IP
func New(requests int, window time.Duration) *Limiter {
	rate := float64(requests) / window.Seconds()
	l := &Limiter{
		buckets:  make(map[string]*bucket),
		rate:     rate,
		capacity: float64(requests),
	}
	// Background sweep — remove stale buckets every 5 minutes to prevent
	// memory growth on high-traffic servers.
	go l.sweep()
	return l
}

// Allow returns true if the IP is within the rate limit.
func (l *Limiter) Allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	b, ok := l.buckets[ip]
	if !ok {
		l.buckets[ip] = &bucket{tokens: l.capacity - 1, lastRefil: time.Now()}
		return true
	}

	// Refill tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(b.lastRefil).Seconds()
	b.tokens += elapsed * l.rate
	if b.tokens > l.capacity {
		b.tokens = l.capacity
	}
	b.lastRefil = now

	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

// Middleware returns an http.Handler middleware that limits requests by IP.
// On limit exceeded it responds with 429 and a safe generic message.
func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := realIP(r)
		if !l.Allow(ip) {
			response.TooManyRequests(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// realIP extracts the client IP, honoring X-Real-IP and X-Forwarded-For
// (chi's RealIP middleware normalises r.RemoteAddr, so this is a fallback).
func realIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		// X-Forwarded-For may be "client, proxy1, proxy2" — take first
		for i := 0; i < len(ip); i++ {
			if ip[i] == ',' {
				return ip[:i]
			}
		}
		return ip
	}
	// chi RealIP middleware already strips the port from RemoteAddr
	return r.RemoteAddr
}

// sweep removes stale buckets (no activity for >10 minutes) every 5 minutes.
func (l *Limiter) sweep() {
	t := time.NewTicker(5 * time.Minute)
	defer t.Stop()
	for range t.C {
		cutoff := time.Now().Add(-10 * time.Minute)
		l.mu.Lock()
		for ip, b := range l.buckets {
			if b.lastRefil.Before(cutoff) {
				delete(l.buckets, ip)
			}
		}
		l.mu.Unlock()
	}
}
