# mcwss
A documentation and implementation of the Minecraft Bedrock Edition websocket protocol.

## Minecraft Websockets
Minecraft has websocket connection commands that may be used to connect to a websocket server. These commands are `/connect` and `/wsserver`.
Both these commands have the same effect. Once connected to a websocket, the connection will last until the server closes it or until the
player completely exits the game. The connection will last when a player merely leaves a world, and will thus last when joining a server.
Websockets in Minecraft can be used for simple command communication between a player and a server. Additionally, the server can choose to
listen for certain events sent by the client and act upon receiving them.

In order to be able to execute the `/connect` and `/wsserver` commands, cheats must be enabled. This means that connecting on third party
servers is not possible, and only on the dedicated server when cheats are enabled. It is possible to connect on a singleplayer world and
join a server after that, but commands executed by the websocket server will not work. Events will still be sent by the client however.

In the settings tab, there is also a setting that enables/disables encrypted websockets. mcwss implements encryption between websockets, 
so changing this setting will not have an effect on the behaviour of mcwss.

## Getting Started

### Prerequisites
mcwss is a Go library. To use it, Go must be installed. The library may be downloaded using `go get github.com/Sandertv/mcwss`.

### Usage
```go
package main

import (
	"github.com/sandertv/mcwss"
)

func main() {
    // Create a new server using the default configuration. To use specific configuration, pass a *wss.Config{} in here.
    server := mcwss.NewServer(nil)
    
    server.OnConnection(func(player *mcwss.Player){
      // Called when a player connects to the server.
    })
    server.OnDisconnection(func(player *mcwss.Player){
      // Called when a player disconnects from the server.
    })
    // Run the server. (blocking)
    server.Run()
}
```
The server may now be connected to by joining a singleplayer game and executing the command `/connect localhost:8000/ws`.

### Documentation
https://godoc.org/github.com/Sandertv/mcwss
