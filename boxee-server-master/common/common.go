package common

import (
	"log"
	"net/http"
)

// ReturnError is a handler to log error texts
// while returning a specified status code
func ReturnError(logger *log.Logger, text string, statusCode int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		logger.Printf(text)
	}
}
