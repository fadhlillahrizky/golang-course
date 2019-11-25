package v1

import (
	"net/http"
	"log"

	Routes "rest-api-server/router/routes"
)

// Middleware ... this is middleware fot v1 routes
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Println("hit V1 middleware")
		next.ServeHTTP(w, r)
	})
}  
	
// GetRoutes ... get v1 routes
func GetRoutes() (SubRoute map[string]Routes.SubRoutePackage){
	SubRoute = map[string]Routes.SubRoutePackage {
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name: "V1HealthRoute", 
					Method: "GET", 
					Pattern: "/health", 
					HandlerFunc: Health(),
				},
			},
			Middleware: Middleware,
		},
	}
	return SubRoute
}