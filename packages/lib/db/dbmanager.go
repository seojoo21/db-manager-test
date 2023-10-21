package dbmanager

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

type DBManger struct {
	conn *sql.DB
}

var dbmanager *DBManger

const (
	dbUser   = "userid"
	dbPwd    = "userpwd"
	dbHost   = "localhost:3306"
	dbSchema = "test"
)

func GetConnection() *DBManger {
	once.Do(func() {
		db, err := sql.Open("mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+")/"+dbSchema+"?charset=utf8")
		if err != nil {
			panic(err)
		}

		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		dbmanager = &DBManger{conn: db}
	})

	return dbmanager
}

func (m *DBManger) QueryMultiRows(query string, args ...interface{}) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	if args != nil && args[0] != "" {
		rows, err = m.conn.Query(query, args)
	} else {
		rows, err = m.conn.Query(query)
	}

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m *DBManger) QueryRow(query string, args ...interface{}) *sql.Row {
	var row *sql.Row

	if args != nil && args[0] != "" {
		row = m.conn.QueryRow(query, args...)
	} else {
		row = m.conn.QueryRow(query)
	}

	return row
}

func (m *DBManger) Exec(query string, args ...interface{}) (sql.Result, error) {
	var result sql.Result
	var err error

	if args != nil && args[0] != "" {
		result, err = m.conn.Exec(query, args...)
	} else {
		result, err = m.conn.Exec(query)
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func (m *DBManger) TExec(query string, args ...interface{}) (sql.Result, error) {
	var result sql.Result
	var err error

	tx, err := m.conn.Begin()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if args != nil && args[0] != "" {
		result, err = tx.Exec(query, args...)
	} else {
		result, err = tx.Exec(query)
	}

	if err != nil {
		fmt.Println("err 2 ", err)

		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return result, err
}
