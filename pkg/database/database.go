package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DatabaseConnector() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occurred on env file please check")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	//Database connection string
	postgressetup := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require", host, port, user, name, password)

	//Using gorm to connect to postgres database
	db, sqlErr := gorm.Open(postgres.Open(postgressetup), &gorm.Config{})
	//sqlErr = db.Ping()
	if sqlErr != nil {
		panic(sqlErr)
	}

	Db = db
}
