package goals_routes

import (
	goals_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/goals-service/goals-handlers"
	"github.com/gin-gonic/gin"
)

func GoalsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness")

	api.POST("/addgoal", goals_handlers.CreateGoalHandler)

	api.GET("/getgoals/:username", goals_handlers.GetGoalsHandler)

	api.PUT("/deletegoal/:title", goals_handlers.DeleteGoalHandler)
}
