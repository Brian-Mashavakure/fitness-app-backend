package token_handlers

import (
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Token struct {
	gorm.Model
	ID          uint
	USERNAME    string `json:"username" gorm:"unique"`
	TOKEN       string `json:"token" gorm:"unique"`
	START_DATE  string
	EXPIRY_DATE string
}

func RefreshToken(c *gin.Context) {
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	tokenString := c.Request.Header.Get("Authorization")

	token := Token{USERNAME: username}

	result := database.Db.Table("tokens").Where("username = ?", username).Scan(&token)
	if result.Error != nil {
		log.Println("Error retrieving user from db")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Failed to find username"})
		return
	}

	if tokenString == token.TOKEN {
		//check if token is still valid
		todayDate := time.Now().Format("02/01/2006")
		status := utils.CompareDates(todayDate, token.EXPIRY_DATE)
		if status == true {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Token is still valid"})
		} else if status == false {
			newTokenString, tokenStartTime, tokenExpiryTime := utils.GenerateToken(username, email)
			newToken := Token{
				USERNAME:    username,
				TOKEN:       newTokenString,
				START_DATE:  tokenStartTime,
				EXPIRY_DATE: tokenExpiryTime,
			}

			newTokenResult := database.Db.Create(&newToken)
			if newTokenResult.Error != nil {
				log.Println("Error generating new token")
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Failed to generate new token"})
				return
			}

			c.String(http.StatusOK, newTokenString)
		}

	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Token provided not in our database"})
	}

}
