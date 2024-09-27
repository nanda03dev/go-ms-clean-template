package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	sql_migration "github.com/nanda03dev/go-ms-template/sql-migration"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/router"
)

func main() {
	var databases = ConnectToDBs()
	defer databases.MongoDB.Client.Disconnect(context.Background())
	defer databases.PostgresDB.DB.Close()

	var ginEngine = gin.Default()

	router.InitializeRoutes(ginEngine, databases)

	ginEngine.Run(":3000")

}

func ConnectToDBs() *db.Databases {
	mongoURI := "mongodb://localhost:27017"
	postgresURI := "postgres://nanda:test@localhost/go-ms-template?sslmode=disable"

	databases := db.NewDatabases(db.DatabaseInfo{MongoURI: mongoURI, PostgresURI: postgresURI})

	// Run migrations before starting the application
	if err := sql_migration.RunMigration(databases.PostgresDB.DB); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	return databases
}
