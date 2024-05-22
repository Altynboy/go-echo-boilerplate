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
