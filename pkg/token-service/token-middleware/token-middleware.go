package token_middleware

import (
	"fmt"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	token_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/token-service/token-handlers"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		username := c.Request.FormValue("username")

		formattedToken := string(tokenString)

		fmt.Printf("Header token is as follows: %s\n", tokenString)

		fmt.Printf("Formatted token is as follows: %s\n", formattedToken)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request not authorized"})
			c.Abort()
			return
		}

		//retrieve user token info
		token := token_handlers.Token{USERNAME: username}
		result := database.Db.Table("tokens").Where("username = ?", username).Scan(&token)
		if result.Error != nil {
			log.Println("Error finding user or token in db")
			c.Abort()
			return
		}

		fmt.Printf("Database token is as follows: %s\n", token.TOKEN)

		//check if token string is the same
		if tokenString != token.TOKEN {
			log.Println("Invalid token string")
			c.JSON(http.StatusOK, gin.H{"error": "Invalid token string provided"})
			return
		}

		//check if token is still valid
		todayDate := time.Now().Format("02-01-2006")
		status := utils.CompareDates(todayDate, token.EXPIRY_DATE)
		if status == false {
			log.Println("Token has expired")
			c.JSON(http.StatusOK, gin.H{"error": "Expired token string provided"})
			return
		}

		//continue
		c.Next()
	}

}