package service

import (
	"github.com/lang-portal/backend_go/internal/models"
)

// StudyService handles general study related business logic.
type StudyService struct {
	db *models.DB
}

// NewStudyService creates a new StudyService.
func NewStudyService(db *models.DB) *StudyService {
	return &StudyService{db: db}
}

// ResetHistory resets study history.
func (s *StudyService) ResetHistory() error {
	// TODO: Implement logic to reset study history
	return nil
}

// FullReset performs a full reset of the database.
func (s *StudyService) FullReset() error {
	// TODO: Implement logic for full database reset
	return nil
} 