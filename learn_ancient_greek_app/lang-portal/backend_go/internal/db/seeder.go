package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"backend_go/internal/models"
)

type SeedManager struct {
	seedPath string
}

type SeedWord struct {
	AncientGreek string `json:"ancient_greek"`
	Greek        string `json:"greek"`
	English      string `json:"english"`
}

func NewSeedManager(path string) *SeedManager {
	return &SeedManager{
		seedPath: path,
	}
}

func (s *SeedManager) LoadSeedFile(filename, groupName string) error {
	data, err := os.ReadFile(filepath.Join(s.seedPath, filename))
	if err != nil {
		return fmt.Errorf("failed to read seed file %s: %w", filename, err)
	}

	var words []SeedWord
	if err := json.Unmarshal(data, &words); err != nil {
		return fmt.Errorf("failed to parse seed file %s: %w", filename, err)
	}

	var group models.Group
	result := models.DB.Where("name = ?", groupName).FirstOrCreate(&group, models.Group{Name: groupName})
	if result.Error != nil {
		return fmt.Errorf("failed to create group %s: %w", groupName, result.Error)
	}

	for _, seedWord := range words {
		word := models.Word{
			AncientGreek: seedWord.AncientGreek,
			Greek:        seedWord.Greek,
			English:      seedWord.English,
		}

		if err := models.DB.Create(&word).Error; err != nil {
			return fmt.Errorf("failed to create word %s: %w", word.AncientGreek, err)
		}

		if err := models.DB.Model(&group).Association("Words").Append(&word); err != nil {
			return fmt.Errorf("failed to associate word with group: %w", err)
		}
	}

	return nil
} 