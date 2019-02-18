package mcwss

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sandertv/mcwss/protocol"
	"github.com/sandertv/mcwss/protocol/event"
	"reflect"
)

// Player is a player connected to the websocket server.
type Player struct {
	*websocket.Conn
	event.Properties
	handlers map[event.Name]func(event interface{})
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
