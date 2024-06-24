package workouts_handlers

import (
	"fmt"
	"github.com/Brian-Mashavakure/fitness-app-backend/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Workout struct {
	WORKOUT_NICKNAME  string `json:"workout_nickname"`
	WARMUP_ACTIVITY   string `json:"warmup_activity"`
	WARMUP_TIME       string `json:"warmupt_time"`
	CARDIO_ACTICITY   string `json:"cardio_activity"`
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

	workout := Workout{
		WORKOUT_NICKNAME:  workout_nickname,
		WARMUP_ACTIVITY:   warmup_activity,
		WARMUP_TIME:       warmup_time,
		CARDIO_ACTICITY:   cardio_activity,
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
	}

	_, dbErr := database.Db.Exec("INSERT INTO workouts(workout_nickname, warmup_activity, warmup_time, cardio_activity, cardio_time, strength_activity, strength_time, core_activity, core_time, flex_activity, flex_time, cooldown_activity, cooldown_time, streak) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", workout.WORKOUT_NICKNAME, workout.WARMUP_ACTIVITY, workout.WARMUP_TIME, workout.CARDIO_ACTICITY, workout.CARDIO_TIME, workout.STRENGTH_ACTIVITY, workout.STRENGTH_TIME, workout.CORE_ACTIVITY, workout.CORE_TIME, workout.FLEX_ACTIVITY, workout.FLEX_TIME, workout.COOLDOWN_ACTIVITY, workout.COOLDOWN_TIME, workout.STREAK)
	if dbErr != nil {
		fmt.Println(dbErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed To Create Workout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout Created Successfully"})
}
