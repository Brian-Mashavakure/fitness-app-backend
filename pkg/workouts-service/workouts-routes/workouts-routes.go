package workouts_routes

import (
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-handlers"
	"github.com/gin-gonic/gin"
)

func WorkoutsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness")

	api.POST("/createworkout", workouts_handlers.CreateWorkoutHandler)
}
