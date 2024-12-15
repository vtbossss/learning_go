package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<h1>Hello, World!</h1>`,
	)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
