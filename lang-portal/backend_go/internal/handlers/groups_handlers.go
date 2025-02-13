package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetGroupsHandler handles the /api/groups endpoint.
func GetGroupsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetGroupsHandler: Handling request to /api/groups")

	// In a real application, you would fetch groups data from your database here.
	// For now, we'll return an empty array as a placeholder.

	response := []interface{}{} // Empty array of groups

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("GetGroupsHandler: Error encoding JSON response:", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	log.Println("GetGroupsHandler: Successfully returned empty groups array")
} 