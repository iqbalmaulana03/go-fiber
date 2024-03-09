package service

import (
	"api/data/request"
	"api/data/response"
)

type NoteService interface {
	Create(note request.NoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
