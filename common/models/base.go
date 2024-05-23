package models

import (
	// "gorm.io/gorm"
	// uuid "github.com/satori/go.uuid"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	// ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	ID        uint 	     `gorm:"primaryKey"`	
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
// func (base *Base) BeforeCreate(scope *gorm.Scope) error {
// 	uid, err := uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uid)
// }
