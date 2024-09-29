package services

import (
	"github.com/nanda03dev/go-ms-template/src/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type UserService interface {
	CreateUser(createUserDTO dto.CreateUserDTO) (*aggregates.User, error)
	GetUserById(id string) (*aggregates.User, error)
}

type userService struct {
	userRepo aggregates.UserRepository
}

func NewUserService(userRepo aggregates.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(createUserDTO dto.CreateUserDTO) (*aggregates.User, error) {

	newUser := aggregates.NewUser(createUserDTO)

	err := s.userRepo.Save(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *userService) GetUserById(id string) (*aggregates.User, error) {
	return s.userRepo.FindById(id)
}
