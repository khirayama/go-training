package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	var messages Messages

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query("SELECT name, message FROM message")

	for rows.Next() {
		var name string
		var text string
		rows.Scan(&name, &text)
		message := Message{}
		message.Name = name
		message.Message = text
		messages = append(messages, message)
	}
	json.NewEncoder(w).Encode(messages)
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message)

	stmt, _ := db.Prepare("INSERT INTO message(name, message) values(?,?)")
	res, _ := stmt.Exec(message.Name, message.Message)
	res.LastInsertId()
}
