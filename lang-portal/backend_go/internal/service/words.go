package service

import (
	"github.com/lang-portal/backend_go/internal/models"
)

// WordService handles word related business logic.
type WordService struct {
	db *models.DB
}

// NewWordService creates a new WordService.
func NewWordService(db *models.DB) *WordService {
	return &WordService{db: db}
}

// GetWord retrieves a word by ID.
func (s *WordService) GetWord(id int) (*models.WordDetailResponse, error) {
	// TODO: Implement database query to get a word by ID
	return nil, nil
}

// ListWords retrieves a paginated list of words.
func (s *WordService) ListWords(page int, pageSize int) (*models.WordsResponse, error) {
	// TODO: Implement database query to get a paginated list of words
	return nil, nil
} 