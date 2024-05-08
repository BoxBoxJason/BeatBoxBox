package custom_errors

import "net/http"

// BadRequestError represents a 400 error
type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{Message: message}
}

// FileTooBigError represents a 400 error
type FileTooBigError struct {
	Message string
}

func (e *FileTooBigError) Error() string {
	return e.Message
}

func NewFileTooBigError(message string) *FileTooBigError {
	return &FileTooBigError{Message: message}
}

// UnauthorizedError represents a 401 error
type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{Message: message}
}

// NotFoundError represents a 404 error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

// DatabaseError represents a 500 error, used for database errors (failed to connect, etc.)
type DatabaseError struct {
	Message string
}

func NewDatabaseError(message string) *DatabaseError {
	return &DatabaseError{Message: message}
}

func (e *DatabaseError) Error() string {
	return e.Message
}

func SendErrorToClient(err error, w http.ResponseWriter, default_error_msg string) {
	switch err.(type) {
	case *BadRequestError, *FileTooBigError:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case *UnauthorizedError:
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case *NotFoundError:
		http.Error(w, err.Error(), http.StatusNotFound)
	case *DatabaseError:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	default:
		if default_error_msg != "" {
			http.Error(w, default_error_msg, http.StatusInternalServerError)
		} else {
			http.Error(w, "Unexpected error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
