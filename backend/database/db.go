package database

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xeenhl/spender/backend/models"
)

type SpendDatastore interface {
	GetAllSpends(ctx context.Context) ([]*models.Spend, error)
	GetSpendById(id int, ctx context.Context) (*models.Spend, error)
	AddSpend(newData models.Spend, ctx context.Context) (*models.Spend, error)
}

type UserDatastore interface {
	GetUserByEmail(email string) (*models.User, error)
	AddNewUser(creds *models.Credentials) (*models.User, error)
}

type Datastore interface {
	SpendDatastore
	UserDatastore
}

type DB struct {
	*sql.DB
}

type DBSettings struct {
	user     string
	password string
	dbName   string
	dbType   string
}

func InitDB() (*DB, error) {

	s := DBSettings{
		user:     "root",
		password: "golang124",
		dbName:   "spenderDB",
		dbType:   "mysql",
	}

	return newDB(s)
}

func newDB(s DBSettings) (*DB, error) {

	conStr := s.user + ":" + s.password + "@/" + s.dbName
	db, err := sql.Open(s.dbType, conStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
