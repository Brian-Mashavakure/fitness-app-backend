package runs_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Run struct {
	gorm.Model
	ID               uint
	SET_DISTANCE     string `json:"set_distance"`
	START_TIME       string `json:"start_time"`
	END_TIME         string `json:"end_time"`
	COVERED_DISTANCE string `json:"covered_distance"`
	TIME_TAKEN       string `json:"time_taken"`
	USERNAME         string `json:"username"`
}

func CreateRunHandler(c *gin.Context) {
	//get run values from request
	set_distance := c.Request.FormValue("set_distance")
	start_time := c.Request.FormValue("start_time")
	end_time := c.Request.FormValue("end_time")
	covered_distance := c.Request.FormValue("covered_distance")
	time_taken := c.Request.FormValue("time_taken")
	username := c.Request.FormValue("username")

	//create run object
	run := Run{
		SET_DISTANCE:     set_distance,
		START_TIME:       start_time,
		END_TIME:         end_time,
		COVERED_DISTANCE: covered_distance,
		TIME_TAKEN:       time_taken,
		USERNAME:         username,
	}
	//run create method on database
	result := database.Db.Create(&run)
	if result.Error != nil {
		fmt.Println("Error running create method on database")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem trying to create run"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully created run"})
}

// TODO:Add deleted at clause to query
func GetRunsHandler(c *gin.Context) {
	//get username
	username := c.Param("username")

	var runs []Run

	result := database.Db.
		Table("runs").
		Select("id, set_distance, start_time, end_time, covered_distance, time_taken, username").
		Where("username = ? AND deleted_at IS NULL", username).
		Scan(&runs)
	if result.Error != nil {
		fmt.Println("Problems trying to query database for runs")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problems trying to query the database"})
		return
	}

	runsJson, jsonErr := json.Marshal(runs)
	if jsonErr != nil {
		fmt.Println("Problems trying to parse runs")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problems trying to return the runs"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(runsJson))
}

// return running leaderboard for community feature
func LeaderboardHandler(c *gin.Context) {
	var leaderboard []Run

	result := database.Db.Table("runs").
		Select("id,set_distance, start_time, end_time, covered_distance, time_taken, username").
		Order("covered_distance desc").
		Scan(&leaderboard)
	if result.Error != nil {
		fmt.Println("Problems trying to query database for leaderboard")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problems trying to query the database"})
		return
	}

	leaderboardJson, jsonErr := json.Marshal(leaderboard)
	if jsonErr != nil {
		fmt.Println("Problems trying to parse leaderboard")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problems trying to return the leaderboard"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(leaderboardJson))

}

// TODO: Add delete run handler
func DeleteRunHandler(c *gin.Context) {
	id := c.Param("id")

	result := database.Db.
		Where("id = ?", id).
		Delete(&Run{})
	if result.Error != nil {
		fmt.Println("Failed to delete run")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete run"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted run"})

}
