package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "simple http exemple")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}
