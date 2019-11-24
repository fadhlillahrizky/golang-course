package routes

import "net/http"

// Routes ...
type Routes []Route

// Route ...
type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

// SubRoutePackage ...
type SubRoutePackage struct {
	Routes 	Routes
	Middleware func(next http.Handler) http.Handler
}