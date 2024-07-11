package main

import (
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	goals_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/goals-service/goals-handlers"
	runs_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/runs-service/runs-handlers"
	runs_routes "github.com/Brian-Mashavakure/fitness-app-backend/pkg/runs-service/runs-routes"
	workouts_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-handlers"
	workouts_routes "github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DatabaseConnector()

	//Run database migrations for tables
	database.Db.AutoMigrate(&goals_handlers.Goal{}, &workouts_handlers.Workout{}, &runs_handlers.Run{})

	router := gin.Default()

	workouts_routes.WorkoutsRoutes(router)
	runs_routes.RunsRoutes(router)

	router.Run("localhost:8080")

}
