package middleware

import (
	"context"
	"net/http"
	"strings"

	"school-ms/config"
	"school-ms/internal/pkg/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

type contextKey string

const (
	CtxUserID     contextKey = "user_id"
	CtxTenantID   contextKey = "tenant_id"
	CtxRole       contextKey = "role"
	CtxSchoolID   contextKey = "school_id"
	CtxTenantObj  contextKey = "tenant_obj"
	CtxTenantSlug contextKey = "tenant_slug"
)

// Claims is embedded in every JWT token.
type Claims struct {
	UserID   int64  `json:"user_id"`
	TenantID int64  `json:"tenant_id"`
	SchoolID int64  `json:"school_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// TenantInfo is the minimum tenant data passed through context.
type TenantInfo struct {
	ID   int64
	Slug string
	Name string
}

//No slug in my table, use domain
func ResolveTenantFromDB(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var tenant TenantInfo
			

			// Priority 2: forwarded host (reverse proxy / nginx)
			if tenant.ID == 0 {
				if fwdHost := r.Header.Get("X-Forwarded-Host"); fwdHost != "" {
					host := extractHost(fwdHost)
					resolveTenantFromHost(db, host, &tenant)
				}
			}

			// Priority 3: actual Host header (direct subdomain access)
			if tenant.ID == 0 {
				host := extractHost(r.Host)
				resolveTenantFromHost(db, host, &tenant)
			}

			ctx := context.WithValue(r.Context(), CtxTenantObj, tenant)
			ctx = context.WithValue(ctx, CtxTenantSlug, tenant.Slug)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// resolveTenantFromHost tries exact domain match then slug match from subdomain.
func resolveTenantFromHost(db *sqlx.DB, host string, tenant *TenantInfo) {
	// Exact domain match against the full host
	err := db.QueryRow(
		`SELECT id, name FROM tenants WHERE domain=? AND is_active=1 LIMIT 1`,
		host).Scan(&tenant.ID, &tenant.Name)
	if err == nil {
		return
	}

	// Try the extracted slug portion as a domain prefix match
	slug := extractSlug(host)
	if slug != "" {
		db.QueryRow(
			`SELECT id, name FROM tenants WHERE domain LIKE ? AND is_active=1 LIMIT 1`,
			slug+"%").Scan(&tenant.ID, &tenant.Name)
	}
}

// ResolveTenant is a simple no-DB fallback kept for tests / environments
// where the DB isn't available during middleware setup.
func ResolveTenant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.Header.Get("X-Tenant-Slug")
		if slug == "" {
			slug = extractSlug(extractHost(r.Host))
		}
		ctx := context.WithValue(r.Context(), CtxTenantSlug, slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Authenticate validates the Bearer JWT and populates context with claims.
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Unauthorized(w, "missing or invalid authorization header")
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.App.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			response.Unauthorized(w, "invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserID, claims.UserID)
		ctx = context.WithValue(ctx, CtxTenantID, claims.TenantID)
		ctx = context.WithValue(ctx, CtxSchoolID, claims.SchoolID)
		ctx = context.WithValue(ctx, CtxRole, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole guards a route to specific roles only.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, _ := r.Context().Value(CtxRole).(string)
			for _, allowed := range roles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}
			response.Forbidden(w, "insufficient permissions")
		})
	}
}

// RequirePermission checks if the JWT role has a given permission via DB lookup
func RequirePermission(db *sqlx.DB, permission string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            role := GetRole(r.Context())
            var count int
            err := db.QueryRow(`
                SELECT COUNT(*) FROM role_permissions rp
                JOIN permissions p ON p.id = rp.permission_id
                WHERE rp.role = ? AND p.name = ?`, role, permission).Scan(&count)
            if err != nil || count == 0 {
                response.Forbidden(w, "you do not have permission to perform this action")
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}

// ── Context helpers ───────────────────────────────────────────────────────────

func GetUserID(ctx context.Context) int64 {
	v, _ := ctx.Value(CtxUserID).(int64)
	return v
}

func GetTenantID(ctx context.Context) int64 {
	v, _ := ctx.Value(CtxTenantID).(int64)
	return v
}

func GetSchoolID(ctx context.Context) int64 {
	v, _ := ctx.Value(CtxSchoolID).(int64)
	return v
}

func GetRole(ctx context.Context) string {
	v, _ := ctx.Value(CtxRole).(string)
	return v
}

func GetResolvedTenant(ctx context.Context) TenantInfo {
	v, _ := ctx.Value(CtxTenantObj).(TenantInfo)
	return v
}

func GetTenantSlug(ctx context.Context) string {
	v, _ := ctx.Value(CtxTenantSlug).(string)
	return v
}

// ── Internal helpers ──────────────────────────────────────────────────────────

// extractHost strips port from "host:port" → "host".
func extractHost(hostHeader string) string {
	if i := strings.LastIndex(hostHeader, ":"); i != -1 {
		// Don't strip if the colon is inside an IPv6 bracket
		if !strings.Contains(hostHeader[i:], "]") {
			return hostHeader[:i]
		}
	}
	return hostHeader
}

// extractSlug derives a tenant slug from a hostname.
//
// Supported patterns:
//   ssms.highwayhigh.ac.ke   →  "highwayhigh"
//   highwayhigh.ac.ke        →  "highwayhigh"
//   highwayhigh.com          →  "highwayhigh"
//   localhost / 127.0.0.1    →  ""
//
// Kenyan ccTLD second-level domains (.ac.ke, .co.ke, .or.ke, etc.) are treated
// specially so the school identifier is the label immediately before the SLD.
func extractSlug(host string) string {
	if host == "" || host == "localhost" {
		return ""
	}

	parts := strings.Split(host, ".")
	n := len(parts)

	// *.localhost TLD: ssms.highway.localhost → "highway"
	if parts[n-1] == "localhost" {
		if n >= 3 {
			return parts[n-2]
		}
		if n == 2 {
			return parts[0]
		}
		return ""
	}

	// Kenyan SLDs
	kenyanSLDs := map[string]bool{
		"ac.ke": true, "co.ke": true, "or.ke": true,
		"go.ke": true, "ne.ke": true, "sc.ke": true,
		"me.ke": true, "mobi.ke": true,
	}

	if n >= 2 {
		sld := parts[n-2] + "." + parts[n-1]
		if kenyanSLDs[sld] {
			if n >= 4 {
				return parts[n-3]
			}
			if n == 3 {
				return parts[0]
			}
		}
	}

	// Generic: foo.bar.com → "bar"
	if n >= 3 {
		return parts[n-3]
	}
	if n == 2 {
		return parts[0]
	}
	return ""
}