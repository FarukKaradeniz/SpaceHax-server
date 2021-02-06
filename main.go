package main

import (
	"fmt"
	"github.com/FarukKaradeniz/SpaceHax-server/app/routes"
	"github.com/FarukKaradeniz/SpaceHax-server/config"
	"github.com/FarukKaradeniz/SpaceHax-server/config/database"
	"github.com/FarukKaradeniz/SpaceHax-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	database.Connect()
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(cors.New())

	routes.HaxAuthRoutes(app)
	routes.HaxGameRoutes(app)
	routes.AdminRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.PORT)))
}
