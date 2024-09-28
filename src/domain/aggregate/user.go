package aggregate

type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

// Business logic inside domain entities
func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

type UserRepository interface {
	Save(user *User) error
	FindById(id string) (*User, error)
}
