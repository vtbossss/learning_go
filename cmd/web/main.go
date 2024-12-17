package main

import(
	"net/http"
	"log"
)

func main(){

	fileServer:=http.FileServer(http.Dir("./ui/static/"))

	mux:=http.NewServeMux()
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))
	mux.HandleFunc("/",home)
	mux.HandleFunc("/handler1",handler1)
	mux.HandleFunc("/handler2",handler2)

	
	log.Println("Starting server at localhost:1000")

	err:=http.ListenAndServe(":1000",mux)
	log.Fatal(err)

}