# Writing Web Applications
https://golang.org/doc/articles/wiki/
http://www.slideshare.net/yasi_life/goweb-16448500

## CLI

```
$ go run wiki.go

$ go build wiki.go
$ ./wiki

$ go fmt wiki.go
```

## imports
fmt
io/ioutil
http -> net/http
html/template

## methods

- range
- fmt.Printf
- fmt.Println
- fmt.Fprintf
- ioutil.WriteFile
- ioutil.ReadFile
- http.HandleFunc
- http.ListenAndServer
- http.Redirect
- http.StatusFound
- http.StatusInternalServerError
- http.Error
- http.URL
- http.Request.FormValue
- template.Execte
- template.ExecteTemplate
- template.Must

## description
- type
- struct
- []byte
- *
- &
- := // 定義 & 代入？
- = // 代入
- [1:] // ex) ```Path[1:]``` : Path内の1番目の文字から最後までの部分スライスを作成する
- w, r // ex) ```func handler(w http.ResponseWriter, r *http.Request) {```

