package goals_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	workouts_handlers "github.com/Brian-Mashavakure/fitness-app-backend/pkg/workouts-service/workouts-handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	ID               uint
	TITLE            string                    `json:"title"`
	GOAL_DESC        string                    `json:"goal_description"`
	DATE_SET         string                    `json:"date_set"`
	FINISH_DATE      string                    `json:"finish_date"`
	USERNAME         string                    `json:"username"`
	WORKOUT_NICKNAME string                    `json:"workout_nickname" gorm:"unique"`
	Workout          workouts_handlers.Workout `gorm:"foreignKey:workout_nickname;references:workout_nickname"`
}

func CreateGoalHandler(c *gin.Context) {
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
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

	//query the database for all workouts related to a single user
	rows, err := database.Db.Table("goals").
		Select("id, title, goal_desc, date_set, finish_date, username, workout_nickname").
		Where("username = $1", username).
		Find(&Goal{}).
		Rows()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem trying to query the database"})
	}

	defer rows.Close()

	var goals []Goal

	for rows.Next() {
		var id uint
		var title string
		var description string
		var date_set string
		var finish_date string
		var workout_nickname string

		if err := rows.Scan(&id, &title, &description, &date_set, &finish_date, &username, &workout_nickname); err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem trying to parse goal"})
			return
		}

		goals = append(goals, Goal{ID: id, TITLE: title, GOAL_DESC: description, DATE_SET: date_set, FINISH_DATE: finish_date, USERNAME: username, WORKOUT_NICKNAME: workout_nickname})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error returning object"})
		return
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
	id := c.Param("id")

	result := database.Db.Where("id = ?", id).Delete(&Goal{})
	if result.Error != nil {
		fmt.Println("Failed to delete goal")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete goal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted goal"})
}
