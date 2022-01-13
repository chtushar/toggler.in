package database

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toggler.in/api/config"
	"toggler.in/api/model"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	dbHost := config.Config("DB_HOST")
	dbUser := config.Config("DB_USER")
	dbPassword := config.Config("DB_PASSWORD")
	dbName := config.Config("DB_NAME")

	dsn := url.URL{
		User:   url.UserPassword(dbUser, dbPassword),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", dbHost, port),
		Path:   dbName,
	}

	// Connect to the Database
	DB, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.User{})

	fmt.Println("Connection Opened to Database")
}
