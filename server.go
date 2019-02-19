package mcwss

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sandertv/mcwss/protocol"
	"log"
	"net/http"
)

// Server is the main entry-point of the mcwss package. It allows interfacing with clients connected to it and
// provides methods to ease the use.
type Server struct {
	config   Config
	upgrader websocket.Upgrader

	players           map[*Player]bool
	connectionFunc    func(conn *Player)
	disconnectionFunc func(conn *Player)
}

// NewServer initialises and returns a new server. If the configuration passed in was non-nil, that
// configuration is used. If nil is passed, the default configuration is used. (see config.go/defaultConfig())
func NewServer(config *Config) *Server {
	server := &Server{
		players:           make(map[*Player]bool),
		config:            defaultConfig(),
		connectionFunc:    func(conn *Player) {},
		disconnectionFunc: func(conn *Player) {},
	}
	if config != nil {
		server.config = *config
	}
	return server
}

// Run runs the server and blocks the current goroutine until the server is stopped.
func (server *Server) Run() error {
	http.HandleFunc(server.config.HandlerPattern, server.handleResponse)
	if err := http.ListenAndServe(server.config.Address, nil); err != nil {
		return err
	}
	return nil
}

// OnConnection sets a handler to handle new connections from players. This method must be used to interact
// with players connected.
func (server *Server) OnConnection(handler func(player *Player)) {
	server.connectionFunc = handler
}

// OnDisconnection sets a handler to handle disconnections from players. Note that when setting the handler,
// sending packets to the player will not arrive.
func (server *Server) OnDisconnection(handler func(player *Player)) {
	server.disconnectionFunc = handler
}

// handleResponse handles the websocket response of a client connecting to the server. It first initialises
// the websocket connection, after which it starts processing and sending packets.
func (server *Server) handleResponse(writer http.ResponseWriter, request *http.Request) {
	ws, err := server.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Panicf("error upgrading request: %v", err)
	}

	// Initialise the player and add it to the players map.
	player := NewPlayer(ws)
	server.players[player] = true

	// Allow the creator of the server to interact with the new player.
	server.connectionFunc(player)

	defer func() {
		server.disconnectionFunc(player)
		if err := ws.Close(); err != nil {
			log.Panicf("error closing websocket connection: %v", err)
		}
		delete(server.players, player)
		player.connected = false
	}()

	for {
		msgType, payload, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error reading message from connection: %v", err)
			break
		}
		if msgType != 1 {
			log.Printf("unexpected message type %v", msgType)
			break
		}
		packet := &protocol.Packet{}
		if err := json.Unmarshal(payload, packet); err != nil {
			log.Printf("malformed packet JSON: %v", err)
			break
		}
		// Find the correct body packet for the message purpose.
		body, found := protocol.Packets[packet.Header.MessagePurpose]
		if !found {
			log.Printf("unknown message purpose %v", packet.Header.MessagePurpose)
			break
		}

		// The packet.Body is currently a map because of some funny JSON behaviour. We marshal and unmarshal
		// it into a pointer and set it back to have a pointer to a struct as body.
		data, _ := json.Marshal(packet.Body)
		cmdResponse, ok := body.(*protocol.CommandResponse)

		if !ok {
			if err := json.Unmarshal(data, &body); err != nil {
				log.Printf("map to struct conversion failed: %v", err)
				break
			}
			fmt.Println(string(payload))
		} else {
			*cmdResponse = protocol.CommandResponse(data)
		}
		packet.Body = body

		if err := player.handleIncomingPacket(*packet); err != nil {
			log.Printf("%v (payload: %v)", err, string(payload))
			break
		}
	}
}
