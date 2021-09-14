package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Heck you 👋!")
	})

	app.Get("/callback", handleCallback())
}

func main() {
	envErr := godotenv.Load()
	app := fiber.New()

	if envErr != nil {
		log.Fatal("Error loding env")
	}

	setupRoutes(app)
	initDB()
	getToken()

  fmt.Println("Heck you world!")
  log.Fatal(app.Listen(":3000"))
}



