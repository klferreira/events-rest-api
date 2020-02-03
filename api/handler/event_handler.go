package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Fetch() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not implemented"))
	})
}

func Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not implemented"))
	})
}

func Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not implemented"))
	})
}

func Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not implemented"))
	})
}

func GetEventHandlers(r *mux.Router) {
	r.Handle("/v1/events", Fetch()).Methods(http.MethodGet)
	r.Handle("/v1/events", Create()).Methods(http.MethodPost)
	r.Handle("/v1/events", Update()).Methods(http.MethodPut)
	r.Handle("/v1/events", Delete()).Methods(http.MethodDelete)
}
