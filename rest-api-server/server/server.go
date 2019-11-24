package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	RouterFactory "rest-api-server/router"

	
)

// Server - thi is the server object
type Server struct {
	Port       int
	Addr       string
	HTTPServer *http.Server
}

// Start - starts the server service
func (s *Server) Start() {
	
	log.Println("Server started on port", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())

}

// NewServer - create a new server
func NewServer(port int) *Server {
	var server Server

	server.Port = port
	server.Addr = ":" + strconv.Itoa(port)

	router := RouterFactory.NewRouter()

	server.HTTPServer = &http.Server{
		Addr:           server.Addr,
		Handler:        router.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server
}
