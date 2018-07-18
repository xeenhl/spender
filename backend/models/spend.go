package models

import (
	"context"
	"errors"

	"github.com/xeenhl/spender/backend/authentication"
)

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

func (db *DB) GetAllSpends(ctx context.Context) ([]*Spend, error) {
	spends := make([]*Spend, 0)

	userID, err := getUserID(ctx.Value(authentication.UserID))

	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM Spends WHERE UserID = " + userID
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		s := new(Spend)

		err := rows.Scan(&s.ID, &s.Amount.Amount, &s.Amount.Currency, &s.User.ID)

		if err != nil {
			return nil, err
		}

		spends = append(spends, s)
	}

	return spends, nil

}

func (db *DB) GetSpendById(id int, ctx context.Context) (*Spend, error) {

	// for _, s := range spends {
	// 	if s.ID == id {
	// 		return s, nil
	// 	}
	// }

	return &Spend{}, errors.New("No spend found by ID for Update")
}

func (db *DB) UpdateSpend(id int, newData Spend, ctx context.Context) (*Spend, error) {

	// for _, s := range spends {
	// 	if s.ID == id {
	// 		s.Update(newData)
	// 		return s, nil
	// 	}
	// }

	return &Spend{}, errors.New("No spend found by ID for Update")
}

func getUserID(i interface{}) (string, error) {

	switch v := i.(type) {
	case string:
		return v, nil
	default:
		return "", errors.New("user id mast be string value in context")
	}

}
