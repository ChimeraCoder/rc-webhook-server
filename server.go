package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:generate gojson -o event.go -name "Event" -pkg "main" -input json/push_example.json

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", EventHandler) // TODO fix
	r.HandleFunc("/postreceive", EventHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("error handling event: %s", err)
		fmt.Fprintf(w, "error handling event: %s", err)
		return
	}
	log.Printf("received %+v", event)
	fmt.Fprintf(w, "received %+v", event)
}
