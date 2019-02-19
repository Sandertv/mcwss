package main

import (
	"github.com/sandertv/mcwss"
	"github.com/sandertv/mcwss/protocol/event"
	"log"
)

func main() {
	server := mcwss.NewServer(nil)
	server.OnConnection(func(player *mcwss.Player) {
		player.OnPlayerMessage(func(event *event.PlayerMessage) {

		})
		player.OnBlockBroken(func(event *event.BlockBroken) {
			player.ExecAs("say I broke a block", func(statusCode int) {

			})
		})
		player.OnBlockPlaced(func(event *event.BlockPlaced) {

		})
		player.OnItemUsed(func(event *event.ItemUsed) {

		})
		player.OnTravelled(func(event *event.PlayerTravelled) {

		})
	})

	if err := server.Run(); err != nil {
		log.Panicf("error running server: %v", err)
	}
}
