package token_routes

import (
	token_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/token-service/token-handlers"
	token_middleware "github.com/Brian-Mashavakure/fitness-app-backend/pkg/token-service/token-middleware"
	"github.com/gin-gonic/gin"
)

func TokenRoutes(router *gin.Engine) {
	api := router.Group("api/fitness")

	api.POST("/refreshtoken", token_middleware.TokenMiddleware(), token_handlers.RefreshToken)
}
