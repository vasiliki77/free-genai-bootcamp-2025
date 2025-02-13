package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetWordsHandler handles the /api/words endpoint.
func GetWordsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetWordsHandler: Handling request to /api/words")

	// In a real application, you would fetch words data from your database here.
	// For now, we'll return an empty array as a placeholder.

	response := map[string]interface{}{
		"items":      []interface{}{}, // Empty array of words
		"pagination": map[string]interface{}{ // Placeholder pagination
			"current_page":   1,
			"items_per_page": 100,
			"total_items":    0,
			"total_pages":    1,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("GetWordsHandler: Error encoding JSON response:", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	log.Println("GetWordsHandler: Successfully returned empty words array with pagination")
} 