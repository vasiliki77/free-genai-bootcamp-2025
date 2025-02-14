package main

import (
	"log"
	"net/http"

	"backend_go/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := models.InitDB("words.db"); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	r := gin.Default()

	// API routes group
	api := r.Group("/api")
	{
		// Dashboard routes
		api.GET("/dashboard/last_study_session", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"id":                123,
				"group_id":          456,
				"group_name":        "Basic Greetings",
				"study_activity_id": 456,
				"created_at":        "2024-03-20T15:30:00Z",
			})
		})

		api.GET("/dashboard/study_progress", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"total_words_studied":   15,
				"total_available_words": 300,
			})
		})

		api.GET("/dashboard/quick-stats", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success_rate":         80.0,
				"total_study_sessions": 4,
				"total_active_groups":  3,
				"study_streak_days":    4,
			})
		})

		// Study activities routes
		api.GET("/study_activities", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"items": []gin.H{
					{
						"id":            1,
						"name":          "Vocabulary Quiz",
						"description":   "Test your knowledge of Greek vocabulary",
						"thumbnail_url": "https://example.com/thumbnail.jpg",
					},
				},
				"pagination": gin.H{
					"current_page":   1,
					"total_pages":    1,
					"total_items":    1,
					"items_per_page": 100,
				},
			})
		})

		api.GET("/study_activities/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"id":            c.Param("id"),
				"name":          "Vocabulary Quiz",
				"description":   "Test your knowledge of Greek vocabulary",
				"thumbnail_url": "https://example.com/thumbnail.jpg",
			})
		})

		api.POST("/study_activities", func(c *gin.Context) {
			var req struct {
				GroupID         uint `json:"group_id" binding:"required"`
				StudyActivityID uint `json:"study_activity_id" binding:"required"`
			}

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"id":                123,
				"group_id":          req.GroupID,
				"study_activity_id": req.StudyActivityID,
				"created_at":        "2024-03-20T15:30:00Z",
			})
		})

		// Words routes
		api.GET("/words", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"items": []gin.H{
					{
						"id":            123,
						"ancient_greek": "Χαῖρε",
						"greek":         "Γειά",
						"english":       "Hello",
						"parts": gin.H{
							"present": "χαίρω",
							"future":  "χαιρήσω",
							"aorist":  "ἐχάρην",
							"perfect": "κεχάρηκα",
						},
						"correct_count": 5,
						"wrong_count":   2,
					},
				},
				"pagination": gin.H{
					"current_page":   1,
					"total_pages":    5,
					"total_items":    500,
					"items_per_page": 100,
				},
			})
		})

		// Study sessions routes
		api.POST("/study_sessions/:id/words/:word_id/review", func(c *gin.Context) {
			var req struct {
				Correct bool `json:"correct" binding:"required"`
			}

			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"word_id":          c.Param("word_id"),
				"success":          true,
				"study_session_id": c.Param("id"),
				"correct":          req.Correct,
				"created_at":       "2024-03-20T15:30:00Z",
			})
		})

		// System routes
		api.POST("/reset_history", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Study history has been reset",
			})
		})

		api.POST("/full_reset", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Database has been reset to initial state",
			})
		})
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
