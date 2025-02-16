package handler

import (
	"net/http"
	"time"

	"github.com/Dhanraj-Patil/event-trigger/internal/models"
	"github.com/Dhanraj-Patil/event-trigger/internal/services"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type Trigger struct {
	UserId   string    `json:"userId"`
	Schedule time.Time `json:"schedule"`
	Message  string    `json:"message"`
	PhoneNo  string    `json:"phoneNo"`
}

type EditTrigger struct {
	Schedule time.Time `json:"schedule"`
	Message  string    `json:"message"`
	PhoneNo  string    `json:"phoneNo"`
}

func EventRouter(g *gin.Engine) {
	router := g.Group("/api")

	router.GET("/triggers", getTriggers)
	router.POST("/trigger", createTrigger)
	router.PUT("/trigger", editTrigger)
	router.DELETE("/trigger", deleteTrigger)

	router.POST("/trigger-test", testTrigger)

}

// @Summary Get all triggers for a user
// @Description Creates a trigger that calls an API at a scheduled time
// @Tags Triggers
// @Accept plain
// @Produce plain
// @Param userId query string true "Enter userId"
// @Success 200 {object} []Trigger
// @Failure 400 {object} map[string]string
// @Router /api/triggers [get]
func getTriggers(c *gin.Context) {
	userId := c.Query("userId")
	userTriggers, err := services.GetAllByUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userTriggers)
}

// @Summary Create a new trigger
// @Description Creates a trigger that calls an API at a scheduled time
// @Tags Triggers
// @Accept json
// @Produce json
// @Param trigger body Trigger true "Trigger Data"
// @Success 200 {object} Trigger
// @Failure 400 {object} map[string]string
// @Router /api/trigger [post]
func createTrigger(c *gin.Context) {
	var trigger *models.Trigger
	if err := c.ShouldBindJSON(&trigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateTrigger(trigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating trigger."})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "created", "data": trigger})
}

// @Summary Test trigger
// @Description Test event trigger
// @Tags Trigger-Test
// @Accept json
// @Produce json
// @Param trigger body Trigger true "Trigger Data"
// @Success 200 {object} Trigger
// @Failure 400 {object} map[string]string
// @Router /api/trigger-test [post]
func testTrigger(c *gin.Context) {
	var trigger *models.Trigger
	if err := c.ShouldBindJSON(&trigger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := services.TestTrigger(trigger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating trigger."})
		return

	}
	c.JSON(http.StatusOK, gin.H{"event trigger": res})
}

// @Summary Edit a trigger
// @Description Creates a trigger that calls an API at a scheduled time
// @Tags Triggers
// @Accept json
// @Produce plain
// @Param triggerId query string true "Enter id to edit trigger"
// @Param trigger body EditTrigger true "Trigger Data"
// @Success 200 {object} EditTrigger
// @Failure 400 {object} map[string]string
// @Router /api/trigger [put]
func editTrigger(c *gin.Context) {
	id := c.Query("triggerId")

	var updateData bson.M

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validUpdate := bson.M{}
	for key, value := range updateData {
		if _, exists := models.AllowedFields[key]; exists {
			validUpdate[key] = value
		}
	}
	if len(validUpdate) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid fields"})
		return
	}

	err := services.EditTrigger(id, validUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated."})
}

// @Summary Delete a trigger by id
// @Description Creates a trigger that calls an API at a scheduled time
// @Tags Triggers
// @Accept plain
// @Produce plain
// @Param triggerId query string true "Enter id for trigger deletion"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/trigger [delete]
func deleteTrigger(c *gin.Context) {
	id := c.Query("triggerId")
	if err := services.DeleteTrigger(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error deletion failed."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Trigger deleted."})
}
