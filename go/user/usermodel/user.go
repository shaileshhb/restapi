package model

import "github.com/shaileshhb/restapi/model"

type User struct {
	model.Base
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
