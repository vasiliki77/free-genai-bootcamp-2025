package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// WordParts represents the different forms of a Greek word
type WordParts struct {
	Present string `json:"present"`
	Future  string `json:"future"`
	Aorist  string `json:"aorist"`
	Perfect string `json:"perfect"`
}

// Implement sql.Scanner interface for WordParts
func (wp *WordParts) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, wp)
}

// Implement driver.Valuer interface for WordParts
func (wp WordParts) Value() (driver.Value, error) {
	return json.Marshal(wp)
}

// Base Models
type Word struct {
	gorm.Model
	AncientGreek string            `json:"ancient_greek" gorm:"not null"`
	Greek        string            `json:"greek" gorm:"not null"`
	English      string            `json:"english" gorm:"not null"`
	Parts        WordParts         `json:"parts" gorm:"type:json"`
	Groups       []Group           `json:"groups" gorm:"many2many:words_groups;"`
}

// Add custom JSON marshaling
func (w Word) MarshalJSON() ([]byte, error) {
	type Alias Word
	return json.Marshal(struct {
		ID           uint              `json:"id"`
		AncientGreek string            `json:"ancient_greek"`
		Greek        string            `json:"greek"`
		English      string            `json:"english"`
		Parts        WordParts         `json:"parts"`
		Groups       []GroupWithStats  `json:"groups"`
		Stats        struct {
			CorrectCount int `json:"correct_count"`
			WrongCount   int `json:"wrong_count"`
		} `json:"stats"`
	}{
		ID:           w.ID,
		AncientGreek: w.AncientGreek,
		Greek:        w.Greek,
		English:      w.English,
		Parts:        w.Parts,
		Groups:       []GroupWithStats{}, // TODO: Convert groups
		Stats: struct {
			CorrectCount int `json:"correct_count"`
			WrongCount   int `json:"wrong_count"`
		}{},
	})
}

type Group struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null;unique"`
	Words []Word `json:"words" gorm:"many2many:words_groups;"`
}

type StudySession struct {
	gorm.Model
	GroupID         uint          `json:"group_id" gorm:"not null"`
	Group           Group         `json:"group"`
	StudyActivityID uint          `json:"study_activity_id" gorm:"not null"`
	StudyActivity   StudyActivity `json:"study_activity"`
	WordReviews     []WordReview  `json:"word_reviews"`
}

type StudyActivity struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type WordReview struct {
	gorm.Model
	WordID         uint         `json:"word_id" gorm:"not null"`
	Word           Word         `json:"word"`
	StudySessionID uint         `json:"study_session_id" gorm:"not null"`
	StudySession   StudySession `json:"study_session"`
	Correct        bool         `json:"correct" gorm:"not null"`
}

// Response Types
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
}

type WordWithStats struct {
	ID           uint              `json:"id"`
	AncientGreek string            `json:"ancient_greek"`
	Greek        string            `json:"greek"`
	English      string            `json:"english"`
	Parts        WordParts         `json:"parts"`
	CorrectCount int               `json:"correct_count"`
	WrongCount   int               `json:"wrong_count"`
}

// Add a constructor to convert Word to WordWithStats
func NewWordWithStats(w *Word) *WordWithStats {
	return &WordWithStats{
		ID:           w.ID,
		AncientGreek: w.AncientGreek,
		Greek:        w.Greek,
		English:      w.English,
		Parts: WordParts{
			Present: "",
			Future:  "",
			Aorist:  "",
			Perfect: "",
		},
		CorrectCount: 0, // TODO: Calculate from word reviews
		WrongCount:   0,
	}
}

type WordDetailResponse struct {
	ID           uint   `json:"id"`
	AncientGreek string `json:"ancient_greek"`
	Greek        string `json:"greek"`
	English      string `json:"english"`
	Stats        struct {
		CorrectCount int `json:"correct_count"`
		WrongCount   int `json:"wrong_count"`
	} `json:"stats"`
	Groups []GroupWithStats `json:"groups"`
}

type GroupWithStats struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Stats struct {
		TotalWordCount int `json:"total_word_count,omitempty"`
	} `json:"stats,omitempty"`
}

type StudySessionResponse struct {
	ID               uint      `json:"id"`
	ActivityName     string    `json:"activity_name"`
	GroupName        string    `json:"group_name"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	ReviewItemsCount int       `json:"review_items_count"`
}

type LastStudySessionResponse struct {
	ID              uint      `json:"id"`
	GroupID         uint      `json:"group_id"`
	GroupName       string    `json:"group_name"`
	StudyActivityID uint      `json:"study_activity_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type StudyProgressResponse struct {
	TotalWordsStudied   int `json:"total_words_studied"`
	TotalAvailableWords int `json:"total_available_words"`
}

type QuickStatsResponse struct {
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

type WordsResponse struct {
	Items      []WordWithStats `json:"items"`
	Pagination *Pagination     `json:"pagination"`
}

type GroupsResponse struct {
	Items      []GroupWithStats `json:"items"`
	Pagination *Pagination      `json:"pagination"`
}

type StudyActivitiesResponse struct {
	Items      []StudyActivity `json:"items"`
	Pagination *Pagination     `json:"pagination"`
}

type StudySessionsResponse struct {
	Items      []StudySession `json:"items"`
	Pagination *Pagination    `json:"pagination"`
}

// Database connection
var DB *gorm.DB

func InitDB(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return err
}
