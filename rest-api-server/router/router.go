package router

import (
	"github.com/gorilla/mux"
	"net/http"
)


const (
	staticDir = "/static/"
)

// RouterHandler - the handler for go api routes
type RouterHandler struct {
	Router *mux.Router
}

// NewRouter - create a new router
func NewRouter() *RouterHandler {
	var router RouterHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	
	
	return &router


}
