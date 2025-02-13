package models

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"
)

// Base Models

// Word represents a vocabulary word.
type Word struct {
	ID           int               `json:"id"`
	AncientGreek string            `json:"ancient_greek"`
	Greek        string            `json:"greek"`
	English      string            `json:"english"`
	Parts        map[string]string `json:"parts"` // JSON map for word parts
}

// Group represents a thematic group of words.
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// StudySession represents a study session.
type StudySession struct {
	ID            int       `json:"id"`
	GroupID       int       `json:"group_id"`
	CreatedAt     time.Time `json:"created_at"`
	StudyActivity int       `json:"study_activity"` // Changed to StudyActivity, assuming it refers to StudyActivity ID
}

// StudyActivity represents a specific study activity.
type StudyActivity struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
	Description  string `json:"description"`
}

// WordReviewItem represents a record of word practice.
type WordReviewItem struct {
	ID             int       `json:"id"`
	WordID         int       `json:"word_id"`
	StudySessionID int       `json:"study_session_id"`
	Correct        bool      `json:"correct"`
	CreatedAt      time.Time `json:"created_at"`
}

// Response Types

// Pagination struct for paginated responses.
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
}

// WordWithStats represents a word with study statistics.
type WordWithStats struct {
	AncientGreek string `json:"ancient_greek"` // Changed to AncientGreek
	Greek        string `json:"greek"`
	English      string `json:"english"`
	CorrectCount int    `json:"correct_count"`
	WrongCount   int    `json:"wrong_count"`
}

// WordDetailResponse represents detailed information for a single word.
type WordDetailResponse struct {
	ID           int               `json:"id"`            // Added ID for consistency
	AncientGreek string            `json:"ancient_greek"` // Changed to AncientGreek
	Greek        string            `json:"greek"`
	English      string            `json:"english"`
	Parts        map[string]string `json:"parts"` // Added Parts
	Stats        struct {
		CorrectCount int `json:"correct_count"`
		WrongCount   int `json:"wrong_count"`
	} `json:"stats"`
	Groups []GroupWithStats `json:"groups"`
}

// GroupWithStats represents a group with associated statistics.
type GroupWithStats struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	WordCount int    `json:"word_count,omitempty"`
	Stats     struct {
		TotalWordCount int `json:"total_word_count,omitempty"`
	} `json:"stats,omitempty"`
}

// StudySessionResponse represents a study session for API responses.
type StudySessionResponse struct {
	ID               int       `json:"id"`
	ActivityName     string    `json:"activity_name"`
	GroupName        string    `json:"group_name"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	ReviewItemsCount int       `json:"review_items_count"`
}

// StudyActivitySessionResponse represents a study activity session for API responses.
type StudyActivitySessionResponse struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
}

// LastStudySession represents data for the last study session.
type LastStudySession struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
	GroupName       string    `json:"group_name"`
}

// StudyProgressResponse represents study progress statistics for dashboard.
type StudyProgressResponse struct {
	TotalWordsStudied   int `json:"total_words_studied"`
	TotalAvailableWords int `json:"total_available_words"`
}

// QuickStatsResponse represents quick statistics for dashboard.
type QuickStatsResponse struct {
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

// WordsResponse is a container for a list of words and pagination info.
type WordsResponse struct {
	Items      []WordWithStats `json:"items"`
	Pagination *Pagination     `json:"pagination"`
}

// GroupsResponse is a container for a list of groups and pagination info.
type GroupsResponse struct {
	Items      []GroupWithStats `json:"items"`
	Pagination *Pagination      `json:"pagination"`
}

// DB struct to hold the database connection.
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection.
func NewDB(dataSourceName string) (*DB, error) {
	fmt.Println("NewDB: Attempting to open database:", dataSourceName)
	absDataSourceName, _ := filepath.Abs(dataSourceName)             // Get absolute path
	fmt.Println("NewDB: Absolute database path:", absDataSourceName) // Log absolute path
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		fmt.Println("NewDB: Error opening database:", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		fmt.Println("NewDB: Database Ping error:", err)
		return nil, err
	}
	fmt.Println("NewDB: Database connection successful.")
	return &DB{db}, nil
}
