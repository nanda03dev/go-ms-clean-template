package service

import (
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type HealthService interface {
	Health() dto.HealthDTO
}

type healthService struct {
	databases *db.Databases
}

func NewHealthService(databases *db.Databases) HealthService {
	return &healthService{
		databases: databases,
	}
}

func (s *healthService) Health() dto.HealthDTO {
	mongoDBHealth, mongoStatus := s.databases.MongoDB.Health()
	postgresDBHealth, postgresStatus := s.databases.PostgresDB.Health()

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
