package main

import (
	"fmt"
	"net/http"
)

func page(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World <br>(build from branch \"master\")\n")
}

func main() {
	http.HandleFunc("/", page)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
