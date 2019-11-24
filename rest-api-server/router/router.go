package router

import (
	"github.com/gorilla/mux"
	"net/http"

	Routes "rest-api-server/router/routes"
)


const (
	staticDir = "/static/"
)

// RouterHandler ... the handler for go api routes
type RouterHandler struct {
	Router *mux.Router
}

// NewRouter ... create a new router
func NewRouter() *RouterHandler {
	var router RouterHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	
	router.Router.Use(Routes.Middleware)

	routes := Routes.GetRoutes()

	for _, route := range routes {
		router.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return &router


}
