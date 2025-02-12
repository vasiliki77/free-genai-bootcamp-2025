package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/service"
)

// StudyActivityHandler holds the service for study activity related operations.
type StudyActivityHandler struct {
	studyActivityService *service.StudyActivityService
}

// NewStudyActivityHandler creates a new StudyActivityHandler.
func NewStudyActivityHandler(studyActivityService *service.StudyActivityService) *StudyActivityHandler {
	return &StudyActivityHandler{studyActivityService: studyActivityService}
}

// GetStudyActivityHandler handles the GET /api/study_activities/:id endpoint.
func (h *StudyActivityHandler) GetStudyActivityHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study activity ID"})
		return
	}

	activity, err := h.studyActivityService.GetStudyActivity(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get study activity"})
		return
	}
	c.JSON(http.StatusOK, activity)
}

// ListStudyActivitiesHandler handles the GET /api/study_activities endpoint.
func (h *StudyActivityHandler) ListStudyActivitiesHandler(c *gin.Context) {
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	activities, err := h.studyActivityService.ListStudyActivities(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list study activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// GetStudySessionsForActivityHandler handles the GET /api/study_activities/:id/study_sessions endpoint.
func (h *StudyActivityHandler) GetStudySessionsForActivityHandler(c *gin.Context) {
	idStr := c.Param("id")
	activityID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study activity ID"})
		return
	}
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	sessions, err := h.studyActivityService.GetStudySessionsForActivity(activityID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get study sessions for activity"})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

// CreateStudyActivitySessionHandler handles the POST /api/study_activities endpoint.
func (h *StudyActivityHandler) CreateStudyActivitySessionHandler(c *gin.Context) {
	var requestBody struct {
		GroupID         int `json:"group_id" binding:"required"`
		StudyActivityID int `json:"study_activity_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionResponse, err := h.studyActivityService.CreateStudyActivitySession(requestBody.GroupID, requestBody.StudyActivityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create study activity session"})
		return
	}
	c.JSON(http.StatusOK, sessionResponse)
} 