package db

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"backend_go/internal/models"
)

type Migration struct {
	ID       string
	Filename string
	SQL      string
}

type MigrationManager struct {
	migrationsPath string
}

func NewMigrationManager(path string) *MigrationManager {
	return &MigrationManager{
		migrationsPath: path,
	}
}

func (m *MigrationManager) LoadMigrations() ([]Migration, error) {
	var migrations []Migration

	err := filepath.Walk(m.migrationsPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".sql") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", path, err)
		}

		id := strings.TrimSuffix(info.Name(), ".sql")
		migrations = append(migrations, Migration{
			ID:       id,
			Filename: info.Name(),
			SQL:      string(content),
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to load migrations: %w", err)
	}

	// Sort migrations by filename
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].ID < migrations[j].ID
	})

	return migrations, nil
}

func (m *MigrationManager) RunMigrations() error {
	migrations, err := m.LoadMigrations()
	if err != nil {
		return err
	}

	// Create migrations table if it doesn't exist
	err = models.DB.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id TEXT PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`).Error
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Run each migration in a transaction
	for _, migration := range migrations {
		var count int64
		models.DB.Model(&struct{ ID string }{}).
			Table("migrations").
			Where("id = ?", migration.ID).
			Count(&count)

		if count > 0 {
			fmt.Printf("Migration %s already applied, skipping\n", migration.ID)
			continue
		}

		fmt.Printf("Applying migration %s...\n", migration.ID)

		tx := models.DB.Begin()
		if err := tx.Exec(migration.SQL).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to apply migration %s: %w", migration.ID, err)
		}

		if err := tx.Create(struct{ ID string }{ID: migration.ID}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %w", migration.ID, err)
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed to commit migration %s: %w", migration.ID, err)
		}

		fmt.Printf("Successfully applied migration %s\n", migration.ID)
	}

	return nil
}
