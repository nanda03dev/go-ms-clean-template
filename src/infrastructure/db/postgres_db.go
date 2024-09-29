package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type PostgresDB struct {
	DB *sql.DB
}

func ConnectPostgresDB(uri string) (*PostgresDB, error) {
	postgresDB := &PostgresDB{}
	err := postgresDB.Connect(uri)
	if err != nil {
		return nil, err
	}

	// Run migrations before starting the application
	if err := RunMigration(postgresDB.DB); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	return postgresDB, nil
}

func (p *PostgresDB) Connect(uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	p.DB = db
	return nil
}

func (p *PostgresDB) Disconnect() {
	p.DB.Close()
}

func (p *PostgresDB) Health() (string, bool) {
	// Set a deadline for the ping operation
	p.DB.SetConnMaxLifetime(2 * time.Second)

	// Ping the PostgreSQL server
	err := p.DB.Ping()
	if err != nil {
		return "PostgreSQL health check failed", false
	} else {
		return "PostgreSQL connection is healthy!", true
	}
}

func RunMigration(db *sql.DB) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %w", err)
	}

	migrationPath := wd + "/sql-migrations/"

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
