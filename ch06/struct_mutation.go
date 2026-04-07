package main

import "fmt"

// Server represents a simple server configuration.
type Server struct {
	Host    string
	Port    int
	Running bool
}

// The pointer is copied, but it still points to the same struct.
// So changes affect the original value.
func updatePort(s *Server, newPort int) {
	s.Port = newPort // Go automatically dereferences the pointer
}

// Mutates the struct by changing its state.
func startServer(s *Server) {
	s.Running = true
	fmt.Printf("Server started: %s:%d\n", s.Host, s.Port)
}

// Safe print with nil check.
func printServer(s *Server) {
	if s == nil {
		fmt.Println("No server provided!")
		return
	}
	fmt.Printf("Host: %s | Port: %d | Running: %v\n", s.Host, s.Port, s.Running)
}

func main() {
	// Create struct and get pointer
	srv := &Server{
		Host:    "localhost",
		Port:    8080,
		Running: false,
	}

	printServer(srv)

	updatePort(srv, 9090)
	startServer(srv)

	printServer(srv)

	// Nil pointer example
	var emptySrv *Server
	printServer(emptySrv)
}
