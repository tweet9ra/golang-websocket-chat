package chat

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
	"log"
)

// Chat server.
type Server struct {
	pattern   string
	clients   map[uint]*Client
	anonymousClients map[uint]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *MessageFromUser
	doneCh    chan bool
	errCh     chan error
}

// Create new chat server.
func NewServer(pattern string) *Server {
	clients := make(map[uint]*Client)
	anonymousClients := make(map[uint]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *MessageFromUser)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		clients,
		anonymousClients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

// Add new client to as anonymous
func (s *Server) Add(c *Client) {
	s.addCh <- c
}

// Delete anonymous/authorized client
func (s *Server) Del(c *Client) {
	s.delCh <- c
}

// Stop server
func (s *Server) Done() {
	s.doneCh <- true
}

// Server error
func (s *Server) Err(err error) {
	s.errCh <- err
}

// Send ws message to clients
func (s *Server) SendToClients(message *MessageToUser, receivers []uint) {
	for _, receiverId := range receivers {
		for _, user := range s.clients {
			if user.id == receiverId {
				user.Write(message)
			}
		}
	}
}

// Move from anonymous users to authorized users
func (s *Server) AuthorizeClient(anonymousId uint, c *Client) {
	s.clients[c.id] = c
	delete(s.anonymousClients, anonymousId)
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen(router *mux.Router) {
	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
	router.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			s.anonymousClients[c.id] = c

		// del a client
		case c := <-s.delCh:
			if c.authorized {
				delete(s.clients, c.id)
			} else {
				delete(s.anonymousClients, c.id)
			}

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
