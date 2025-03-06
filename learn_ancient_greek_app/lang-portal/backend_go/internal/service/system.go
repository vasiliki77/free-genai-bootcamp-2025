package service

import (
	"backend_go/internal/models"
	"fmt"
)

type SystemService struct{}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (s *SystemService) ResetHistory() error {
	// Delete all data from relevant tables
	if err := models.DB.Exec(`
		DELETE FROM word_reviews;
		DELETE FROM study_sessions;
		DELETE FROM words;
		DELETE FROM groups;
		DELETE FROM study_activities;
		
		-- Reset SQLite sequences
		DELETE FROM sqlite_sequence WHERE name IN 
		('word_reviews', 'study_sessions', 'words', 'groups', 'study_activities');
	`).Error; err != nil {
		return err
	}
	return nil
}

func (s *SystemService) FullReset() error {
	// Delete all data from relevant tables
	if err := models.DB.Exec(`
		DELETE FROM word_reviews;
		DELETE FROM study_sessions;
		DELETE FROM words;
		DELETE FROM groups;
		DELETE FROM study_activities;
		
		-- Reset SQLite sequences
		DELETE FROM sqlite_sequence WHERE name IN 
		('word_reviews', 'study_sessions', 'words', 'groups', 'study_activities');
	`).Error; err != nil {
		return err
	}
	return nil
}

func (s *SystemService) ReloadTestData() error {
	fmt.Println("Reloading test data...")

	// First clear existing data
	if err := s.ResetHistory(); err != nil {
		fmt.Println("Error resetting history:", err)
		return err
	}

	// Then reload test data with explicit IDs
	if err := models.DB.Exec(`
		-- Restore core data first
		INSERT INTO groups (id, name, created_at, updated_at) VALUES 
		(1, 'Basic Verbs', DATETIME('now'), DATETIME('now')),
		(2, 'Common Nouns', DATETIME('now'), DATETIME('now'));

		INSERT INTO study_activities (id, name, description, created_at, updated_at) VALUES 
		(1, 'Vocabulary Quiz', 'Test your vocabulary', DATETIME('now'), DATETIME('now'));

		INSERT INTO words (id, ancient_greek, greek, english, parts, created_at, updated_at) VALUES 
		(1, 'λύω', 'λύνω', 'to loose', '{}', DATETIME('now'), DATETIME('now')),
		(2, 'γράφω', 'γράφω', 'to write', '{}', DATETIME('now'), DATETIME('now')),
		(3, 'χαίρω', 'χαίρω', 'to rejoice', '{}', DATETIME('now'), DATETIME('now'));

		-- Then add study data
		INSERT INTO study_sessions (id, group_id, study_activity_id, created_at, updated_at) VALUES
		(1, 1, 1, DATETIME('now'), DATETIME('now')),
		(2, 2, 1, DATETIME('now'), DATETIME('now')),
		(3, 1, 1, DATETIME('now', '-1 day'), DATETIME('now', '-1 day'));

		INSERT INTO word_reviews (id, word_id, study_session_id, correct, created_at, updated_at) VALUES
		(1, 1, 1, true, DATETIME('now'), DATETIME('now')),
		(2, 2, 1, true, DATETIME('now'), DATETIME('now')),
		(3, 3, 1, true, DATETIME('now'), DATETIME('now')),
		(4, 1, 2, false, DATETIME('now'), DATETIME('now'));
	`).Error; err != nil {
		fmt.Println("Error reloading data:", err)
		return err
	}

	fmt.Println("Test data reloaded successfully")
	return nil
}
