package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Message: "Hello, World!!"}
		json.NewEncoder(w).Encode(resp)
		log.Println("Hello, World!")
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8090", router))
}
