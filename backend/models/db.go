package models

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SpendDatastore interface {
	GetAllSpends(ctx context.Context) ([]*Spend, error)
	GetSpendById(id int, ctx context.Context) (*Spend, error)
	UpdateSpend(id int, newData Spend, ctx context.Context) (*Spend, error)
}

type Datastore interface {
	SpendDatastore
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

	//conStr := s.user + ":" + s.password + "@/" + s.dbName
	//db, err := sql.Open(s.dbType, conStr)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}
	//
	//return &DB{db}, nil

	return nil, nil
}
