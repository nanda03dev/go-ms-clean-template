package service

import (
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type HealthService interface {
	Health() dto.HealthDTO
}

type healthService struct {
	dbs *db.Databases
}

func NewHealthService(dbs *db.Databases) HealthService {
	return &healthService{
		dbs: dbs,
	}
}

func (s *healthService) Health() dto.HealthDTO {
	mongoDBHealth, mongoStatus := s.dbs.MongoDB.Health()
	postgresDBHealth, postgresStatus := s.dbs.PostgresDB.Health()

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
