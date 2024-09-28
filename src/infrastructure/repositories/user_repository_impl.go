package repositories

import (
	"errors"

	domain "github.com/nanda03dev/go-ms-template/src/domain/user"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
)

// In-memory storage for simplicity; replace with actual DB connection
var users = map[string]*domain.User{}

type UserRepositoryImpl struct {
	dbs *db.Databases
}

func NewUserRepository(dbs *db.Databases) domain.UserRepository {
	return &UserRepositoryImpl{dbs: dbs}
}

func (r *UserRepositoryImpl) Save(user *domain.User) error {
	users[user.ID] = user
	return nil
}

func (r *UserRepositoryImpl) FindById(id string) (*domain.User, error) {
	user, exists := users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
