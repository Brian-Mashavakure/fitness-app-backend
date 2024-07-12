package workouts_routes

import (
	workouts_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-handlers"
	"github.com/gin-gonic/gin"
)

func WorkoutsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness")

	api.POST("/createworkout", workouts_handlers.CreateWorkoutHandler)

	api.GET("/getworkouts/:username", workouts_handlers.GetWorkoutsHandler)

	api.PUT("/updatestreak/:username/:workout_nickname", workouts_handlers.UpdateStreakHandler)

	api.PUT("/deleteworkout/:username/:workout_nickname", workouts_handlers.DeleteWorkoutHandler)
}
