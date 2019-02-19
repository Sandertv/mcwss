// Package mcwss implements a websocket server to be used with Minecraft Bedrock Edition. It aims to implement
// most of the protocol and abstraction around it to achieve and easy to use API for interacting with the
// client.
//
// When using mcwss.NewServer(nil), the default configuration is used. Once the server is ran, (Server.Run())
// your client may connect to it using `/connect localhost:8000/ws'. The client will remain connected to the
// websocket until it either connects to a new one or closes the entire game, or when the websocket server
// chooses to forcibly close the connection.
package mcwss
