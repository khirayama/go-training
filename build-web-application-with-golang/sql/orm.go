package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "test/xiemengjun/123456")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)

	type Userinfo struct {
		Uid        int `PK`
		Username   string
		Departname string
		Created    time.Time
	}
}
