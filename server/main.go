package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/krazy-code/devlink/configs"
	"github.com/krazy-code/devlink/routes"
	"github.com/krazy-code/devlink/utils"
)

func init() {
	if os.Getenv("STAGE_STATUS") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}
}
func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)
	routes.InitRouter(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
