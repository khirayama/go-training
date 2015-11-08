package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

// template
const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
	Username: <input type="text" id="name" name="name">
	Password: <input type="password" id="password" name="password">
	<button type="submit">Login</button>
</form>
`

const internalPage = `
<h1>Internal</h1>
<small>User: %s</small>
<form method="post" action="/logout">
	<button type="submit">Logout</button>
</form>
`

// Handlers
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexPage)
}

func internalPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		fmt.Fprintf(w, internalPage, userName)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		setSession(name, w)
		redirectTarget = "/internal"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

// helpers
func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
