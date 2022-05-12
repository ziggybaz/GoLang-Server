package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(res http.ResponseWriter,req *http.Request){
	if err := req.ParseForm();err !=nil{
		fmt.Fprintf(res,"ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(res,"Your POST request was succesful \n")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res,"Name = %s\n",name)
	fmt.Fprintf(res,"Address = %s\n",address)
}

func helloHandler( res http.ResponseWriter,req *http.Request){
	if req.URL.Path !="/hello"{
		http.Error(res, "404 NOT FOUND", http.StatusNotFound)
		return
	}
	if req.Method !="GET"{
		http.Error(res,"Method Not Supported", http.StatusNotFound)
	}
	fmt.Fprintf(res,"Hello Bro")
}//takes in two params, the reponse and the request

func main() {
	fileServer := http.FileServer(http.Dir("Static")) //using http package to inform Golang to read static files
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	//listening port - server code
	fmt.Printf("Starting on PORT 3000\n")
	if err:=http.ListenAndServe(":3000",nil); err !=nil{
		log.Fatal(err)
	}
}