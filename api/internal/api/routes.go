package api

import (
	// STD LIB
	"encoding/json"
	"net/http"

	// INTERNAL LIB
	"emilandersson.se/internal/modules/foobar"

	//EXTERNAL LIB
	"github.com/gorilla/mux"
)

func setUpRoutes(r *mux.Router) *mux.Router {
	// testroutes
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		resp := Healthcheck{Status: true}
		json.NewEncoder(w).Encode(resp)
	})
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Message: "Hello, World!!!"}
		json.NewEncoder(w).Encode(resp)
	}).Methods("GET")

	r.HandleFunc("/foobar", foobar.Foobar).Methods("GET")

	return r
}
