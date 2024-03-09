package main

import (
	"api/config"
	"api/controller"
	"api/model"
	"api/repository"
	"api/route"
	"api/service"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Run service ...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}

	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	noteRepository := repository.NewNoteRepositoryImpl(db)

	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	noteController := controller.NewNoteController(noteService)

	route := route.NewRoute(noteController)

	app := fiber.New()

	app.Mount("/api", route)

	log.Fatal(app.Listen(":9000"))
}
