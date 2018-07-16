package models

import "errors"

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

func (s *Spend) Update(nSpend Spend) {

	s.Amount.update(*nSpend.Amount)
	s.User.update(nSpend.User)

}

func (a *Amount) update(n Amount) {

	a.Amount = n.Amount
	a.Currency = n.Currency

}

func (u *User) update(n User) {
	u.ID = n.ID
}

func (db *DB) GetAllSpends() ([]*Spend, error) {
	return make([]*Spend, 10), nil
}

func (db *DB) GetSpendById(id int) (*Spend, error) {

	// for _, s := range spends {
	// 	if s.ID == id {
	// 		return s, nil
	// 	}
	// }

	return &Spend{}, errors.New("No spend found by ID for Update")
}

func (db *DB) UpdateSpend(id int, newData Spend) (*Spend, error) {

	// for _, s := range spends {
	// 	if s.ID == id {
	// 		s.Update(newData)
	// 		return s, nil
	// 	}
	// }

	return &Spend{}, errors.New("No spend found by ID for Update")
}
