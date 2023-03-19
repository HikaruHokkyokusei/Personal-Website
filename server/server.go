package main

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	envName    string
	portNumber string
)

func initialize() {
	fmt.Println("こんにちは　世界...")

	envName, portNumber = os.Getenv("EnvName"), os.Getenv("PORT")
	if envName == "" {
		envName = "prd"
	}
	if portNumber == "" {
		portNumber = "6969"
	}
}

func createServer() *fiber.App {
	var app = fiber.New(fiber.Config{})
	if envName == "dev" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173, http://localhost:5173/, " +
				"http://127.0.0.1:5173, http://127.0.0.1:5173/, " +
				"ws://localhost:5173, ws://localhost:5173/, " +
				"ws://127.0.0.1:5173, ws://127.0.0.1:5173/",
		}))
	}

	return app
}

func configureWebsocket(server *fiber.App) {
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

func configureServerRoutes(server *fiber.App) {
	server.Get("/healthCheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	server.Static("/", "./../build", fiber.Static{})
}

func main() {
	initialize()

	var server = createServer()
	configureWebsocket(server)
	configureServerRoutes(server)

	log.Fatal(server.Listen(":" + portNumber))
}
