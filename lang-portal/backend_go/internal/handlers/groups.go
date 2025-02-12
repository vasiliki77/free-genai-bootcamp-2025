package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/service"
)

// GroupHandler holds the service for group related operations.
type GroupHandler struct {
	groupService *service.GroupService
}

// NewGroupHandler creates a new GroupHandler.
func NewGroupHandler(groupService *service.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

// GetGroupHandler handles the GET /api/groups/:id endpoint.
func (h *GroupHandler) GetGroupHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := h.groupService.GetGroup(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get group"})
		return
	}
	c.JSON(http.StatusOK, group)
}

// ListGroupsHandler handles the GET /api/groups endpoint.
func (h *GroupHandler) ListGroupsHandler(c *gin.Context) {
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	groups, err := h.groupService.ListGroups(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list groups"})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// GetWordsInGroupHandler handles the GET /api/groups/:id/words endpoint.
func (h *GroupHandler) GetWordsInGroupHandler(c *gin.Context) {
	idStr := c.Param("id")
	groupID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	words, err := h.groupService.GetWordsInGroup(groupID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get words in group"})
		return
	}
	c.JSON(http.StatusOK, words)
}

// GetStudySessionsForGroupHandler handles the GET /api/groups/:id/study_sessions endpoint.
func (h *GroupHandler) GetStudySessionsForGroupHandler(c *gin.Context) {
	idStr := c.Param("id")
	groupID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	sessions, err := h.groupService.GetStudySessionsForGroup(groupID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get study sessions for group"})
		return
	}
	c.JSON(http.StatusOK, sessions)
} 