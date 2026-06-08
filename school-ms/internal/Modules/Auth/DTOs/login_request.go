package dtos

// LoginRequest – the client sends only email + password.
// Tenant is resolved server-side from the HTTP Host header.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
