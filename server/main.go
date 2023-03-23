package main

import (
	"fmt"
	"log"
	"os"
	re "regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"

	WS "Server/WebSocket"
)

var (
	envName        string
	portNumber     string
	allowedOrigins string
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
	allowedOrigins = getEnv("AllowedOrigins", "")
}

func createServer() *fiber.App {
	var app = fiber.New(fiber.Config{})

	if envName == "dev" {
		allowedOrigins = "http://localhost:5173, https://localhost:5173, ws://localhost:5173, wss://localhost:5173"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
	}))

	app.Use(cache.New(cache.Config{
		Next: func(ctx *fiber.Ctx) bool {
			if path := ctx.Path(); re.MustCompile(`.*/ws(/.*)?$`).MatchString(path) {
				return true
			}
			return false
		},
		CacheControl: true,
		Expiration:   24 * time.Hour,
		ExpirationGenerator: func(ctx *fiber.Ctx, config *cache.Config) time.Duration {
			// Redundant function kept just to remind of it's existence
			return config.Expiration
		},
		Methods: []string{fiber.MethodGet},
	}))

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
