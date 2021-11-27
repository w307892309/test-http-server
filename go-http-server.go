package main

import (
	"fmt"
	"net/http"
)

func page(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><body>Hello World <br>(build from branch <b>\"DEV\"</b>)</body></html>")
}

func main() {
	http.HandleFunc("/", page)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
