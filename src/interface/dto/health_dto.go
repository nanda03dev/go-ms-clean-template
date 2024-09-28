package dto

type HealthDTO struct {
	Service    string `json:"service"`
	MongoDB    string `json:"mongoDB"`
	PostgresDB string `json:"postgresDB"`
}
