package dao

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type SQLHelper struct {
	db *sql.DB
}

func NewSQLHelper(driver, connstr string) (*SQLHelper, error) {
	db, err := sql.Open(driver, connstr)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS article (
			id varchar(128),
			name varchar(128) NOT NULL,
			author varchar(128) NOT NULL,
			brief TEXT NOT NULL,
			content TEXT NOT NULL,
			PRIMARY KEY (id)
			)`)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}
	logrus.Debug("database initialized")
	return &SQLHelper{db:db}, nil
}