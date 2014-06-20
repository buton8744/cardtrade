package net

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Mydb struct {
	db *sql.DB
}

var (
	err error
)

func (m *Mydb) Connect(user string, password string, addr string) bool {
	service := user + "@" + password + "tcp(" + addr + ")/lego"
	m.db, err = sql.Open("mysql", service)
	if err != nil {
		Log(err)
		return false
	}
	err = m.db.Ping()
	if err != nil {
		Log(err)
		return false
	}
	return true
	// defer m.db.Close()
}

func (m *Mydb) Close() {
	m.db.Close()
}

func (m *Mydb) Query(query string, args ...interface{}) *sql.Rows {
	rows, err := m.db.Query(query, args)
	if err != nil {
	}
	return rows
}

func (m *Mydb) Execute(query string, args ...interface{}) error {
	result, err := m.db.Exec(query, args)
	if err != nil {
		return err
	}
	return nil
}
