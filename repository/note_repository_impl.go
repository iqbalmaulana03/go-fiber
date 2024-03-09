package repository

import (
	"api/data/request"
	"api/helper"
	"api/model"
	"errors"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	Db *gorm.DB
}

func NewNoteRepositoryImpl(Db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{Db: Db}
}

func (r *NoteRepositoryImpl) Save(note model.Note) {
	result := r.Db.Create(&note)
	helper.ErrorPanic(result.Error)
}

func (r *NoteRepositoryImpl) Update(note model.Note) {
	var updatedNote = request.UpdateNoteRequest{
		Id:      note.Id,
		Content: note.Content,
	}
	result := r.Db.Model(&note).Updates(updatedNote)
	helper.ErrorPanic(result.Error)
}

func (r *NoteRepositoryImpl) Delete(noteId int) {
	var note model.Note
	result := r.Db.Where("id = ?", noteId).Delete(&note)
	helper.ErrorPanic(result.Error)
}

func (r *NoteRepositoryImpl) FindById(noteId int) (model.Note, error) {
	var note model.Note
	result := r.Db.Find(&note, noteId)
	if result != nil {
		return note, nil
	} else {
		return note, errors.New("note is not found")
	}
}

func (r *NoteRepositoryImpl) FindAll() []model.Note {
	var note []model.Note
	result := r.Db.Find(&note)
	helper.ErrorPanic(result.Error)
	return note
}
