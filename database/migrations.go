package database

import (
	// BlogModels "go-echo-boilerplate/blogs/models"
	UserModels "go-echo-boilerplate/src/models/users/model"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "2020080206",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&UserModels.User{}).Error; err != nil {
					return err
				}
				// if err := tx.AutoMigrate(&BlogModels.Blog{}).Error; err != nil {
				// 	return err
				// }
				// if err := tx.AutoMigrate(&InfoModels.AbonentInfo{}).Error; err != nil {
				// 	return err
				// }
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.DropTable("blogs").Error; err != nil {
					return nil
				}
				if err := tx.DropTable("users").Error; err != nil {
					return nil
				}
				if err := tx.DropTable("userinfo").Error; err != nil {
					return nil
				}
				return nil
			},
		},
	})
}
