package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var SqlitePool *sql.DB

func InitPool() {
	var err error
	SqlitePool, err = sql.Open("sqlite3", "file:todo.db")
	if err != nil {
		panic("Cannot connect to todo.db")
	}
}
