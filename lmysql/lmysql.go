package lmysql

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type ldb struct {
	dbMap map[string]*sql.DB
	sync.Mutex
}

var lightDb *ldb = &ldb{}

func init() {
	lightDb.dbMap = make(map[string]*sql.DB)
}

func Conn(dsn string) *sql.DB {
	if conn, ok := lightDb.dbMap[dsn]; ok {
		return conn
	}
	return setConn(dsn)
}

func setConn(dsn string) *sql.DB {
	lightDb.Lock()
	defer lightDb.Unlock()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Print(err.Error())
	}
	lightDb.dbMap[dsn] = conn
	return conn
}
