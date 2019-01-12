package main

import (
	"io"
	"net/http"
)

func handle(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hola servidor go!")
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}
