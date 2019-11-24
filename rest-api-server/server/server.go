package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	
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

	server.HTTPServer = &http.Server{
		Addr:           server.Addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server
}
