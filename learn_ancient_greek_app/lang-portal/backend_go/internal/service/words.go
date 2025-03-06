package service

import (
	"backend_go/internal/models"
	"errors"
)

type WordService struct {
}

func (s *WordService) GetWords(page, perPage int) (*models.WordsResponse, error) {
	var words []models.Word
	var total int64

	models.DB.Model(&models.Word{}).Count(&total)
	result := models.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to WordWithStats
	wordStats := make([]models.WordWithStats, len(words))
	for i, word := range words {
		stats := models.NewWordWithStats(&word)
		wordStats[i] = *stats
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

type WordResponse struct {
	ID           uint              `json:"id"`
	AncientGreek string            `json:"ancient_greek"`
	Greek        string            `json:"greek"`
	English      string            `json:"english"`
	Parts        map[string]string `json:"parts"`
	Stats        struct {
		CorrectCount int `json:"correct_count"`
		WrongCount   int `json:"wrong_count"`
	} `json:"stats"`
	Groups []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
}

func (s *WordService) GetWord(id uint) (*WordResponse, error) {
	var word models.Word
	result := models.DB.Preload("Groups").First(&word, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("word not found")
		}
		return nil, result.Error
	}

	// Convert to response format
	response := &WordResponse{
		ID:           word.ID,
		AncientGreek: word.AncientGreek,
		Greek:        word.Greek,
		English:      word.English,
		Parts: map[string]string{
			"present": word.Parts.Present,
			"future":  word.Parts.Future,
			"aorist":  word.Parts.Aorist,
			"perfect": word.Parts.Perfect,
		},
		Stats: struct {
			CorrectCount int `json:"correct_count"`
			WrongCount   int `json:"wrong_count"`
		}{
			CorrectCount: 0, // TODO: Calculate from reviews
			WrongCount:   0,
		},
	}

	// Convert groups
	response.Groups = make([]struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}, len(word.Groups))
	for i, g := range word.Groups {
		response.Groups[i].ID = g.ID
		response.Groups[i].Name = g.Name
	}

	return response, nil
}

func NewWordService() *WordService {
	return &WordService{}
}
