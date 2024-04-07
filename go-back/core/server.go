package core

import (
	"encoding/json"
	"strconv"
)

type Server struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewServer() *Server {
	return &Server{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (s *Server) Run() {

	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
		case message := <-s.broadcast:
			var result map[string]string
			_ = json.Unmarshal(message, &result)
			for client := range s.clients {
				userid, _ := strconv.Atoi(result["userid"])
				if userid == client.userid {
					select {
					case client.send <- []byte(result["message"]):
					default:
						close(client.send)
						delete(s.clients, client)
					}
				}
			}
		}
	}
}
