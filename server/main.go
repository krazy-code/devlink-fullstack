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

// var ctx = context.Background()

func init() {
	if os.Getenv("STAGE_STATUS") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}
}
func main() {
	config := configs.FiberConfig()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })

	app := fiber.New(config)

	// err := rdb.Set(ctx, "account:user1", 1000, 0).Err()
	// if err != nil {
	// 	log.Fatalf("Could not set initial account data: %v", err)
	// }
	// err = rdb.Set(ctx, "account:user2", 500, 0).Err()
	// if err != nil {
	// 	log.Fatalf("Could not set initial account data: %v", err)
	// }

	routes.InitRouter(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
