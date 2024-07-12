package workouts_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	ID                uint
	WORKOUT_NICKNAME  string `json:"workout_nickname"`
	WARMUP_ACTIVITY   string `json:"warmup_activity"`
	WARMUP_TIME       string `json:"warmupt_time"`
	CARDIO_ACTIVITY   string `json:"cardio_activity"`
	CARDIO_TIME       string `json:"cardio_time"`
	STRENGTH_ACTIVITY string `json:"strength_activity"`
	STRENGTH_TIME     string `json:"strength_time"`
	CORE_ACTIVITY     string `json:"core_activity"`
	CORE_TIME         string `json:"core_time"`
	FLEX_ACTIVITY     string `json:"flex_activity"`
	FLEX_TIME         string `json:"flex_time"`
	COOLDOWN_ACTIVITY string `json:"cooldown_activity"`
	COOLDOWN_TIME     string `json:"cooldown_time"`
	STREAK            string `json:"streak"`
	USERNAME          string `json:"username"`
}

func CreateWorkoutHandler(c *gin.Context) {
	//get values from request
	workout_nickname := c.Request.FormValue("workout_nickname")
	warmup_activity := c.Request.FormValue("warmup_activity")
	warmup_time := c.Request.FormValue("warmup_time")
	cardio_activity := c.Request.FormValue("cardio_activity")
	cardio_time := c.Request.FormValue("cardio_time")
	strength_activity := c.Request.FormValue("strength_activity")
	strength_time := c.Request.FormValue("strength_time")
	core_activity := c.Request.FormValue("core_activity")
	core_time := c.Request.FormValue("core_time")
	flex_activity := c.Request.FormValue("flex_activity")
	flex_time := c.Request.FormValue("flex_time")
	cooldown_activity := c.Request.FormValue("cooldown_activity")
	cooldown_time := c.Request.FormValue("cooldown_time")
	streak := "0"
	username := c.Request.FormValue("username")

	workout := Workout{
		WORKOUT_NICKNAME:  workout_nickname,
		WARMUP_ACTIVITY:   warmup_activity,
		WARMUP_TIME:       warmup_time,
		CARDIO_ACTIVITY:   cardio_activity,
		CARDIO_TIME:       cardio_time,
		STRENGTH_ACTIVITY: strength_activity,
		STRENGTH_TIME:     strength_time,
		CORE_ACTIVITY:     core_activity,
		CORE_TIME:         core_time,
		FLEX_ACTIVITY:     flex_activity,
		FLEX_TIME:         flex_time,
		COOLDOWN_ACTIVITY: cooldown_activity,
		COOLDOWN_TIME:     cooldown_time,
		STREAK:            streak,
		USERNAME:          username,
	}

	result := database.Db.Create(&workout)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create workout"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Workout Created Successfully"})
	}

}

func GetWorkoutsHandler(c *gin.Context) {
	username := c.Param("username")

	var workouts []Workout
	result := database.Db.Table("workouts").
		Select("workout_nickname, warmup_activity, warmup_time, cardio_activity, cardio_time, strength_activity, strength_time, core_activity, core_time, flex_activity, flex_time, cooldown_activity, cooldown_time, streak").
		Where("username = ? AND deleted_at IS NULL", username).
		Scan(&workouts)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem returning workouts"})
		return
	}

	workoutsJSON, err := json.Marshal(workouts)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error Trying To Return Object"})
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(workoutsJSON))
}

func UpdateStreakHandler(c *gin.Context) {
	username := c.Param("username")
	workout_nickname := c.Param("workout_nickname")

	var workout Workout
	result := database.Db.Select("streak").Where("username = $1 AND workout_nickname = $2", username, workout_nickname).First(&workout)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ran into a problem getting the streak"})
		return
	}

	streak := workout.STREAK
	fmt.Println(streak)

	//convert streak to string
	intStreak, intErr := strconv.Atoi(streak)
	if intErr != nil {
		fmt.Printf("Error trying to parse streak value: %v\n", intErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problems parsing the streak"})
		return
	}

	newStreak := intStreak + 1

	updatedStreak := strconv.Itoa(newStreak)
	fmt.Println(updatedStreak)

	//update the database
	updateResult := database.Db.
		Model(&Workout{}).
		Where("username = ? AND workout_nickname = ?", username, workout_nickname).
		Update("streak", updatedStreak)
	if updateResult.Error != nil {
		fmt.Printf("Error trying to update the streak: %v\n", result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem updating the streak"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully updated streak"})
	}

}

func DeleteWorkoutHandler(c *gin.Context) {
	username := c.Param("username")
	workout_nickname := c.Param("workout_nickname")

	result := database.Db.Where("username = ? AND workout_nickname = ?", username, workout_nickname).Delete(&Workout{})
	if result.Error != nil {
		fmt.Println("Failed to delete workout")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete workout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted workout"})
}
