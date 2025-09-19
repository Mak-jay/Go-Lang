package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w,"Welcom to my server")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Print(w, "ParseForm() error:",err)
		return
	}
	fmt.Fprintf(w, "POST request succuessfull\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Email = %s\n",email)
}

func main(){

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello",helloHandler)

   fmt.Printf("Server started at port 8080 \n")
   if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
   }


 
}
