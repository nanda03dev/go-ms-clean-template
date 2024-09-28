package domain

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

// Business logic inside domain entities
func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}
