package main

import (
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-handlers"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DatabaseConnector()

	//Run database migrations for tables
	database.Db.AutoMigrate(&workouts_handlers.Workout{})

	router := gin.Default()

	workouts_routes.WorkoutsRoutes(router)

	router.Run("localhost:8080")

}
