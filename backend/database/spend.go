package database

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	appCtx "github.com/xeenhl/spender/backend/context"
	"github.com/xeenhl/spender/backend/models"
)

func (db *DB) GetAllSpends(ctx context.Context) ([]*models.Spend, error) {

	spends := make([]*models.Spend, 0)
	userID := getID(ctx)

	query := "SELECT * FROM Spends WHERE UserID = " + userID
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		s := models.NewSpend()

		err := rows.Scan(&s.ID, &s.Amount.Amount, &s.Amount.Currency, &s.User.ID)

		if err != nil {
			return nil, err
		}

		spends = append(spends, s)
	}

	return spends, nil

}

func (db *DB) GetSpendById(id int, ctx context.Context) (*models.Spend, error) {

	userID := getID(ctx)

	query := fmt.Sprintf(`SELECT * FROM Spends WHERE UserID = %#v AND ID = %#v LIMIT 1`, userID, strconv.Itoa(id))
	fmt.Println(query)
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		s := models.NewSpend()

		err := rows.Scan(&s.ID, &s.Amount.Amount, &s.Amount.Currency, &s.User.ID)

		if err != nil {
			return nil, err
		}

		return s, nil
	}

	return &models.Spend{}, errors.New("No spend found by ID for Update")
}

func (db *DB) AddSpend(newData models.Spend, ctx context.Context) (*models.Spend, error) {

	userID := getID(ctx)

	if userID != strconv.Itoa(int(newData.User.ID)) {
		return nil, errors.New("cand add spend for wrong userId")
	}

	query := fmt.Sprintf(`INSERT INTO Spends (Amount, Currency, UserId) VALUES (%#v, %#v, %#v)`, newData.Amount.Amount, newData.Amount.Currency, newData.User.ID)
	fmt.Println(query)
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		s := models.NewSpend()

		err := rows.Scan(&s.ID, &s.Amount.Amount, &s.Amount.Currency, &s.User.ID)

		if err != nil {
			return nil, err
		}

		return s, nil
	}

	return &models.Spend{}, errors.New("No spend found by ID for Update")
}

func getID(context context.Context) string {
	var userID int
	x := context.Value(appCtx.UserID)

	switch v := x.(type) {
	//jwt-go lib parse int claims from token as float64 by default
	case float64:
		userID = int(v)
	default:
		userID = -1
	}

	return strconv.Itoa(userID)
}
