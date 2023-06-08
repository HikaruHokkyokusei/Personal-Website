package main

import (
	"fmt"
	"log"
	"os"
	re "regexp"
	"time"

	WS "Server/WebSocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var (
	envName         string
	portNumber      string
	allowedOrigins  string
	enableRouteLogs string
)

func getEnv(key string, defaultValue string) string {
	if value, isPresent := os.LookupEnv(key); isPresent {
		return value
	} else {
		return defaultValue
	}
}

func createFiberApp() *fiber.App {
	var app = fiber.New(fiber.Config{
		CaseSensitive: true,
		//DisableDefaultContentType: true,
		//EnablePrintRoutes: true,
	})
	if enableRouteLogs == "true" {
		app.Use(logger.New())
	}

	if envName == "dev" {
		allowedOrigins = "http://localhost:5173, https://localhost:5173, http://127.0.0.1:5173, https://127.0.0.1:5173"
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

func configureFiberApp(app *fiber.App) {
	app.Get("/healthCheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	})
	app.Get("/metrics", monitor.New())

	WS.ConfigureWebsocket(app)

	app.Static("/", "./../build", fiber.Static{
		Index: "index.html",
	})
	app.Static("/*", "./../build", fiber.Static{
		Index: "index.html",
	})
}

func init() {
	fmt.Println("こんにちは　世界...")
	envName = getEnv("EnvName", "prd")
	portNumber = getEnv("PORT", "6969")
	allowedOrigins = getEnv("AllowedOrigins", "")
	enableRouteLogs = getEnv("enableRouteLogs", "false")
}

func main() {
	var fiberApp = createFiberApp()
	configureFiberApp(fiberApp)

	log.Fatal(fiberApp.Listen(":" + portNumber))
}
