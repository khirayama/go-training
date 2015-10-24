package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		// return nil, err ?
		return nil
	}
	// return &Page{Title: title, Body: body}[, nil ?
	return &Page{Title: title, Body: body}
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	// p2, _ := loadPage("TestPage") ?
	p2 := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
