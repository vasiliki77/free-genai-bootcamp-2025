package service

import (
	"backend_go/internal/models"
	"errors"
	"time"
)

type GroupService struct{}

type GroupWithStats struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	WordCount int    `json:"word_count"`
}

type GroupResponse struct {
	ID    uint `json:"id"`
	Name  string `json:"name"`
	Stats struct {
		TotalWordCount int `json:"total_word_count"`
	} `json:"stats"`
}

type StudySessionWithStats struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID uint      `json:"study_activity_id"`
}

func (s *GroupService) GetGroups(page, perPage int) (*models.GroupsResponse, error) {
	var groups []models.Group
	var total int64

	models.DB.Model(&models.Group{}).Count(&total)
	result := models.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&groups)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert to GroupWithStats
	groupStats := make([]models.GroupWithStats, len(groups))
	for i, group := range groups {
		wordCount := models.DB.Model(&group).Association("Words").Count()
		
		groupStats[i] = models.GroupWithStats{
			ID:        group.ID,
			Name:      group.Name,
			WordCount: int(wordCount),
		}
	}

	return &models.GroupsResponse{
		Items: groupStats,
		Pagination: &models.Pagination{
			CurrentPage:  page,
			TotalPages:   (int(total) + perPage - 1) / perPage,
			TotalItems:   int(total),
			ItemsPerPage: perPage,
		},
	}, nil
}

func (s *GroupService) GetGroupStudySessions(groupID uint, page, perPage int) (*models.StudySessionsResponse, error) {
	var sessions []models.StudySession
	var total int64

	result := models.DB.Where("group_id = ?", groupID).Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("group not found")
	}

	// Convert to StudySessionWithStats
	sessionStats := make([]models.StudySessionWithStats, len(sessions))
	for i, session := range sessions {
		sessionStats[i] = models.StudySessionWithStats{
			ID:              session.ID,
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

func (s *GroupService) GetGroupWords(groupID uint, page, perPage int) (*models.WordsResponse, error) {
	var words []models.Word
	var total int64

	result := models.DB.Joins("JOIN words_groups ON words.id = words_groups.word_id").
		Where("words_groups.group_id = ?", groupID).
		Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("group not found")
	}

	// Convert to WordWithStats
	wordStats := make([]models.WordWithStats, len(words))
	for i, word := range words {
		wordStats[i] = models.WordWithStats{
			ID:           word.ID,
			AncientGreek: word.AncientGreek,
			Greek:        word.Greek,
			English:      word.English,
		}
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

func (s *GroupService) GetGroup(id uint) (*GroupResponse, error) {
	var group models.Group
	result := models.DB.First(&group, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("group not found")
		}
		return nil, result.Error
	}

	wordCount := models.DB.Model(&group).Association("Words").Count()

	response := &GroupResponse{
		ID:   group.ID,
		Name: group.Name,
	}
	response.Stats.TotalWordCount = int(wordCount)

	return response, nil
}

func NewGroupService() *GroupService {
	return &GroupService{}
}
