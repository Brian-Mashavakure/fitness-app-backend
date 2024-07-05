package workouts_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to create workout"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout Created Successfully"})
}

func GetWorkoutsHandler(c *gin.Context) {
	username := c.Param("username")

	//query the database for all workouts related to a single user
	rows, err := database.Db.Table("workouts").
		Select("workout_nickname, warmup_activity, warmup_time, cardio_activity, cardio_time, strength_activity, strength_time, core_activity, core_time, flex_activity, flex_time, cooldown_activity, cooldown_time, streak").
		Where("username = $1", username).
		Find(&Workout{}).
		Rows()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem trying to query the database"})
	}

	defer rows.Close()

	var workouts []Workout

	for rows.Next() {
		var workout_nickname string
		var warmup_activity string
		var warmup_time string
		var cardio_activity string
		var cardio_time string
		var strength_activity string
		var strength_time string
		var core_activity string
		var core_time string
		var flex_activity string
		var flex_time string
		var cooldown_activity string
		var cooldown_time string
		var streak string

		if err := rows.Scan(&workout_nickname, &warmup_activity, &warmup_time, &cardio_activity, &cardio_time, &strength_activity, &strength_time, &core_activity, &core_time, &flex_activity, &flex_time, &cooldown_activity, &cooldown_time, &streak); err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem returning workouts"})
			return
		}

		workouts = append(workouts, Workout{WORKOUT_NICKNAME: workout_nickname, WARMUP_ACTIVITY: warmup_activity, WARMUP_TIME: warmup_time, CARDIO_ACTIVITY: cardio_activity, CARDIO_TIME: cardio_time, STRENGTH_ACTIVITY: strength_activity, STRENGTH_TIME: strength_time, CORE_ACTIVITY: core_activity, CORE_TIME: core_time, FLEX_ACTIVITY: flex_activity, FLEX_TIME: flex_time, COOLDOWN_ACTIVITY: cooldown_activity, COOLDOWN_TIME: cooldown_time, STREAK: streak})

	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
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

func GetSingleWorkoutHandler(c *gin.Context) {
	username := c.Param("username")
	workout_nickname := c.Param("workout_nickname")

	//query the database for a single workout using the workout nickname and username
	rows, err := database.Db.Table("workouts").
		Select("workout_nickname, warmup_activity, warmup_time, cardio_activity, cardio_time, strength_activity, strength_time, core_activity, core_time, flex_activity, flex_time, cooldown_activity, cooldown_time, streak").
		Where("username = $1 AND workout_nickname = $2", username, workout_nickname).
		Find(&Workout{}).
		Rows()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error Querying The Database"})
	}

	defer rows.Close()

	var workout []Workout

	for rows.Next() {
		var workout_nickname string
		var warmup_activity string
		var warmup_time string
		var cardio_activity string
		var cardio_time string
		var strength_activity string
		var strength_time string
		var core_activity string
		var core_time string
		var flex_activity string
		var flex_time string
		var cooldown_activity string
		var cooldown_time string
		var streak string

		if err := rows.Scan(&workout_nickname, &warmup_activity, &warmup_time, &cardio_activity, &cardio_time, &strength_activity, &strength_time, &core_activity, &core_time, &flex_activity, &flex_time, &cooldown_activity, &cooldown_time, &streak); err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Problem trying to parse workout"})
			return
		}

		workout = append(workout, Workout{WORKOUT_NICKNAME: workout_nickname, WARMUP_ACTIVITY: warmup_activity, WARMUP_TIME: warmup_time, CARDIO_ACTIVITY: cardio_activity, CARDIO_TIME: cardio_time, STRENGTH_ACTIVITY: strength_activity, STRENGTH_TIME: strength_time, CORE_ACTIVITY: core_activity, CORE_TIME: core_time, FLEX_ACTIVITY: flex_activity, FLEX_TIME: flex_time, COOLDOWN_ACTIVITY: cooldown_activity, COOLDOWN_TIME: cooldown_time, STREAK: streak})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error returning object"})
		return
	}

	workoutJson, err := json.Marshal(workout)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error trying to parse object"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(workoutJson))
}
