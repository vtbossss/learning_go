package main

import (
"net/http"
"fmt"
"strconv"
"html/template" // New 
"log"
)


func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/"{
		http.NotFound(w,r)
		return
	}
	ts,err:=template.ParseFiles("./ui/html/pages/home.tmpl")
	if err!=nil{
		log.Println(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
	err=ts.Execute(w,nil)
	if(err!=nil){
		log.Println(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}
}

func handler1(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}
	fmt.Fprintf(w,"Display a specific id %d...",id)


}

func handler2(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost { 
		w.Header().Set("Allow", http.MethodPost) 
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) 
		return 
	}

	w.Write([]byte("Create a new snippet...")) 

}
	

