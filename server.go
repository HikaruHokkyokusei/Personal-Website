package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	var envName, port = os.Getenv("EnvName"), os.Getenv("PORT")
	if envName == "" {
		envName = "prd"
	}
	if port == "" {
		port = "6969"
	}

	fmt.Println("こんにちは　世界...")

	var app = fiber.New(fiber.Config{})

	if envName == "dev" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173, http://localhost:5173/, " +
				"http://127.0.0.1:5173, http://127.0.0.1:5173/, " +
				"ws://localhost:5173, ws://localhost:5173/, " +
				"ws://127.0.0.1:5173, ws://127.0.0.1:5173/",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	app.Use("/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(conn *websocket.Conn) {
		fmt.Println("New Connection:", conn.Params("id"))
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("Error when trying to close the connection...")
			}
		}(conn)

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = conn.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))
	app.Get("/healthCheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	app.Static("/", "./build", fiber.Static{})

	log.Fatal(app.Listen(":" + port))
}
