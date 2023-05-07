package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/api",func(w http.ResponseWriter, r *http.Request){
		log.Println("entering ")
		fmt.Fprintf(w, "Hi!")
	})

  log.Fatal(http.ListenAndServe(":8080", nil))
}