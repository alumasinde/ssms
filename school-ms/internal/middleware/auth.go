package middleware

import (
	"context"
	"net/http"
	"strings"

	"school-ms/config"
	"school-ms/internal/pkg/permcache"
	"school-ms/internal/pkg/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)

type contextKey string

const (
	CtxUserID         contextKey = "user_id"
	CtxTenantID       contextKey = "tenant_id"
	CtxRoles          contextKey = "roles"
	CtxSchoolID       contextKey = "school_id"
	CtxAcademicYearID contextKey = "academic_year_id"
	CtxTermID         contextKey = "term_id"
	CtxTenantObj      contextKey = "tenant_obj"
	CtxTenantSlug     contextKey = "tenant_slug"
	CtxPermCache      contextKey = "perm_cache"
)

// Claims is embedded in every JWT token.
// All fields match the real schema and middleware helpers exactly.
type Claims struct {
	UserID         int64    `json:"user_id"`
	TenantID       int64    `json:"tenant_id"`
	SchoolID       *int64   `json:"school_id"`
	Roles          []string `json:"roles"`
	AcademicYearID *int64   `json:"academic_year_id,omitempty"`
	TermID         *int64   `json:"term_id,omitempty"`
	jwt.RegisteredClaims
}

// TenantInfo is the minimum tenant data passed through context.
type TenantInfo struct {
	ID   int64
	Slug string
	Name string
}

func InjectPermCache(pc *permcache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), CtxPermCache, pc)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ResolveTenantFromDB(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var tenant TenantInfo
			if fwdHost := r.Header.Get("X-Forwarded-Host"); fwdHost != "" {
				resolveTenantFromHost(db, extractHost(fwdHost), &tenant)
			}
			if tenant.ID == 0 {
				resolveTenantFromHost(db, extractHost(r.Host), &tenant)
			}
			ctx := context.WithValue(r.Context(), CtxTenantObj, tenant)
			ctx = context.WithValue(ctx, CtxTenantSlug, tenant.Slug)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func resolveTenantFromHost(db *sqlx.DB, host string, tenant *TenantInfo) {
	err := db.QueryRow(
		`SELECT id, name FROM tenants WHERE domain=? AND is_active=1 LIMIT 1`,
		host).Scan(&tenant.ID, &tenant.Name)
	if err == nil {
		return
	}
	slug := extractSlug(host)
	if slug != "" {
		db.QueryRow(
			`SELECT id, name FROM tenants WHERE domain LIKE ? AND is_active=1 LIMIT 1`,
			slug+"%").Scan(&tenant.ID, &tenant.Name)
	}
}

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
		ctx = context.WithValue(ctx, CtxRoles, claims.Roles)
		ctx = context.WithValue(ctx, CtxAcademicYearID, claims.AcademicYearID)
		ctx = context.WithValue(ctx, CtxTermID, claims.TermID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole guards a route to specific roles only.
func RequireRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRoles := GetRoles(r.Context())
			for _, allowed := range roles {
				for _, role := range userRoles {
					if role == allowed {
						next.ServeHTTP(w, r)
						return
					}
				}
			}
			response.Forbidden(w, "insufficient permissions")
		})
	}
}

// RequirePermission checks a permission via the in-process cache (5-min TTL).
// Falls back to a direct DB query when no cache is in context.
//
// DB fallback uses the real schema:
//   role_permissions(role_id, permission_id) → roles(id, name)
// The user may have multiple roles — any role granting the permission allows access.
func RequirePermission(db *sqlx.DB, permission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			roles := GetRoles(r.Context())
			var allowed bool

			if pc, ok := r.Context().Value(CtxPermCache).(*permcache.Cache); ok {
				for _, role := range roles {
					if pc.Has(role, permission) {
						allowed = true
						break
					}
				}
			} else {
				// Fallback: direct DB — schema-aligned query
				// roles table exists; role_permissions.role_id is a FK to roles.id
				var count int
				for _, role := range roles {
					db.QueryRow(`
						SELECT COUNT(*)
						FROM role_permissions rp
						INNER JOIN permissions p ON p.id  = rp.permission_id
						INNER JOIN roles       r ON r.id  = rp.role_id
						WHERE r.name  = ?
						  AND p.name  = ?
						  AND r.is_active = 1
					`, role, permission).Scan(&count)
					if count > 0 {
						allowed = true
						break
					}
				}
			}

			if !allowed {
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
func GetSchoolID(ctx context.Context) *int64 {
	v, _ := ctx.Value(CtxSchoolID).(*int64)
	return v
}
func GetRoles(ctx context.Context) []string {
	v, ok := ctx.Value(CtxRoles).([]string)
	if !ok {
		return []string{}
	}
	return v
}
func GetAcademicYearID(ctx context.Context) *int64 {
	v, _ := ctx.Value(CtxAcademicYearID).(*int64)
	return v
}
func GetTermID(ctx context.Context) *int64 {
	v, _ := ctx.Value(CtxTermID).(*int64)
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

func extractHost(hostHeader string) string {
	if i := strings.LastIndex(hostHeader, ":"); i != -1 {
		if !strings.Contains(hostHeader[i:], "]") {
			return hostHeader[:i]
		}
	}
	return hostHeader
}

func extractSlug(host string) string {
	if host == "" || host == "localhost" {
		return ""
	}
	parts := strings.Split(host, ".")
	n := len(parts)

	if parts[n-1] == "localhost" {
		if n >= 3 {
			return parts[n-2]
		}
		if n == 2 {
			return parts[0]
		}
		return ""
	}

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
	if n >= 3 {
		return parts[n-3]
	}
	if n == 2 {
		return parts[0]
	}
	return ""
}
