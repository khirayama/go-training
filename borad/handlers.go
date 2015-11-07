package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	var messages Messages

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query("SELECT id, name, text FROM message")

	for rows.Next() {
		var id int
		var name string
		var text string
		rows.Scan(&id, &name, &text)
		message := Message{}
		message.Id = id
		message.Name = name
		message.Text = text
		messages = append(messages, message)
	}
	json.NewEncoder(w).Encode(messages)
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message Message
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	json.Unmarshal(body, &message)

	stmt, _ := db.Prepare("INSERT INTO message(name, text) values(?, ?)")
	stmt.Exec(message.Name, message.Text)
}

func MessageDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageId := vars["messageId"]
	stmt, _ := db.Prepare("delete from message where id=?")
	stmt.Exec(messageId)
}
