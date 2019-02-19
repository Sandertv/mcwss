package main

import (
	"bufio"
	"fmt"
	"github.com/sandertv/mcwss"
	"github.com/sandertv/mcwss/protocol/event"
	"log"
	"os"
)

func main() {
	server := mcwss.NewServer(nil)
	server.OnConnection(func(player *mcwss.Player) {
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() && player.Connected() {
				player.Exec(scanner.Text(), func(data map[string]interface{}) {
					fmt.Println(data)
				})
			}
		}()
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
