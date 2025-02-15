package service

import (
	"backend_go/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type DashboardService struct {
}

type LastStudySessionResponse struct {
	ID              uint   `json:"id"`
	GroupID         uint   `json:"group_id"`
	GroupName       string `json:"group_name"`
	StudyActivityID uint   `json:"study_activity_id"`
	CreatedAt       string `json:"created_at"`
}

type StudyProgressResponse struct {
	TotalWordsStudied   int `json:"total_words_studied"`
	TotalAvailableWords int `json:"total_available_words"`
}

type QuickStatsResponse struct {
	SuccessRate        json.Number `json:"success_rate"`
	TotalStudySessions int         `json:"total_study_sessions"`
	TotalActiveGroups  int         `json:"total_active_groups"`
	StudyStreakDays    int         `json:"study_streak_days"`
}

func (s *DashboardService) GetLastStudySession() (*models.LastStudySessionResponse, error) {
	var session models.StudySession
	var group models.Group

	result := models.DB.Order("created_at DESC").First(&session)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("no study sessions found")
		}
		return nil, result.Error
	}

	if err := models.DB.First(&group, session.GroupID).Error; err != nil {
		return nil, err
	}

	return &models.LastStudySessionResponse{
		ID:              session.ID,
		GroupID:         session.GroupID,
		GroupName:       group.Name,
		StudyActivityID: session.StudyActivityID,
		CreatedAt:       session.CreatedAt,
	}, nil
}

func (s *DashboardService) GetStudyProgress() (*StudyProgressResponse, error) {
	var totalWords int64
	var studiedWords int64

	models.DB.Model(&models.Word{}).Count(&totalWords)
	models.DB.Model(&models.WordReview{}).Distinct("word_id").Count(&studiedWords)

	return &StudyProgressResponse{
		TotalWordsStudied:   int(studiedWords),
		TotalAvailableWords: int(totalWords),
	}, nil
}

func (s *DashboardService) GetQuickStats() (*models.QuickStatsResponse, error) {
	var totalSessions int64
	var totalGroups int64
	var correctReviews int64
	var totalReviews int64

	// Count total study sessions
	models.DB.Model(&models.StudySession{}).Count(&totalSessions)

	// Count active groups (groups with study sessions)
	models.DB.Model(&models.Group{}).
		Joins("JOIN study_sessions ON study_sessions.group_id = groups.id").
		Distinct().Count(&totalGroups)

	// Calculate success rate
	models.DB.Model(&models.WordReview{}).Where("correct = ?", true).Count(&correctReviews)
	models.DB.Model(&models.WordReview{}).Count(&totalReviews)

	var successRate float64
	if totalReviews > 0 {
		successRate = float64(correctReviews) / float64(totalReviews) * 100.0
	}

	return &models.QuickStatsResponse{
		SuccessRate:        json.Number(fmt.Sprintf("%.2f", successRate)),
		TotalStudySessions: int(totalSessions),
		TotalActiveGroups:  int(totalGroups),
		StudyStreakDays:    s.calculateStreak(),
	}, nil
}

func hasStudySessionOnDate(date time.Time) bool {
	var count int64
	models.DB.Model(&models.StudySession{}).
		Where("DATE(created_at) = DATE(?)", date).
		Count(&count)
	return count > 0
}

func (s *DashboardService) calculateStreak() int {
	var lastSession models.StudySession

	// Get last study session
	result := models.DB.Order("created_at desc").First(&lastSession)
	if result.Error != nil {
		return 0
	}

	streak := 0
	today := time.Now()

	// Check consecutive days backwards from today
	for i := 0; i < 365; i++ { // Cap at 1 year
		checkDate := today.AddDate(0, 0, -i)
		if !hasStudySessionOnDate(checkDate) {
			break
		}
		streak++
	}

	return streak
}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}
