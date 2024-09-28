package db

func ConnectMongoDB(uri string) (*MongoDB, error) {
	mongoDB := &MongoDB{}
	err := mongoDB.Connect(uri)
	if err != nil {
		return nil, err
	}
	return mongoDB, nil
}
