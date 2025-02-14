package service

import (
	"backend_go/internal/models"
)

type WordService struct {
}

func (s *WordService) GetWords(page, perPage int) (*models.WordsResponse, error) {
	var words []models.Word
	var total int64

	offset := (page - 1) * perPage

	models.DB.Model(&models.Word{}).Count(&total)
	result := models.DB.
		Limit(perPage).
		Offset(offset).
		Find(&words)

	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to WordWithStats
	wordStats := make([]models.WordWithStats, len(words))
	for i, word := range words {
		var correct, wrong int64
		models.DB.Model(&models.WordReview{}).
			Where("word_id = ? AND correct = ?", word.ID, true).
			Count(&correct)
		models.DB.Model(&models.WordReview{}).
			Where("word_id = ? AND correct = ?", word.ID, false).
			Count(&wrong)

		wordStats[i] = models.WordWithStats{
			ID:           word.ID,
			AncientGreek: word.AncientGreek,
			Greek:        word.Greek,
			English:      word.English,
			CorrectCount: int(correct),
			WrongCount:   int(wrong),
		}
	}

	totalPages := (int(total) + perPage - 1) / perPage

	return &models.WordsResponse{
		Items: wordStats,
		Pagination: &models.Pagination{
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   int(total),
			ItemsPerPage: perPage,
		},
	}, nil
}

func (s *WordService) GetWord(id uint) (*models.WordDetailResponse, error) {
	var word models.Word
	result := models.DB.
		Preload("Groups").
		First(&word, id)

	if result.Error != nil {
		return nil, result.Error
	}

	var correct, wrong int64
	models.DB.Model(&models.WordReview{}).
		Where("word_id = ? AND correct = ?", word.ID, true).
		Count(&correct)
	models.DB.Model(&models.WordReview{}).
		Where("word_id = ? AND correct = ?", word.ID, false).
		Count(&wrong)

	groups := make([]models.GroupWithStats, len(word.Groups))
	for i, group := range word.Groups {
		var wordCount int64
		wordCount = models.DB.Model(&group).Association("Words").Count()

		groups[i] = models.GroupWithStats{
			ID:   group.ID,
			Name: group.Name,
			Stats: struct {
				TotalWordCount int `json:"total_word_count,omitempty"`
			}{
				TotalWordCount: int(wordCount),
			},
		}
	}

	return &models.WordDetailResponse{
		ID:           word.ID,
		AncientGreek: word.AncientGreek,
		Greek:        word.Greek,
		English:      word.English,
		Stats: struct {
			CorrectCount int `json:"correct_count"`
			WrongCount   int `json:"wrong_count"`
		}{
			CorrectCount: int(correct),
			WrongCount:   int(wrong),
		},
		Groups: groups,
	}, nil
}

func NewWordService() *WordService {
	return &WordService{}
}
