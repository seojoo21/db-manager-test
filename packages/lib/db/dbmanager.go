package dbmanager

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

type DBManger struct {
	conn *sql.DB
}

var dbmanager *DBManger

func GetConnection() *DBManger {
	once.Do(func() {
		db, err := sql.Open("mysql", "userid:password@tcp(localhost:3306)/testschema?charset=utf8")
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

func (m *DBManger) QueryMultiRows(query string) (*sql.Rows, error) {
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
