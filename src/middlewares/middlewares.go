package middlewares

import (
	"devbook/src/responses"
	"devbook/src/security"
	"fmt"
	"net/http"
)


// LogRequest logs ai requests
func LogRequest (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// CheckAuthenticatedRequest check if a user token is valid for authenticated request
func CheckAuthenticatedRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := security.ValidateToken(r); error != nil{
			responses.ErrorResponse(w, http.StatusUnauthorized, error)
			return
		}
		next(w, r)
	}
}