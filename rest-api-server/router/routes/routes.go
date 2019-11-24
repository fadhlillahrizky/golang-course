package routes

import (
	"net/http"
)

// Middleware ... this is the main middlerware for the application
func Middleware(next http.Handler) http.Handler  {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// log.Println("inside middleware")s
		next.ServeHTTP(w, r)
	})
}

// GetRoutes ... 
func GetRoutes() Routes {
	
	return Routes {
		Route{Name: "HealthChack", Method:"GET", Pattern: "/health", HandlerFunc: Health()},
	}

}