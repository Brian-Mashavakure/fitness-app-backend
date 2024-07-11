package runs_routes

import (
	runs_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/runs-service/runs-handlers"
	"github.com/gin-gonic/gin"
)

func RunsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness")

	api.POST("/createrun", runs_handlers.CreateRunHandler)
	api.GET("/getruns/:username", runs_handlers.GetRunsHandler)
	api.GET("/runsleaderboard", runs_handlers.LeaderboardHandler)

}
