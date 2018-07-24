package models

import (
	"fmt"
	"errors"
)

type Credentials struct {
	Email    string
	Password string
}

func (db *DB) GetUserByEmail(email string) (*User, error) {

	query := fmt.Sprintf("SELECT * FROM User WHERE ID = %#v LIMIT 1", email)

	rows, err := db.Query(query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		s := new(User)

		err := rows.Scan(&s.ID, &s.Verified, &s.Email)

		if err != nil {
			return nil, err
		}

		return s, nil
	}

	return &User{}, errors.New("No user found by email")
}

func (db *DB) AddNewUser(user *User) (*User, error) {

	query := fmt.Sprintf(`INSERT INTO Spends (Email, Verified, Password) VALUES (%#v, %#v, %#v)`, user.Email, user.Verified, user.Password)

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		u := new(User)

		err := rows.Scan(&u.ID, &u.Verified)

		if err != nil {
			return nil, err
		}

		return u, nil
	}

	return nil, errors.New("No spend found by ID for Update")
}
