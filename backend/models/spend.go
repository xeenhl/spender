package models

import "time"

type Amount struct {
	Amount   float32
	Currency string
}

type User struct {
	ID       int32
	Verified bool
	Credentials
}

type Spend struct {
	ID     int
	User   *User
	Amount *Amount
	date   time.Time
}

func (s *Spend) Update(nSpend Spend) {

	s.Amount.update(*nSpend.Amount)
	s.User.update(*nSpend.User)

}

func (a Amount) update(n Amount) {
	a.Amount = n.Amount
	a.Currency = n.Currency

}

func (u *User) update(n User) {
	u.ID = n.ID
}

func NewSpend() *Spend {
	return &Spend{
		ID: -1,
		Amount: &Amount{
			Amount:   -1,
			Currency: "",
		},
		User: &User{
			ID: -1,
		},
	}
}
