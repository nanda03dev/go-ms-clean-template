package db

func ConnectPostgresDB(uri string) (*PostgresDB, error) {
	postgresDB := &PostgresDB{}
	err := postgresDB.Connect(uri)
	if err != nil {
		return nil, err
	}
	return postgresDB, nil
}
