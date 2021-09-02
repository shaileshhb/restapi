package user

import "github.com/shaileshhb/restapi/model/general"

type User struct {
	general.Base
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Username string `gorm:"type:varchar(30)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
