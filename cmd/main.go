package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	





	app := fiber.New()
	


	app.Listen(":3000")
}