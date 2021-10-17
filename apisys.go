package main

import(
	"fmt"
	"net/http"
	"log"
)

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "API system works")
}

func handleRequests(){
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main(){
	handleRequests()
}