package service

import (
	"backend_go/internal/models"
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
