package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	log.Fatal(app.Listen(":3000"))
}