package goals_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	ID               uint
	TITLE            string `json:"title"`
	GOAL_DESC        string `json:"goal_description"`
	DATE_SET         string `json:"date_set"`
	FINISH_DATE      string `json:"finish_date"`
	USERNAME         string `json:"username"`
	WORKOUT_NICKNAME string `json:"workout_nickname"`
}

func CreateGoalHandler(c *gin.Context) {
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("goal_description")
	date_set := c.Request.FormValue("date_set")
	finish_date := c.Request.FormValue("finish_date")
	username := c.Request.FormValue("username")
	workout_nickname := c.Request.FormValue("workout_nickname")

	goal := Goal{
		TITLE:            title,
		GOAL_DESC:        description,
		DATE_SET:         date_set,
		FINISH_DATE:      finish_date,
		USERNAME:         username,
		WORKOUT_NICKNAME: workout_nickname,
	}

	result := database.Db.Create(&goal)
	if result.Error != nil {
		fmt.Println("Failed to create goal")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully added goal"})
}

func GetGoalsHandler(c *gin.Context) {
	username := c.Param("username")

	var goals []Goal
	result := database.Db.Table("goals").
		Select("id, title, goal_desc, date_set, finish_date, username, workout_nickname").
		Where("username = ? AND deleted_at IS NULL", username).
		Scan(&goals)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem trying to query the database"})

	}

	goalsJson, err := json.Marshal(goals)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error trying to parse object"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(goalsJson))
}

func DeleteGoalHandler(c *gin.Context) {
	title := c.Param("title")

	result := database.Db.
		Where("title = ?", title).
		Delete(&Goal{})
	if result.Error != nil {
		fmt.Println("Failed to delete goal")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted goal"})
}
