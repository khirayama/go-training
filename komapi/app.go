package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type Kome struct {
	name        string
	url         string
	description string
}

func randomKomeInfo() *Kome {
	kome := &Kome{
		name:        "コシヒカリ",
		url:         "aaa",
		description: "うまい",
	}
	return kome
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	// kome := json.Marshal()
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		kome := randomKomeInfo()
		log.Print(kome)
		w.WriteJson(kome)
	}))
	log.Print("app start")
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
