package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var Db *sql.DB

func DatabaseConnector() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occured on env file please check")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	postgressetup := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable ", host, port, user, name, password)

	db, sqlErr := sql.Open("postgres", postgressetup)
	if sqlErr != nil {
		fmt.Printf("Error Trying To Connect TO Postgres DB %s\n", sqlErr)
		panic(sqlErr)
	}

	Db = db
	sqlErr = db.Ping()
	if sqlErr != nil {
		panic(sqlErr)
	}

	fmt.Println("Connected To Database Successfully")
}
