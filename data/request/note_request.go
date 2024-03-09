package request

type NoteRequest struct {
	Content string `validator:"required,min=2,max=100" json:"content"`
}

type UpdateNoteRequest struct {
	Id      int    `validator:"required"`
	Content string `validator:"required,min=2,max=100" json:"content"`
}
