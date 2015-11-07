package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db, err = sql.Open("sqlite3", "./db/board.db")

func main() {
	db.Exec("create table message(id integer primary key autoincrement, name text, text text)")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
