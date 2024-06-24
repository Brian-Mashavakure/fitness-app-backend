package main

import (
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DatabaseConnector()

	router := gin.Default()

	workouts_routes.WorkoutsRoutes(router)

	router.Run("localhost:8080")

}
