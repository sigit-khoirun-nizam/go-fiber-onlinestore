package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sigit-khoirun-nizam/go-fiber-onlinestore/app/routes"
	"github.com/sigit-khoirun-nizam/go-fiber-onlinestore/app/utils"
)

func main() {
	app := fiber.New()
	utils.InitDatabase()
	defer utils.DB.Close()

	routes.SetupRoutes(app)

	port := 3000
	fmt.Printf("Server started on :%d\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
