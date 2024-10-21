package goals_routes

import (
	goals_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/goals-service/goals-handlers"
	token_moddleware "github.com/Brian-Mashavakure/fitness-app-backend/pkg/token-service/token-middleware"
	"github.com/gin-gonic/gin"
)

func GoalsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness/goals")

	api.POST("/addgoal", token_moddleware.TokenMiddleware(), goals_handlers.CreateGoalHandler)

	api.GET("/getgoals", token_moddleware.TokenMiddleware(), goals_handlers.GetGoalsHandler)

	api.PUT("/deletegoal", token_moddleware.TokenMiddleware(), goals_handlers.DeleteGoalHandler)
}
