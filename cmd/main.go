package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/omrfrkazt/fiber-gorm-simple-api/db"
	"github.com/omrfrkazt/fiber-gorm-simple-api/user"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)
	app := fiber.New()
	app.Get("/user/:id", handler.Get)
	app.Post("/user", handler.Create)
	app.Listen(":3001")
	fmt.Println("Server running on port 3000")
}
