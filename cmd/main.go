package main

import (
	"context"
	"log"

	"github.com/frostnzx/go-myapi-assignment/internal/adapters"
	"github.com/frostnzx/go-myapi-assignment/internal/core"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()


	rdb := redis.NewClient(&redis.Options{
		Addr: "myapi-redis:6379",
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Successfully connected to Redis!")
	}

	profileRepo := adapters.NewRedisProfileRepository(rdb)
	profileService := core.NewProfileService(profileRepo)
	profileHandler := adapters.NewHttpProfileHandler(profileService)

	// Routes
	app.Put("/profile", profileHandler.CreateProfile)
	app.Get("/profiles", profileHandler.GetProfiles)

	app.Listen(":8081")
}
