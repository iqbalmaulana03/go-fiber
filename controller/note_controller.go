package controller

import (
	"api/data/request"
	"api/data/response"
	"api/helper"
	"api/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{noteService: service}
}

func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	createNoteRequest := request.NoteRequest{}
	err := ctx.BodyParser(&createNoteRequest)
	helper.ErrorPanic(err)

	controller.noteService.Create(createNoteRequest)

	webResponse := response.Response{
		Code:    201,
		Status:  "Created",
		Message: "successfully create note data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	updateNoteRequest := request.UpdateNoteRequest{}
	err := ctx.BodyParser(&updateNoteRequest)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	updateNoteRequest.Id = id

	controller.noteService.Update(updateNoteRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "successfully update note data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	controller.noteService.Delete(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "successfully delete note data!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	noteResponse := controller.noteService.FindById(id)

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "successfully get by id note data!",
		Data:    noteResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	noteResponse := controller.noteService.FindAll()

	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "successfully get all note data!",
		Data:    noteResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
