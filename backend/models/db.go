package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SpendDatastore interface {
	GetAllSpends() ([]*Spend, error)
	GetSpendById(id int) (*Spend, error)
	UpdateSpend(id int, newData Spend) (*Spend, error)
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

	conStr := s.user + ":" + s.password + "@/" + s.dbName
	fmt.Println(conStr)
	db, err := sql.Open(s.dbType, conStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
