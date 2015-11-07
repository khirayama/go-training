package main

var currentId int

var messages Messages

func init() {
	RepoCreateMessage(Message{Name: "khirayama", Message: "Hello world 1"})
	RepoCreateMessage(Message{Name: "khirayama", Message: "Hello world 2"})
}

func RepoCreateMessage(m Message) Message {
	currentId += 1
	m.Id = currentId
	messages = append(messages, m)
	return m
}
