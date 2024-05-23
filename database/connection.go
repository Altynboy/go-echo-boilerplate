package database

import (
	"fmt"
	config "go-echo-boilerplate/setup"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var onceDb sync.Once

var instance *gorm.DB

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		databaseConfig := config.DatabaseNew().(*config.DatabaseConfig)
		db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			databaseConfig.Psql.DbHost,
			databaseConfig.Psql.DbPort,
			databaseConfig.Psql.DbUsername,
			databaseConfig.Psql.DbDatabase,
			databaseConfig.Psql.DbPassword,
			databaseConfig.Psql.DbSslmode),
		}), &gorm.Config{})
		// 	, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		// 	databaseConfig.Psql.DbHost,
		// 	databaseConfig.Psql.DbPort,
		// 	databaseConfig.Psql.DbUsername,
		// 	databaseConfig.Psql.DbDatabase,
		// 	databaseConfig.Psql.DbPassword,
		// 	databaseConfig.Psql.DbSslmode,
		// ))
		if err != nil {
			fmt.Println("Could not connect to database :%v", err)
			log.Fatalf("Could not connect to database :%v", err)
			panic(0)
		}
		dbCon, err := db.DB()

		dbCon.SetMaxIdleConns(10)

		dbCon.SetMaxOpenConns(100)

		dbCon.SetConnMaxLifetime(time.Hour)

		fmt.Println("Database connection initialized successfully")	
		instance = db
	})
	return instance
}
