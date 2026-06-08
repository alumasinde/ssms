package dtos

type LoginResponse struct {
    AccessToken  string      `json:"access_token"`
    RefreshToken string      `json:"refresh_token"`
    User         UserSummary `json:"user"`
    Permissions  []string    `json:"permissions"`
}

type UserSummary struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	TenantID int64  `json:"tenant_id"`
	SchoolID int64  `json:"school_id"`
}
