package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UID         uuid.UUID `json:"uid" db:"uid"`
	Login       string    `json:"login" binding:"required" validate:"required,min=6,alphanum"`
	Password    string    `json:"-" binding:"required" validate:"required,min=6"`
	Images      []Image   `json:"images"`
	Description string    `json:"description" binding:"required"`
	Salt        string    `json:"-"`
}

type Image struct {
	ID     int       `json:"id" db:"id"`
	URL    string    `json:"url" binding:"required"`
	UserID uuid.UUID `json:"userid" db:"userid"`
}

type Profile struct {
	Login       string `json:"login" binding:"required" validate:"required,min=6,alphanum"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type SignInInput struct {
	Login    string `json:"login" binding:"required" validate:"required,min=6,alphanum"`
	Password string `json:"password" binding:"required" validate:"required,min=6,alphanumunicode"`
}

type ChangePassword struct {
	Password string `json:"password" binding:"required" validate:"required,min=6,alphanumunicode"`
}

type Likes struct {
	UIDFirstLike  uuid.UUID `json:"UIDFirstLike" db:"UIDFirstLike"`
	UIDSecondLike uuid.UUID `json:"UIDSecondLike" db:"UIDSecondLike"`
	Date          time.Time `json:"Date"`
	Mutual        bool      `json:"Mutual"`
}

type LikesUID struct {
	UIDFirstLike  string `json:"UIDFirstLike" db:"UIDFirstLike"`
	UIDSecondLike string `json:"UIDSecondLike" db:"UIDSecondLike"`
}
