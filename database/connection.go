package database

import (
	"fmt"
	config "go-echo-boilerplate/setup"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var onceDb sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := config.DatabaseNew().(*config.DatabaseConfig)
		db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			databaseConfig.Psql.DbHost,
			databaseConfig.Psql.DbPort,
			databaseConfig.Psql.DbUsername,
			databaseConfig.Psql.DbDatabase,
			databaseConfig.Psql.DbPassword,
			databaseConfig.Psql.DbSslmode,
		))
		if err != nil {
			fmt.Println("Could not connect to database :%v", err)
			log.Fatalf("Could not connect to database :%v", err)
			panic(0)
		}
		fmt.Println("Ok")
		db.DB().SetMaxIdleConns(10)

		db.DB().SetMaxOpenConns(100)

		db.DB().SetConnMaxLifetime(time.Hour)

		fmt.Println("Database connection initialized successfully")	
		instance = db
	})
	return instance
}
