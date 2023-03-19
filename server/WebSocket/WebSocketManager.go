package WebSocketManager

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ConfigureWebsocket(server *fiber.App) {
	server.Use("/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			return ctx.Next()
		}
		return ctx.SendStatus(fiber.StatusUpgradeRequired)
	})

	server.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		fmt.Println("New Connection")
		var isCloseMessage = false

		defer func(conn *websocket.Conn) {
			var err = conn.Close()
			if err != nil {
				fmt.Println("Error when trying to close the connection...")
			}
		}(conn)
		conn.SetCloseHandler(func(code int, text string) error {
			isCloseMessage = true
			fmt.Println("Connection Closed. Code:", code, "Message:", text)
			return nil
		})

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				if !isCloseMessage {
					fmt.Println("Error when reading the message:", err)
				}
				break
			}
			fmt.Println("Received Message:", msg)

			if err = conn.WriteMessage(mt, msg); err != nil {
				fmt.Println("Error when sending message:", err)
				break
			}
		}
	}))
}
