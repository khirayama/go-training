package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func KomeIndex(w http.ResponseWriter, r *http.Request) {
	komes := Komes{
		Kome{Name: "コシヒカリ", Url: "aaa", Description: "うまい"},
		Kome{Name: "ヒノヒカリ", Url: "bbb", Description: "これもうまい"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(komes); err != nil {
		panic(err)
	}
}

func KomeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func choice(s []int) int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}
