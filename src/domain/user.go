package domain

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
