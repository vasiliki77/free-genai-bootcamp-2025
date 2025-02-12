package service

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lang-portal/backend_go/internal/models"
)

// DashboardService handles dashboard related business logic.
type DashboardService struct {
	db *models.DB
}

// NewDashboardService creates a new DashboardService.
func NewDashboardService(db *models.DB) *DashboardService {
	return &DashboardService{db: db}
}

// LastStudySession represents data for the last study session.
type LastStudySession struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
	GroupName       string    `json:"group_name"`
}

// StudyProgress represents study progress statistics.
type StudyProgress struct {
	TotalWordsStudied   int `json:"total_words_studied"`
	TotalAvailableWords int `json:"total_available_words"`
}

// QuickStats represents quick statistics for the dashboard.
type QuickStats struct {
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
	WordsLearned       int     `json:"words_learned"`
	WordsInProgress    int     `json:"words_in_progress"`
}

// GetLastStudySession retrieves the last study session from the database.
func (s *DashboardService) GetLastStudySession() (*LastStudySession, error) {
	query := `
		SELECT
			ss.id, ss.group_id, ss.created_at, ss.study_activity, g.name
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		ORDER BY ss.created_at DESC
		LIMIT 1
	`

	log.Printf("Executing query: %s", query)
	row := s.db.QueryRow(query)
	err := row.Err()
	if err != nil {
		log.Printf("QueryRow error: %v", err)
		return nil, err
	}

	var session LastStudySession
	err = row.Scan(
		&session.ID,
		&session.GroupID,
		&session.CreatedAt,
		&session.StudyActivityID,
		&session.GroupName,
	)
	if err != nil {
		log.Printf("Error during Scan: %v", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no study sessions found")
	}
	return &session, nil
}

// GetStudyProgress retrieves study progress statistics.
func (s *DashboardService) GetStudyProgress() (*StudyProgress, error) {
	query := `
		SELECT
			COUNT(DISTINCT word_id) as studied,
			(SELECT COUNT(*) FROM words) as total
		FROM word_review_items
	`

	var progress StudyProgress
	err := s.db.QueryRow(query).Scan(
		&progress.TotalWordsStudied,
		&progress.TotalAvailableWords,
	)
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

// GetQuickStats retrieves summary statistics for the learning progress.
func (s *DashboardService) GetQuickStats() (*QuickStats, error) {
	// Get success rate and word counts
	successRateQuery := `
		SELECT
			CAST(COALESCE(CAST(SUM(CASE WHEN correct THEN 1 ELSE 0 END) AS REAL) / NULLIF(CAST(COUNT(*) AS REAL), 0) * 100, 0.0) AS REAL),
			COALESCE(COUNT(DISTINCT CASE WHEN correct THEN word_id END), 0),
			COALESCE(COUNT(DISTINCT CASE WHEN NOT correct THEN word_id END), 0)
		FROM word_review_items
	`

	// Get total study sessions
	sessionsQuery := `SELECT COUNT(*) FROM study_sessions`

	// Get total active groups
	groupsQuery := `
		SELECT COUNT(DISTINCT group_id)
		FROM study_sessions
		WHERE created_at >= datetime('now', '-30 days')
	`

	// Get study streak (simplified version - counts consecutive days)
	streakQuery := `
		WITH RECURSIVE dates AS (
			SELECT date(created_at) as study_date
			FROM study_sessions
			GROUP BY date(created_at)
			ORDER BY study_date DESC
		),
		streak AS (
			SELECT study_date, 1 as streak
			FROM dates
			WHERE study_date = date('now', '-1 day')
			UNION ALL
			SELECT d.study_date, s.streak + 1
			FROM dates d
			JOIN streak s ON date(d.study_date, '+1 day') = s.study_date
		)
		SELECT COALESCE(MAX(streak), 0) FROM streak
	`

	var stats QuickStats
	if err := s.db.QueryRow(successRateQuery).Scan(&stats.SuccessRate, &stats.WordsLearned, &stats.WordsInProgress); err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(sessionsQuery).Scan(&stats.TotalStudySessions); err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(groupsQuery).Scan(&stats.TotalActiveGroups); err != nil {
		return nil, err
	}
	if err := s.db.QueryRow(streakQuery).Scan(&stats.StudyStreakDays); err != nil {
		return nil, err
	}

	return &stats, nil
}
