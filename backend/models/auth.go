package models

import (
	"errors"
	"fmt"
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

func (db *DB) AddNewUser(creds Credentials) (*User, error) {

	query := fmt.Sprintf(`INSERT INTO User (Email, Verified, Password) VALUES (%#v, %#v, %#v)`, creds.Email, false, creds.Password)

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		u := new(User)

		err := rows.Scan(&u.ID, &u.Verified, &u.Email, &u.Password)

		if err != nil {
			return nil, err
		}

		return u, nil
	}

	return nil, errors.New("No User has been added")
}
