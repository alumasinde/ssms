package dtos

// LoginResponse is returned on a successful POST /auth/login.
type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    int64       `json:"expires_in"` // seconds

	User        UserSummary `json:"user"`
	Roles       []string    `json:"roles"`
	Permissions []string    `json:"permissions"`
	Context     AuthContext `json:"context"`
}

// UserSummary is the safe user projection embedded in LoginResponse.
// Password hash and internal audit fields are never exposed.
type UserSummary struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string `json:"name"`       // computed: first_name + " " + last_name
	Email     string `json:"email"`
	TenantID  int64  `json:"tenant_id"`
	SchoolID  *int64 `json:"school_id,omitempty"` // omitted for superadmin
	IsActive  bool   `json:"is_active"`
}

// AuthContext carries the current academic year and term IDs so the frontend
// never needs a separate round-trip immediately after login.
type AuthContext struct {
	AcademicYearID *int64 `json:"academic_year_id,omitempty"`
	TermID         *int64 `json:"term_id,omitempty"`
}

// RefreshRequest is the body for POST /auth/refresh.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
