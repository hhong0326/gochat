package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ConnectChatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		r_id := c.Param("room_id")

		fmt.Println(r_id)

		// push

		c.JSON(200, gin.H{
			"message": "Joined the chat",
		})
	}
}
