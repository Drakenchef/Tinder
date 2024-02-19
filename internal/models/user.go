package models

import "github.com/google/uuid"

// satori uid
type User struct {
	UID         uuid.UUID `json:"uid" db:"uid"`
	Login       string    `json:"login" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Salt        string    `json:"salt"`
}

type SignInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
