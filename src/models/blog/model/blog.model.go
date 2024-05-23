package models

import (
	CommonModels "go-echo-boilerplate/common/models"
)


type Blog struct {
	CommonModels.Base
	UserId  uint   `gorm:"type:integer"`
	Title   string `gorm:"type:character varying(255)"`
	Content string `gorm:"type:text"`
}

// GORM pluralizes struct name to snake_cases as table name,
// for struct Blog, its table name is blogs by convention

// If you want to use custom table name, you can define TableName method
// func (User) TableName() string {
// 	return "other_table_name"
// }