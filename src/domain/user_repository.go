package domain

type UserRepository interface {
	Save(user *User) error
	FindById(id string) (*User, error)
}
