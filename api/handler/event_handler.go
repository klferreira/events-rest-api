package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/internal/event"
	"github.com/klferreira/events-rest-api/pkg/httputil"
)

func Fetch(service event.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		events, err := service.Fetch(r.Context(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		response := httputil.GetJSONResponse("events", events, err)

		json.NewEncoder(w).Encode(response)
	})
}

func Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not implemented"))
	})
}

func Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not implemented"))
	})
}

func Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not implemented"))
	})
}

func GetEventHandlers(r *mux.Router, service event.Service) {
	r.Handle("/v1/events", Fetch(service)).Methods(http.MethodGet)
	r.Handle("/v1/events", Create()).Methods(http.MethodPost)
	r.Handle("/v1/events", Update()).Methods(http.MethodPut)
	r.Handle("/v1/events", Delete()).Methods(http.MethodDelete)
}
