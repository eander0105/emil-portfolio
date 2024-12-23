package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Healthcheck struct {
	Status bool `json:"status"`
}

type Response struct {
	Message string `json:"message"`
}

func Init() {
	router := mux.NewRouter()
	router.Use(LogMW)

	// router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
	// 	resp := Healthcheck{Status: true}
	// 	json.NewEncoder(w).Encode(resp)
	// })
	// router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	resp := Response{Message: "Hello, World!!!"}
	// 	json.NewEncoder(w).Encode(resp)
	// }).Methods("GET")

	router = setUpRoutes(router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
