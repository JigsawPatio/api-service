package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// HTTPError represents an HTTP error
type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP error %d: %s", e.Code, e.Message)
}

// NewHTTPError returns a new HTTPError
func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{Code: code, Message: message}
}

// JSONError returns a JSON error response
func JSONError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

// JSONErrorf returns a JSON error response with a formatted message
func JSONErrorf(w http.ResponseWriter, code int, message string, a ...interface{}) {
	http.Error(w, fmt.Sprintf(message, a...), code)
}

// ValidateEmail checks if an email is valid
func ValidateEmail(email string) bool {
	return len(email) > 0 && strings.Contains(email, "@")
}

// ValidatePassword checks if a password is valid
func ValidatePassword(password string) bool {
	return len(password) >= 8
}

// DBError represents a database error
type DBError struct {
	gorm.Error
}

func (e *DBError) Error() string {
	return fmt.Sprintf("DB error: %s", e.Error)
}

// NewDBError returns a new DBError
func NewDBError(err error) *DBError {
	return &DBError{Error: err}
}

// GetLogger returns a logger instance
func GetLogger() *logrus.Entry {
	return logrus.NewEntry(logrus.New())
}

// GetDB returns a database instance
func GetDB() *gorm.DB {
	// TO DO: implement database connection
	return nil
}

// GetRouter returns a router instance
func GetRouter() *mux.Router {
	return mux.NewRouter()
}

// GetHTTPHandler returns an HTTP handler instance
func GetHTTPHandler() http.Handler {
	return http.DefaultServeMux
}

// GetTime returns the current time
func GetTime() time.Time {
	return time.Now()
}