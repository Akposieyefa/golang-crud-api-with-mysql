package models

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" validate:"required"`
	Body  string `json:"body"`
}

func (p *Post) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
