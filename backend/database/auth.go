package database

import (
	"errors"
	"fmt"

	"github.com/xeenhl/spender/backend/models"
)

func (db *DB) GetUserByEmail(email string) (*models.User, error) {

	query := fmt.Sprintf("SELECT ID, EMAIL, PASSWORD, VERIFIED,  FROM User WHERE ID = %#v LIMIT 1", email)

	rows, err := db.Query(query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		s := new(models.User)

		err := rows.Scan(&s.ID, &s.Email, &s.Password, &s.Verified)

		if err != nil {
			return nil, err
		}

		return s, nil
	}

	return &models.User{}, errors.New("No user found by email")
}

func (db *DB) AddNewUser(creds *models.Credentials) (*models.User, error) {

	query := fmt.Sprintf(`INSERT INTO User (Email, Verified, Password) VALUES (%#v, %#v, %#v)`, creds.Email, false, creds.Password)

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		u := new(models.User)

		err := rows.Scan(&u.ID, &u.Verified, &u.Email, &u.Password)

		if err != nil {
			return nil, err
		}

		return u, nil
	}

	return nil, errors.New("No User has been added")
}
