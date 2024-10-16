package runs_routes

import (
	runs_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/runs-service/runs-handlers"
	token_moddleware "github.com/Brian-Mashavakure/fitness-app-backend/pkg/token-service/token-middleware"
	"github.com/gin-gonic/gin"
)

func RunsRoutes(router *gin.Engine) {
	api := router.Group("api/fitness/runs")

	api.POST("/createrun", token_moddleware.TokenMiddleware(), runs_handlers.CreateRunHandler)

	api.GET("/getruns/:username", token_moddleware.TokenMiddleware(), runs_handlers.GetRunsHandler)

	api.GET("/runsleaderboard", token_moddleware.TokenMiddleware(), runs_handlers.LeaderboardHandler)

	api.PUT("/deleterun/:id", token_moddleware.TokenMiddleware(), runs_handlers.DeleteRunHandler)

}
