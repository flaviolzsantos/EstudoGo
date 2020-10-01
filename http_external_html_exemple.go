package main

import (
	"html/template"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("files/hello.html"))
	tmpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}
