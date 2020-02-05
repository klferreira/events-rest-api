package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/internal/event"
	"github.com/klferreira/events-rest-api/internal/model"
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

func Create(service event.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		event := &model.Event{}

		err = json.Unmarshal(body, event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		event, err = service.Create(r.Context(), event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		httputil.GetJSONResponse("event", event, nil).Write(w)
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
	r.Handle("/v1/events", Create(service)).Methods(http.MethodPost)
	r.Handle("/v1/events", Update()).Methods(http.MethodPut)
	r.Handle("/v1/events", Delete()).Methods(http.MethodDelete)
}
