package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type Healthcheck struct {
	Status bool `json:"status"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()
	// router.Use(LogMW)

	router = setUpRoutes(router)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func setUpRoutes(r *mux.Router) *mux.Router {
	// testroutes
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		resp := Healthcheck{Status: true}
		json.NewEncoder(w).Encode(resp)
	})

	return r
}
