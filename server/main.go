package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	WS "Server/WebSocket"
)

var (
	envName    string
	portNumber string
)

func getEnv(key string, defaultValue string) string {
	if value, isPresent := os.LookupEnv(key); isPresent {
		return value
	} else {
		return defaultValue
	}
}

func initialize() {
	fmt.Println("こんにちは　世界...")
	envName = getEnv("EnvName", "prd")
	portNumber = getEnv("PORT", "6969")
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

func configureServerRoutes(server *fiber.App) {
	server.Get("/healthCheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	})

	server.Static("/", "./../build", fiber.Static{})
}

func main() {
	initialize()

	var server = createServer()
	WS.ConfigureWebsocket(server)
	configureServerRoutes(server)

	log.Fatal(server.Listen(":" + portNumber))
}