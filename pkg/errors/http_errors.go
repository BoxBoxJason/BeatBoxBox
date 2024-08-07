package custom_errors

import "net/http"

type HTTPError interface {
	Error() string
	StatusCode() int
}

// BadRequestError represents a 400 error
type BadRequestError struct{ Message string }

func (e *BadRequestError) Error() string                 { return e.Message }
func NewBadRequestError(message string) *BadRequestError { return &BadRequestError{Message: message} }
func (e *BadRequestError) StatusCode() int               { return http.StatusBadRequest }

// FileTooBigError represents a 400 error
type FileTooBigError struct{ Message string }

func (e *FileTooBigError) Error() string                 { return e.Message }
func NewFileTooBigError(message string) *FileTooBigError { return &FileTooBigError{Message: message} }
func (e *FileTooBigError) StatusCode() int               { return http.StatusRequestEntityTooLarge }

// UnauthorizedError represents a 401 error
type UnauthorizedError struct{ Message string }

func (e *UnauthorizedError) Error() string { return e.Message }
func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{Message: message}
}
func (e *UnauthorizedError) StatusCode() int { return http.StatusUnauthorized }

// ForbiddenError represents a 403 error
type ForbiddenError struct{ Message string }

func (e *ForbiddenError) Error() string                { return e.Message }
func NewForbiddenError(message string) *ForbiddenError { return &ForbiddenError{Message: message} }
func (e *ForbiddenError) StatusCode() int              { return http.StatusForbidden }

// NotFoundError represents a 404 error
type NotFoundError struct{ Message string }

func (e *NotFoundError) Error() string               { return e.Message }
func NewNotFoundError(message string) *NotFoundError { return &NotFoundError{Message: message} }
func (e *NotFoundError) StatusCode() int             { return http.StatusNotFound }

// DatabaseError represents a 500 error, used for database errors (failed to connect, etc.)
type DatabaseError struct{ Message string }

func NewDatabaseError(message string) *DatabaseError { return &DatabaseError{Message: message} }
func (e *DatabaseError) Error() string               { return e.Message }
func (e *DatabaseError) StatusCode() int             { return http.StatusInternalServerError }

// InternalServerError represents a 500 error
type InternalServerError struct{ Message string }

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{Message: message}
}
func (e *InternalServerError) Error() string   { return e.Message }
func (e *InternalServerError) StatusCode() int { return http.StatusInternalServerError }

// ConflictError represents a 409 error
type ConflictError struct{ Message string }

func NewConflictError(message string) *ConflictError { return &ConflictError{Message: message} }
func (e *ConflictError) Error() string               { return e.Message }
func (e *ConflictError) StatusCode() int             { return http.StatusConflict }

func SendErrorToClient(w http.ResponseWriter, err error) {
	if http_err, ok := err.(HTTPError); ok {
		http.Error(w, http_err.Error(), http_err.StatusCode())
	} else {
		http.Error(w, http_err.Error(), http.StatusInternalServerError)
	}
}
