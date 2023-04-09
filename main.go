package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	fileserver := http.FileServer(http.Dir("./sites"))
	http.Handle("/" , fileserver)
	http.HandleFunc("/greet" , greetingFunction)
	http.HandleFunc("/info" , infoFunction)


	fmt.Println("Starting at port number: 8080")

	if err := http.ListenAndServe(":8080" , nil); err != nil {
		log.Fatal(err)	
	} 

}

func greetingFunction(w http.ResponseWriter , r *http.Request){
	if r.URL.Path != "/greet" {
		fmt.Fprintf(w, "Error 404: not found")
	}
	if r.Method != "GET" {
		fmt.Fprintf(w, "This method is Not expected")
	}
	fmt.Fprintf(w , "HELLO VIEWER!!!")
}

func infoFunction(w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() method error")
		return
	}
	fmt.Fprintf(w , "Post request Successful\n")
	names := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w , "Name : %s\n" , names)
	fmt.Fprintf(w , "email : %s" , email)
}