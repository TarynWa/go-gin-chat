package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID uint
	Username  string `json:"username"`
	Password  string `json:"password"`
	AvatarId  string `json:"avatar_id"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}


func FindUserByField(field, value string) User {
	var u User

	if field == "id" || field == "username" {
		ChatDB.Where(field+" = ?", value).First(&u)
	}

	return u
}