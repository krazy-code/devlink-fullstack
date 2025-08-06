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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
