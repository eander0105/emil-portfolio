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

	router = setUpRoutes(router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
