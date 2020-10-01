package main

import (
	"html/template"
	"net/http"
)

type Teste struct {
	Name string
}

func helloTemplate(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("files/hello_template.html"))

	testes := []Teste{
		{Name: "Teste1"},
		{Name: "Teste2"},
		{Name: "Teste3"},
	}

	tmpl.Execute(res, testes)
}

func main() {
	http.HandleFunc("/", helloTemplate)
	http.ListenAndServe(":8888", nil)
}
