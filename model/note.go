package model

type Note struct {
	Id      int    `grom:"type:int;primary_key"`
	Content string `gorm:"not null" json:"content,omitempty"`
}
