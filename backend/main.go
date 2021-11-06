package main

import (
	"log"
	"product-list/src/database"
	"product-list/src/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.DBConn = database.Connect()
}

func main() {
	app := fiber.New()
	routes.UserRoutes(app)

	log.Fatal(app.Listen(":8080"))
}