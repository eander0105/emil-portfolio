package foobar

import (
	"net/http"
)

func Foobar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foobar"))
}
