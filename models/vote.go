package models

import (
	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	CommentId uint `json: "commentId" gorm:"not null"`
	UserId    uint `json: "userId" gorm:"not null"`
	Value     bool `json: "value" gorm:"not null"` //si ha votado
}
