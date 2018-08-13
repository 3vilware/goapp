package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId   uint      `json: "userId"`
	ParentId uint      `json: "parentId"`
	Votes    int32     `json: "votes"`
	Content  string    `json: "content"`
	HasVote  int8      `json: "hasVote" gorm:"-"` //ver si el usuario tiene votado un comentario
	User     []User    `json: "user,omitempty"`
	Children []Comment `json: "children,omitempty"`
}
