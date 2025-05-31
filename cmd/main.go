package main

import (
	"github.com/frostnzx/go-myapi-assignment/internal/adapters"
	"github.com/frostnzx/go-myapi-assignment/internal/core"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	profileRepo := adapters.NewRedisProfileRepository(rdb)
	profileService := core.NewProfileService(profileRepo)
	profileHandler := adapters.NewHttpProfileHandler(profileService)

	// routes
	app.Put("/profile", profileHandler.CreateProfile) 
	app.Get("/profiles", profileHandler.GetProfiles)

	app.Listen(":5000")
}