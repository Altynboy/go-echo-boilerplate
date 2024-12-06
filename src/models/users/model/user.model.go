package models

import (
	"go-echo-boilerplate/common"
	CommonModels "go-echo-boilerplate/common/models"
	"go-echo-boilerplate/common/utils"

	"gorm.io/gorm"
)

type User struct {
	CommonModels.Base
	Phone      string `gorm:"type:character varying(12);unique_index"`
	Role       common.UserRole
	Password   string
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
}

func (User) TableName() string {
	return "users"
}

// func (user User) String() string {
// 	return user.Name
// }

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	hashed, err := utils.GetPasswordUtil().HashPassword(user.Password)
	user.Password = hashed
	return
}
