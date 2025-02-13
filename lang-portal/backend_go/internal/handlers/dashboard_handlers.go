package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// DashboardQuickStatsHandler handles the /api/dashboard/quick-stats endpoint.
func DashboardQuickStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DashboardQuickStatsHandler: Handling request to /api/dashboard/quick-stats")

	// In a real application, you would fetch quick stats data from your backend here.
	// For now, we'll return a placeholder response.

	response := map[string]interface{}{
		"totalWords":       100, // Placeholder values
		"totalGroups":      5,   // Placeholder values
		"studySessionsToday": 2,   // Placeholder values
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("DashboardQuickStatsHandler: Error encoding JSON response:", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	log.Println("DashboardQuickStatsHandler: Successfully returned quick stats")
} 