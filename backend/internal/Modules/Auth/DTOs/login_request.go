package dtos

import "strings"

// LoginRequest is the JSON body for POST /auth/login.
// Tenant is resolved server-side from Host — never sent by the client.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Normalize lower-cases and trims both fields in-place.
func (r *LoginRequest) Normalize() {
	r.Email    = strings.TrimSpace(strings.ToLower(r.Email))
	r.Password = strings.TrimSpace(r.Password)
}

// Validate returns a map of field → message when invalid, nil when valid.
func (r *LoginRequest) Validate() map[string]string {
	errs := map[string]string{}
	if r.Email == "" {
		errs["email"] = "email is required"
	}
	if r.Password == "" {
		errs["password"] = "password is required"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
