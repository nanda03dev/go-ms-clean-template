package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInfo struct {
	MongoURI    string
	PostgresURI string
}

type Databases struct {
	MongoDB    *MongoDB
	PostgresDB *PostgresDB
}

type MongoDB struct {
	Client *mongo.Client
}

type PostgresDB struct {
	DB *sql.DB
}

func NewDatabases(databaseInfo DatabaseInfo) *Databases {
	mongoDB, err := ConnectMongoDB(databaseInfo.MongoURI)

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	} else {
		fmt.Println("successfully connected to MongoDB")
	}

	postgresDB, err := ConnectPostgresDB(databaseInfo.PostgresURI)
	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	} else {
		fmt.Println("successfully connected to PostgreSQL")
	}

	return &Databases{
		MongoDB:    mongoDB,
		PostgresDB: postgresDB,
	}
}

func (m *MongoDB) Connect(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	m.Client = client
	return nil
}

func (m *MongoDB) Health() (string, bool) {
	// Define a context with a timeout for the ping operation
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Run a ping to the MongoDB server
	err := m.Client.Ping(ctx, nil)

	if err != nil {
		return "MongoDB health check failed", false
	} else {
		return "MongoDB connection is healthy!", true
	}

}

func (p *PostgresDB) Connect(uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	p.DB = db
	return nil
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
