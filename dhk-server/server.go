package main

import "sync"

// Server holds all active sessions.
// TODO: sessions map value type depends on what you decide for your session/Key struct.
type Server struct {
	mu       sync.Mutex
	sessions map[string]interface{} // TODO: replace interface{} once you define your session type
}

func NewServer() *Server {
	return &Server{
		sessions: make(map[string]interface{}),
	}
}
