package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/app_module"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/router"
)

func main() {
	var databases = db.GetDBConnection()

	defer databases.MongoDB.Disconnect()
	defer databases.PostgresDB.Disconnect()

	app_module.GetAppModule()

	var ginEngine = gin.Default()

	router.InitializeRoutes(ginEngine)

	ginEngine.Run(":3000")

}
