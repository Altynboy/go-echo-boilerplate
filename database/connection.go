package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDb sync.Once

var instance *gorm.DB

func Init() {
	onceDb.Do(func() {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
				viper.GetString("Db.Host"),
				viper.GetString("Db.Port"),
				viper.GetString("Db.Name"),
				viper.GetString("Db.User"),
				viper.GetString("Db.Password"),
				viper.GetString("Db.SslMode"),
			)}), &gorm.Config{})
		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}

		dbCon, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB instance: %v", err)
		}

		dbCon.SetMaxIdleConns(10)
		dbCon.SetMaxOpenConns(100)
		dbCon.SetConnMaxLifetime(time.Hour)

		fmt.Println("Database connection initialized successfully")
		instance = db
	})
}

func Instance() *gorm.DB {
	return instance
}
