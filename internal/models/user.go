package models

type User struct {
	UID      int    `json:"uid" db:"uid"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Profile struct {
	UID         int    `json:"uid" db:"uid"`
	Login       string `json:"login" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
}
