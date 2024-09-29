package entity

type User struct {
	UserID   string `bson:"userID,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
