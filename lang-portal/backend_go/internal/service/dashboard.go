package service

import (
	"backend_go/internal/models"
	"time"
)

type DashboardService struct {
}

func (s *DashboardService) GetLastStudySession() (*models.LastStudySessionResponse, error) {
	var session models.StudySession
	result := models.DB.
		Preload("Group").
		Order("created_at DESC").
		First(&session)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.LastStudySessionResponse{
		ID:              session.ID,
		GroupID:         session.GroupID,
		GroupName:       session.Group.Name,
		StudyActivityID: session.StudyActivityID,
		CreatedAt:       session.CreatedAt,
	}, nil
}

func (s *DashboardService) GetStudyProgress() (*models.StudyProgressResponse, error) {
	var totalWords int64
	var studiedWords int64

	models.DB.Model(&models.Word{}).Count(&totalWords)
	models.DB.Model(&models.WordReview{}).Distinct("word_id").Count(&studiedWords)

	return &models.StudyProgressResponse{
		TotalWordsStudied:   int(studiedWords),
		TotalAvailableWords: int(totalWords),
	}, nil
}

func (s *DashboardService) GetQuickStats() (*models.QuickStatsResponse, error) {
	var totalSessions int64
	var totalGroups int64
	var correctReviews int64
	var totalReviews int64

	models.DB.Model(&models.StudySession{}).Count(&totalSessions)
	models.DB.Model(&models.Group{}).Count(&totalGroups)
	models.DB.Model(&models.WordReview{}).Count(&totalReviews)
	models.DB.Model(&models.WordReview{}).Where("correct = ?", true).Count(&correctReviews)

	var successRate float64
	if totalReviews > 0 {
		successRate = float64(correctReviews) / float64(totalReviews) * 100
	}

	// Calculate study streak
	var streak int
	var lastSession models.StudySession
	result := models.DB.Order("created_at DESC").First(&lastSession)
	if result.Error == nil {
		streakDate := lastSession.CreatedAt
		for {
			if !hasStudySessionOnDate(streakDate) {
				break
			}
			streak++
			streakDate = streakDate.AddDate(0, 0, -1)
		}
	}

	return &models.QuickStatsResponse{
		SuccessRate:        successRate,
		TotalStudySessions: int(totalSessions),
		TotalActiveGroups:  int(totalGroups),
		StudyStreakDays:    streak,
	}, nil
}

func hasStudySessionOnDate(date time.Time) bool {
	var count int64
	models.DB.Model(&models.StudySession{}).
		Where("DATE(created_at) = DATE(?)", date).
		Count(&count)
	return count > 0
}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}
