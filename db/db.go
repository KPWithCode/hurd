package db

import (
	"fmt"
	"log"
	"os"

	"github.com/KPWithCode/hurd/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	//load env
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	// ==============without gorm setup===================
	// db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	// if err != nil {
	// 	return err
	// }

	// Database Settings
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT") // Default port
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/New_York", host, port, user, password, dbname)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to db \n", err)
		os.Exit(2)
		//======without gorm setup===========
		// return err
	}
	db.Debug().AutoMigrate(&models.UserModel{})
	// makes db a global variable as DB was declared outside of function. ex: Now usable in User controller
	DB = db

	log.Printf("Successfully connected to the Waviest DB: JUU %v", dbname)
	// ==============without gorm setup===============
	// if err = db.Ping(); err != nil {
	// 	return err
	// }
	// return nil
}
