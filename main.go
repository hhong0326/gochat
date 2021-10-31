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

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	//create
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	type Message struct {
		Text string `json:"message"`
	}

	// socket connection
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel, roomID Message) {

		fmt.Println("Connected", c.Id())
		c.Join(roomID.Text)

		// return
	})

	// socket disconnection
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		fmt.Println("Disconnected", c.Id())

		// handles when someone closes the tab
		c.Leave("Room")
	})

	// chat socket
	server.On("/chat", func(c *gosocketio.Channel, message Message) string {
		fmt.Println(message.Text)
		c.BroadcastTo("Room", "/message", message.Text)
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
