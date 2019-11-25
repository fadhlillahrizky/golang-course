package router

import (
	"github.com/gorilla/mux"
	"net/http"

	Routes "rest-api-server/router/routes"
	V1Routes "rest-api-server/router/routes/v1"
)


const (
	staticDir = "/static/"
)

// RouterHandler ... the handler for go api routes
type RouterHandler struct {
	Router *mux.Router
}

// AttachSubRouterWithMiddleware ... allows yo to attach
// 	subroutes to router
func (r *RouterHandler) AttachSubRouterWithMiddleware(
	path string,
	subroutes Routes.Routes,
	Middleware mux.MiddlewareFunc,
) (SubRouter *mux.Router){
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(Middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	
	return
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

	v1SubRoutes := V1Routes.GetRoutes()
	for name, pack := range v1SubRoutes {
		router.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return &router


}
