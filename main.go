package main

import (
	"fmt"
	"github.com/FarukKaradeniz/SpaceHax-server/app/routes"
	"github.com/FarukKaradeniz/SpaceHax-server/config"
	"github.com/FarukKaradeniz/SpaceHax-server/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	routes.HaxAuthRoutes(app)
	routes.HaxGameRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.PORT)))
}
