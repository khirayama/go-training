# Writing Web Applications
https://golang.org/doc/articles/wiki/

## imports
fmt
io/ioutil
http -> net/http
html/template

## methods

- fmt.Printf
- fmt.Println
- fmt.Fprintf
- ioutil.WriteFile
- ioutil.ReadFile
- http.HandleFunc
- http.ListenAndServer
- template.Execte

## description
- type
- struct
- []byte
- *
- &
- :=
- [1:] // ex) ```Path[1:]``` : Path内の1番目の文字から最後までの部分スライスを作成する
- w, r // ex) ```func handler(w http.ResponseWriter, r *http.Request) {```
- : // ex) ```r.URL.Path[len("/edit/"):]```

