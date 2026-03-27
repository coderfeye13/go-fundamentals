package main

import "fmt"

// Server represents a simple server configuration.
type Server struct {
	Host    string
	Port    int
	Running bool
}

// updatePort modifies the server via pointer.
// Without the *, this function would work on a COPY and changes would be lost.
func updatePort(s *Server, newPort int) {
	s.Port = newPort // Go auto-dereferences: this is (*s).Port = newPort
}

func startServer(s *Server) {
	s.Running = true
	fmt.Printf("Server started: %s:%d\n", s.Host, s.Port)
}

func printServer(s *Server) {
	// Nil check — defensive programming
	if s == nil {
		fmt.Println("No server provided!")
		return
	}
	fmt.Printf("Host: %s | Port: %d | Running: %v\n", s.Host, s.Port, s.Running)
}

func main() {
	// & before struct literal → returns a pointer directly
	srv := &Server{
		Host:    "localhost",
		Port:    8080,
		Running: false,
	}

	printServer(srv) // Host: localhost | Port: 8080 | Running: false

	updatePort(srv, 9090)
	startServer(srv)

	printServer(srv) // Host: localhost | Port: 9090 | Running: true

	// Nil pointer example — safe because printServer checks for nil
	var emptySrv *Server
	printServer(emptySrv) // No server provided!
}
