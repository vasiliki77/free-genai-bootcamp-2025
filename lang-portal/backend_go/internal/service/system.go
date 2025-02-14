package service

import (
	"backend_go/internal/models"
)

type SystemService struct{}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (s *SystemService) ResetHistory() error {
	return models.DB.Delete(&models.WordReview{}).Error
}

func (s *SystemService) FullReset() error {
	// Delete all data from all tables
	if err := models.DB.Delete(&models.WordReview{}).Error; err != nil {
		return err
	}
	if err := models.DB.Delete(&models.StudySession{}).Error; err != nil {
		return err
	}
	if err := models.DB.Delete(&models.Word{}).Error; err != nil {
		return err
	}
	return models.DB.Delete(&models.Group{}).Error
} 