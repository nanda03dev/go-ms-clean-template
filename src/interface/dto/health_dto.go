package dto

type HealthResponseDTO struct {
	Service    string `json:"service"`
	MongoDB    string `json:"mongoDB"`
	PostgresDB string `json:"postgresDB"`
}
