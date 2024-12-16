package main

import(
	"net/http"
	"log"
)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/handler1",handler1)
	mux.HandleFunc("/handler2",handler2)

	log.Println("Starting server at localhost:1000")

	err:=http.ListenAndServe(":1000",mux)
	log.Fatal(err)

}