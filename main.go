package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"server-bof/api/routes"
	"server-bof/config"
	"server-bof/database"
)

func main()  {

	database.MysqlConnection()

	port, app := config.GetEnv("DEFAULT_PORT"), fiber.New()
	routes.Routes(app)
	//middlewares
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	err := app.Listen(port)
	if err != nil {
		return
	}

}
