package app

import (
	"net/http"
)

func AddRoutes(mux *http.ServeMux) {
	mux.Handle("/", handlerRoot())
	mux.Handle("/healthz", handleHealthz())
	mux.Handle("/login", handleLogin())
}
