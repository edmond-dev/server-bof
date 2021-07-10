package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"server-bof/api/routes"
	"server-bof/config"
	"server-bof/database"
)

func main() {

	database.MysqlConnection()
	port := config.GetEnv("DEFAULT_PORT")
	app := fiber.New()

	//CORS middleware config
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST, GET, DELETE, PATCH, PUT",
		AllowHeaders:     "",
		AllowCredentials: true,
		MaxAge:           0,
	}))
	app.Use(logger.New(logger.ConfigDefault))

	routes.Routes(app)

	err := app.Listen(port)
	if err != nil {
		return
	}

}
