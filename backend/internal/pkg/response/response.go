package response

import (
	"encoding/json"
	"net/http"

	"school-ms/internal/pkg/logger"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func JSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func Success(w http.ResponseWriter, data interface{}, message string) {
	JSON(w, http.StatusOK, APIResponse{Success: true, Message: message, Data: data})
}

func Created(w http.ResponseWriter, data interface{}, message string) {
	JSON(w, http.StatusCreated, APIResponse{Success: true, Message: message, Data: data})
}

func BadRequest(w http.ResponseWriter, err string) {
	JSON(w, http.StatusBadRequest, APIResponse{Success: false, Error: err})
}

func Unauthorized(w http.ResponseWriter, err string) {
	JSON(w, http.StatusUnauthorized, APIResponse{Success: false, Error: err})
}

func Forbidden(w http.ResponseWriter, err string) {
	JSON(w, http.StatusForbidden, APIResponse{Success: false, Error: err})
}

func NotFound(w http.ResponseWriter, err string) {
	JSON(w, http.StatusNotFound, APIResponse{Success: false, Error: err})
}

func TooManyRequests(w http.ResponseWriter) {
	JSON(w, http.StatusTooManyRequests, APIResponse{
		Success: false,
		Error:   "too many requests — please wait and try again",
	})
}

func Paginated(w http.ResponseWriter, data interface{}, meta Meta) {
	JSON(w, http.StatusOK, APIResponse{Success: true, Data: data, Meta: &meta})
}

// ServerError logs the real error internally and returns a safe generic
// message to the client. Use this everywhere instead of InternalError.
// Never call err.Error() in a response — it exposes DB schema and internals.
func ServerError(w http.ResponseWriter, err error) {
	logger.Error.Printf("internal server error: %v", err)
	JSON(w, http.StatusInternalServerError, APIResponse{
		Success: false,
		Error:   "an unexpected error occurred, please try again",
	})
}

// InternalError is kept for backward compatibility but now also logs safely.
// Prefer ServerError(w, err) for new code.
// IMPORTANT: this no longer forwards err.Error() to the client.
func InternalError(w http.ResponseWriter, _ string) {
	JSON(w, http.StatusInternalServerError, APIResponse{
		Success: false,
		Error:   "an unexpected error occurred, please try again",
	})
}
