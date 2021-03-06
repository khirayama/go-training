package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"

	"../trace"
)

type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}
type chatUser struct {
	common.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します

	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_SECRET_KEY"), "http://localhost:8080/auth/callback/google"),
	)

	// staticファイルを配布する場合
	// http.Handle("/assets/",
	// 	http.StripPrefix("/assets",
	// 		http.FileServer(http.Dir("/assets/"))))

	r := newRoom(UseFileSystemAvatar)
	r.tracer = trace.New(os.Stdout)
	// ルート
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/room", r)
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	// TODO: p77を再読
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars/"))))

	// チャットルームを開始する
	go r.run()

	// Webサーバーを開始します
	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
