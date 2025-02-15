package service

import (
	"backend_go/internal/models"
	"errors"
)

type StudyService struct {
}

func NewStudyService() *StudyService {
	return &StudyService{}
}

func (s *StudyService) GetStudyActivities(page, perPage int) (*models.StudyActivitiesResponse, error) {
	var activities []models.StudyActivity
	var total int64

	offset := (page - 1) * perPage

	models.DB.Model(&models.StudyActivity{}).Count(&total)
	result := models.DB.
		Limit(perPage).
		Offset(offset).
		Find(&activities)

	if result.Error != nil {
		return nil, result.Error
	}

	totalPages := (int(total) + perPage - 1) / perPage

	return &models.StudyActivitiesResponse{
		Items: activities,
		Pagination: &models.Pagination{
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   int(total),
			ItemsPerPage: perPage,
		},
	}, nil
}

func (s *StudyService) GetStudyActivity(id uint) (*models.StudyActivity, error) {
	var activity models.StudyActivity
	result := models.DB.First(&activity, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("activity not found")
		}
		return nil, result.Error
	}

	return &activity, nil
}

func (s *StudyService) CreateStudySession(groupID, studyActivityID uint) (*models.StudySession, error) {
	session := &models.StudySession{
		GroupID:         groupID,
		StudyActivityID: studyActivityID,
	}

	result := models.DB.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func (s *StudyService) ReviewWord(sessionID, wordID uint, correct bool) (*models.WordReview, error) {
	review := &models.WordReview{
		WordID:         wordID,
		StudySessionID: sessionID,
		Correct:        correct,
	}

	result := models.DB.Create(review)
	if result.Error != nil {
		return nil, result.Error
	}

	return review, nil
}

func (s *StudyService) GetActivityStudySessions(activityID uint, page, perPage int) (*models.StudySessionsResponse, error) {
	var sessions []models.StudySession
	var total int64

	query := models.DB.Model(&models.StudySession{}).Where("study_activity_id = ?", activityID)
	query.Count(&total)

	result := query.Limit(perPage).Offset((page - 1) * perPage).Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to StudySessionWithStats
	sessionStats := make([]models.StudySessionWithStats, len(sessions))
	for i, session := range sessions {
		sessionStats[i] = models.StudySessionWithStats{
			ID:              session.ID,
			GroupID:         session.GroupID,
			CreatedAt:       session.CreatedAt,
			StudyActivityID: session.StudyActivityID,
		}
	}

	return &models.StudySessionsResponse{
		Items: sessionStats,
		Pagination: &models.Pagination{
			CurrentPage:  page,
			TotalPages:   (int(total) + perPage - 1) / perPage,
			TotalItems:   int(total),
			ItemsPerPage: perPage,
		},
	}, nil
}

func (s *StudyService) GetStudySessions(page, perPage int) (*models.StudySessionsResponse, error) {
	var sessions []models.StudySession
	var total int64

	query := models.DB.Model(&models.StudySession{})
	query.Count(&total)

	result := query.Limit(perPage).Offset((page - 1) * perPage).Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to StudySessionWithStats
	sessionStats := make([]models.StudySessionWithStats, len(sessions))
	for i, session := range sessions {
		sessionStats[i] = models.StudySessionWithStats{
			ID:              session.ID,
			GroupID:         session.GroupID,
			StudyActivityID: session.StudyActivityID,
			CreatedAt:       session.CreatedAt,
		}
	}

	return &models.StudySessionsResponse{
		Items: sessionStats,
		Pagination: &models.Pagination{
			CurrentPage:  page,
			TotalPages:   (int(total) + perPage - 1) / perPage,
			TotalItems:   int(total),
			ItemsPerPage: perPage,
		},
	}, nil
}

func (s *StudyService) GetStudySession(id uint) (*models.StudySessionWithStats, error) {
	var session models.StudySession
	result := models.DB.First(&session, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("session not found")
		}
		return nil, result.Error
	}

	return &models.StudySessionWithStats{
		ID:              session.ID,
		GroupID:         session.GroupID,
		StudyActivityID: session.StudyActivityID,
		CreatedAt:       session.CreatedAt,
	}, nil
}

func (s *StudyService) GetStudySessionWords(sessionID uint, page, perPage int) (*models.WordsResponse, error) {
	var session models.StudySession
	if err := models.DB.First(&session, sessionID).Error; err != nil {
		return nil, err
	}

	var words []models.Word
	var total int64

	// Get words through reviews
	query := models.DB.Model(&models.Word{}).
		Joins("JOIN word_reviews ON words.id = word_reviews.word_id").
		Where("word_reviews.study_session_id = ?", sessionID)

	query.Count(&total)
	result := query.Limit(perPage).Offset((page - 1) * perPage).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to WordWithStats
	wordStats := make([]models.WordWithStats, len(words))
	for i, word := range words {
		wordStats[i] = *models.NewWordWithStats(&word)
	}

	return &models.WordsResponse{
		Items: wordStats,
		Pagination: &models.Pagination{
			CurrentPage:  page,
			TotalPages:   (int(total) + perPage - 1) / perPage,
			TotalItems:   int(total),
			ItemsPerPage: perPage,
		},
	}, nil
}
