package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	sql_migration "github.com/nanda03dev/go-ms-template/sql-migration"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/router"
)

func main() {
	var mongoDB, postgresDB = ConnectToDBs()
	defer mongoDB.Client.Disconnect(context.Background())
	defer postgresDB.DB.Close()

	var ginEngine = gin.Default()

	router.InitializeRoutes(ginEngine, mongoDB, postgresDB)

	ginEngine.Run(":3000")

}

func ConnectToDBs() (*db.MongoDB, *db.PostgresDB) {
	mongoURI := "mongodb://localhost:27017"
	postgresURI := "postgres://nanda:test@localhost/go-ms-template?sslmode=disable"

	mongoDB, err := db.ConnectMongoDB(mongoURI)

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	} else {
		fmt.Println("successfully connected to MongoDB")
	}

	postgresDB, err := db.ConnectPostgresDB(postgresURI)
	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	} else {
		fmt.Println("successfully connected to PostgreSQL")
	}

	// Run migrations before starting the application
	if err := sql_migration.RunMigration(postgresDB.DB); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	return mongoDB, postgresDB
}
