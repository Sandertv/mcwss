package main

import (
	"fmt"
	"github.com/sandertv/mcwss"
	"github.com/sandertv/mcwss/protocol/event"
	"log"
)

func main() {
	server := mcwss.NewServer(nil)
	server.OnConnection(func(player *mcwss.Player) {
		player.OnPlayerMessage(func(event *event.PlayerMessage) {
			fmt.Printf("%v sent a message: %v (%v)\n", event.Sender, event.Message, event.MessageType)
		})
		player.OnBlockPlace(func(event *event.BlockPlaced) {

		})
	})

	if err := server.Run(); err != nil {
		log.Panicf("error running server: %v", err)
	}
}
