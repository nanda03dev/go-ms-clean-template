package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type DatabaseInfo struct {
	MongoURI    string
	PostgresURI string
}

type Databases struct {
	MongoDB    *MongoDB
	PostgresDB *PostgresDB
}

var (
	once               sync.Once
	DatabaseConnection *Databases
	DatabaseURI        DatabaseInfo
)

func init() {
	err := godotenv.Load()

	DatabaseURI = DatabaseInfo{
		MongoURI:    "mongodb://localhost:27017",
		PostgresURI: "postgres://nanda:test@localhost/go-ms-template?sslmode=disable",
	}

	if err != nil {
		log.Println("No .env file found")
	}

	if os.Getenv("MONGO_URI") != "" {
		DatabaseURI.MongoURI = os.Getenv("MONGO_URI")
	}
	if os.Getenv("POSTGRES_URI") != "" {
		DatabaseURI.PostgresURI = os.Getenv("POSTGRES_URI")
	}
}

func GetDBConnection() *Databases {
	once.Do(func() {
		mongoDB, err := ConnectMongoDB(DatabaseURI.MongoURI)

		if err != nil {
			log.Fatalf("Could not connect to MongoDB: %v", err)
		} else {
			fmt.Println("successfully connected to MongoDB")
		}

		postgresDB, err := ConnectPostgresDB(DatabaseURI.PostgresURI)
		if err != nil {
			log.Fatalf("Could not connect to PostgreSQL: %v", err)
		} else {
			fmt.Println("successfully connected to PostgreSQL")
		}

		DatabaseConnection = &Databases{
			MongoDB:    mongoDB,
			PostgresDB: postgresDB,
		}
	})
	return DatabaseConnection
}
