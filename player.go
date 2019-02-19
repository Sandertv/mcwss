package mcwss

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sandertv/mcwss/protocol"
	"github.com/sandertv/mcwss/protocol/command"
	"github.com/sandertv/mcwss/protocol/event"
	"log"
	"reflect"
)

// Player is a player connected to the websocket server.
type Player struct {
	*websocket.Conn

	name      string
	connected bool
	event.Properties

	handlers         map[event.Name]func(event interface{})
	commandCallbacks map[string]reflect.Value
}

// Name returns the name of the player.
func (player *Player) Name() string {
	return player.name
}

// Connected checks if a player is currently connected. If not, the reference to this player should be
// released as soon as possible.
func (player *Player) Connected() bool {
	return player.connected
}

// Exec sends a command string with a callback that can process the output of the command. The callback passed
// must have one input argument, being the container of the output.
// This may be done using either a pointer to a struct, or a map, like so:
//
// player.Exec("getlocalplayername", func(response *command.LocalPlayerName){})
// player.Exec("getlocalplayername", func(response map[string]interface{}){})
func (player *Player) Exec(commandLine string, callback interface{}) {
	val := reflect.ValueOf(callback)
	t := val.Type()
	// Do some basic function validation.
	if t.Kind() != reflect.Func || t.NumIn() != 1 || (t.In(0).Kind() != reflect.Ptr && t.In(0).Kind() != reflect.Map) {
		panic("invalid callback type passed. must be of type func(*commandResponse)")
	}

	packet := protocol.NewCommandRequest(commandLine)
	player.commandCallbacks[packet.Header.RequestID] = val
	_ = player.WriteJSON(packet)
}

// ExecAs sends a command string as if it were sent by the player itself with a callback that can process the
// output of the command. The callback passed must have one input argument, being the container of the output.
// The output of the command is captured by the player, not by the websocket server. Only a status code is
// captured by the server that indicates if the command was successful.
func (player *Player) ExecAs(commandLine string, callback func(statusCode int)) {
	player.Exec(fmt.Sprintf("execute %v ~ ~ ~ %v", player.name, commandLine), func(response map[string]interface{}) {
		codeInterface, exists := response["statusCode"]
		if !exists {
			log.Printf("exec as: invalid response JSON")
			return
		}
		code, _ := codeInterface.(int)
		callback(code)
	})
}

// OnTransform subscribes to transformations done to the player, usually sent by means such as teleporting.
func (player *Player) OnTransform(handler func(event *event.PlayerTransform)) {
	_ = player.subscribeTo(event.NamePlayerTransform, func(e interface{}) {
		handler(e.(*event.PlayerTransform))
	})
}

// OnTravelled subscribes to the player travelling to places.
func (player *Player) OnTravelled(handler func(event *event.PlayerTravelled)) {
	_ = player.subscribeTo(event.NamePlayerTravelled, func(e interface{}) {
		handler(e.(*event.PlayerTravelled))
	})
}

// OnItemUsed subscribes to items used by the player.
func (player *Player) OnItemUsed(handler func(event *event.ItemUsed)) {
	_ = player.subscribeTo(event.NameItemUsed, func(e interface{}) {
		handler(e.(*event.ItemUsed))
	})
}

// OnBlockPlaced subscribes to blocks placed by the player.
func (player *Player) OnBlockPlaced(handler func(event *event.BlockPlaced)) {
	_ = player.subscribeTo(event.NameBlockPlaced, func(e interface{}) {
		handler(e.(*event.BlockPlaced))
	})
}

// OnBlockBroken subscribes to blocks broken by the player.
func (player *Player) OnBlockBroken(handler func(event *event.BlockBroken)) {
	_ = player.subscribeTo(event.NameBlockBroken, func(e interface{}) {
		handler(e.(*event.BlockBroken))
	})
}

// OnPlayerMessage subscribes to player messages sent and received by the client. Note that an event
// is called both when the player chats and when the player receives its own chat, resulting in a duplicate
// event when the player chats.
func (player *Player) OnPlayerMessage(handler func(event *event.PlayerMessage)) {
	_ = player.subscribeTo(event.NamePlayerMessage, func(e interface{}) {
		handler(e.(*event.PlayerMessage))
	})
}

// UnsubscribeFrom unsubscribes from events with the event name passed. The handler used to handle the event
// will no longer be called.
func (player *Player) UnsubscribeFrom(eventName event.Name) error {
	if err := player.WriteJSON(protocol.NewEventRequest(eventName, protocol.Unsubscribe)); err != nil {
		return fmt.Errorf("error writing event unsubscribe request: %v", err)
	}
	delete(player.handlers, eventName)
	return nil
}

// newPlayer returns an initialised player for a websocket connection.
func newPlayer(conn *websocket.Conn) *Player {
	player := &Player{
		connected:        true,
		Conn:             conn,
		handlers:         make(map[event.Name]func(event interface{})),
		commandCallbacks: make(map[string]reflect.Value),
	}
	player.Exec("getlocalplayername", func(response *command.LocalPlayerName) {
		player.name = response.LocalPlayerName
	})
	return player
}

// subscribeTo subscribes to an arbitrary event. It is recommended to use the methods to listen specifically
// to events above.
func (player *Player) subscribeTo(eventName event.Name, handler func(event interface{})) error {
	player.handlers[eventName] = handler
	if err := player.WriteJSON(protocol.NewEventRequest(eventName, protocol.Subscribe)); err != nil {
		return fmt.Errorf("error writing event subscribe request: %v", err)
	}
	return nil
}

// handleIncomingPacket handles an incoming packet, processing in particular the body of the packet.
func (player *Player) handleIncomingPacket(packet protocol.Packet) error {
	switch body := packet.Body.(type) {
	default:
		// Unknown or invalid packet. Don't try to process this.
		return fmt.Errorf("unknown packet %v", reflect.TypeOf(body).Name())
	case *protocol.ErrorResponse:
		return fmt.Errorf("a client side error occurred (code = %v): %v", body.StatusCode, body.StatusMessage)
	case *protocol.CommandResponse:
		callback, ok := player.commandCallbacks[packet.Header.RequestID]
		if !ok {
			return fmt.Errorf("command response: got command response with unknown requestID %v", packet.Header.RequestID)
		}
		// Remove the command callback from the map.
		delete(player.commandCallbacks, packet.Header.RequestID)

		commandResponseValue := reflect.New(callback.Type().In(0)).Interface()
		if err := json.Unmarshal([]byte(*body), commandResponseValue); err != nil {
			return fmt.Errorf("command response: malformed response JSON %v: %v", string(*body), err)
		}
		callback.Call([]reflect.Value{reflect.ValueOf(commandResponseValue).Elem()})
	case *protocol.EventResponse:
		properties := event.Properties{}
		if err := json.Unmarshal(body.Properties, &properties); err != nil {
			return fmt.Errorf("event response: malformed properties JSON: %v", err)
		}
		// Update the player's properties to the latest.
		player.Properties = properties

		eventData, ok := event.Events[body.EventName]
		if !ok {
			return fmt.Errorf("event response: unknown event with name %v", body.EventName)
		}
		_ = json.Unmarshal(body.Properties, &eventData)

		if measurable, ok := eventData.(event.Measurable); ok {
			// Parse measurements if the event requires them.
			measurable.ConsumeMeasurements(body.Measurements)
		}

		// Find the handler by the event name.
		handler, ok := player.handlers[body.EventName]
		if !ok {
			return fmt.Errorf("event response: unhandled event response for event %v", body.EventName)
		}
		// Finally call the handler with the event data processed.
		handler(eventData)
	}
	return nil
}
