package service

import (
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type HealthSearvice struct {
	MongoDB    *db.MongoDB
	PostgresDB *db.PostgresDB
}

func NewHealthService(mongoDB *db.MongoDB, postgresDB *db.PostgresDB) *HealthSearvice {
	return &HealthSearvice{
		MongoDB:    mongoDB,
		PostgresDB: postgresDB,
	}
}

func (s *HealthSearvice) Health() dto.HealthDTO {
	mongoDBHealth, mongoStatus := s.MongoDB.Health()
	postgresDBHealth, postgresStatus := s.PostgresDB.Health()

	var serviceHealth = "go ms template is waiting for requests"

	if !mongoStatus || !postgresStatus {
		serviceHealth = "go ms template is down due to database un-healthy"
	}
	return dto.HealthDTO{
		Service:    serviceHealth,
		MongoDB:    mongoDBHealth,
		PostgresDB: postgresDBHealth,
	}
}
