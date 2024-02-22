package models

import "github.com/google/uuid"

type User struct {
	UID         uuid.UUID `json:"uid" db:"uid"`
	Login       string    `json:"login" binding:"required" validate:"required,min=6,alphanum"`
	Password    string    `json:"password" binding:"required" validate:"required,min=6"`
	Image       string    `json:"image" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Salt        string    `json:"-"`
}

type SignInInput struct {
	Login    string `json:"login" binding:"required" validate:"required,min=6,alphanum"`
	Password string `json:"password" binding:"required" validate:"required,min=6,alphanumunicode"`
}
