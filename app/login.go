package app

import (
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am logging in... :)"))
}
