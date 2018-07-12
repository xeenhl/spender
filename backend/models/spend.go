package models

type Amount struct {
	Amount   float32
	Currency string
}

type User struct {
	ID int32
}

type Spend struct {
	ID     int
	User   User
	Amount *Amount
}
