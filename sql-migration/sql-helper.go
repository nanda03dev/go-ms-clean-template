package sql_migration

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
)

func RunMigration(db *sql.DB) error {
	wd, err := os.Getwd()

	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	migrationPath := wd + "/sql-migration/migrations/"

	files, err := os.ReadDir(migrationPath)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			// Open the migration file
			f, err := os.Open(migrationPath + file.Name())
			if err != nil {
				return fmt.Errorf("failed to open migration file %s: %w", file.Name(), err)
			}
			defer f.Close()

			// Read the content of the migration file
			script, err := io.ReadAll(f)
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
			}

			_, err = db.Exec(string(script))
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
			}
			log.Printf("Executed migration: %s", file.Name())
		}
	}

	return nil
}
