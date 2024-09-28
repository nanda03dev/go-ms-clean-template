package service

import (
	"github.com/google/uuid"
	"github.com/nanda03dev/go-ms-template/src/domain/aggregate"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type UserService interface {
	CreateUser(createUserDTO dto.CreateUserDTO) (*aggregate.User, error)
	GetUserById(id string) (*aggregate.User, error)
}

type userService struct {
	userRepo aggregate.UserRepository
}

func NewUserService(userRepo aggregate.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(createUserDTO dto.CreateUserDTO) (*aggregate.User, error) {
	newUser := &aggregate.User{
		UserID:   Generate16DigitUUID(), // Generate unique ID (UUID or similar)
		Name:     createUserDTO.Name,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	}

	err := s.userRepo.Save(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *userService) GetUserById(id string) (*aggregate.User, error) {
	return s.userRepo.FindById(id)
}

func Generate16DigitUUID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
