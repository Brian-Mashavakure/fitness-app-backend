package workouts_routes

import (
	token_moddleware "github.com/Brian-Mashavakure/fitness-app-backend/pkg/token-service/token-middleware"
	workouts_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-handlers"
	"github.com/gin-gonic/gin"
)

func WorkoutsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness/workouts")

	api.POST("/createworkout", token_moddleware.TokenMiddleware(), workouts_handlers.CreateWorkoutHandler)

	api.GET("/getworkouts/:username", token_moddleware.TokenMiddleware(), workouts_handlers.GetWorkoutsHandler)

	api.PUT("/updatestreak/:username/:workout_nickname", token_moddleware.TokenMiddleware(), workouts_handlers.UpdateStreakHandler)

	api.PUT("/deleteworkout/:username/:workout_nickname", token_moddleware.TokenMiddleware(), workouts_handlers.DeleteWorkoutHandler)
}
