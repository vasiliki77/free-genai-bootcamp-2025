package main

import (
	"log"

	"backend_go/internal/models"
	"backend_go/internal/service"
	"backend_go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := models.InitDB("words.test.db"); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize services
	wordService := service.NewWordService()
	groupService := service.NewGroupService()
	systemService := service.NewSystemService()
	dashboardService := service.NewDashboardService()
	studyService := service.NewStudyService()

	// Initialize handlers
	handlers := handlers.NewHandlers(
		dashboardService,
		wordService,
		studyService,
		groupService,
		systemService,
	)

	// Setup router
	router := gin.Default()
	handlers.Register(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
