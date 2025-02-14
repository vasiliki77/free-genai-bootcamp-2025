package service

import (
	"backend_go/internal/models"
	"errors"
)

type GroupService struct{}

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
		groupStats[i] = models.GroupWithStats{
			ID:   group.ID,
			Name: group.Name,
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

	return &models.StudySessionsResponse{
		Items: sessions,
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

func (s *GroupService) GetGroup(id uint) (*models.Group, error) {
	var group models.Group
	result := models.DB.First(&group, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("group not found")
		}
		return nil, result.Error
	}
	return &group, nil
}

func NewGroupService() *GroupService {
	return &GroupService{}
}
