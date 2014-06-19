package net

import (
	"container/list"
	"database/sql"
	"github.com/go-sql-drive/mysql"
)

type Mydb struct {
	db sql.DB
}

func (m *Mydb) Connect(user, password, addr) bool {
	service := user + "@" + password + "tcp(" + addr + ")/lego"
	m.db, err := sql.Open("mysql", service)
	if err != nil {
		Log(err)
		return false
	}
	err := m.db.Ping()
	if err != nil {
		Log(err)
		return false
	}
	return true
	// defer m.db.Close()
}

func (m *Mydb) Query(query string, args ...interface{}) sql.Rows {
	rows, err := m.db.Query(query, args)
}

func (m *Mydb) Execute(query string, args ..interface{}) bool {
	err := m.db.Execute(query, args)
}
