package main

import (
	"net/http"

	// _ "github.com/Dhanraj-Patil/event-trigger/config"

	"github.com/Dhanraj-Patil/event-trigger/internal/database"
	"github.com/Dhanraj-Patil/event-trigger/internal/handler"
	"github.com/Dhanraj-Patil/event-trigger/internal/scheduler"
	"github.com/Dhanraj-Patil/event-trigger/internal/utils"

	_ "github.com/Dhanraj-Patil/event-trigger/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Event-Trigger
// @version 1.0
// @description This is a sample API documentation using Swagger in Go.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	utils.InitLogger()
	database.InitDB()
	scheduler.InitAsynqClient()
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler.EventRouter(router)

	router.GET("/", home)

	router.Run(":8080")
}

// Home godoc
// @Summary Ping the server
// @Description Health check endpoint
// @Tags Health
// @Produce  plain
// @Success 200 {string} string "Event Trigger Platform running."
// @Router / [get]
func home(c *gin.Context) {
	c.String(http.StatusOK, "Event Trigger Platform running.")
}
