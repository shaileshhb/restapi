package model

type User struct {
	Base
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Username string `gorm:"type:varchar(30)" json:"username"`
	Password string `gorm:"type:varchar(30)" json:"password"`
}

var JwtKey = []byte("some_secret_key")
