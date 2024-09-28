package domain

type Order struct {
	ID       string
	UserID   int
	ItemName string
	Amount   float64
}
