package entity

// User will consist of registration details of user.
type User struct {
	Base
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Username string `gorm:"type:varchar(30)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
