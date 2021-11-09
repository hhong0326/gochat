package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func main() {

	fmt.Println("Chat Service")
	fmt.Println("migration test")

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	//create
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	type Message struct {
		ID     string `json:"id"`
		RoomID string `json:"room_id"`
		Text   string `json:"message"`
	}

	// make room
	server.On("/room", func(c *gosocketio.Channel, message Message) string {
		fmt.Println("Join: ", message.RoomID)
		c.Join(message.RoomID)

		return "room made successfully."
	})

	// leave room
	server.On("/leave", func(c *gosocketio.Channel, message Message) string {
		fmt.Println("Leave: ", message.RoomID)
		c.Leave(message.RoomID)

		return "leave room successfully."
	})

	// socket connection
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {

		fmt.Println("Connected: ", c.Id())
	})

	// socket disconnection
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		fmt.Println("Disconnected: ", c.Id())

		c.Leave("Room")
	})

	// chat socket
	server.On("/chat", func(c *gosocketio.Channel, payload Message) string {
		fmt.Println("Chat: ", payload.Text)

		c.BroadcastTo(payload.RoomID, "/message", payload.Text)
		return "message sent successfully."
	})

	router.GET("/socket.io/", gin.WrapH(server))
	router.StaticFile("/", "./client/home.htm")
	router.StaticFile("/app.js", "./client/app.js")

	fmt.Println("http://localhost:5001")

	if err := router.Run(":5001"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
