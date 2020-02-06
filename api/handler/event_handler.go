package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/api/validator"
	"github.com/klferreira/events-rest-api/internal/event"
	"github.com/klferreira/events-rest-api/pkg/httputil"
)

func Fetch(service event.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		filters, err := validator.FetchEventRequestForm(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		events, err := service.Fetch(r.Context(), filters)
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

		event, err := validator.ValidateEventCreateRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		result, err := service.Create(r.Context(), event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		w.WriteHeader(http.StatusCreated)
		httputil.GetJSONResponse("event", result, nil).Write(w)
	})
}

func Update(service event.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		event, err := validator.ValidateEventUpdateRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		result, err := service.Update(r.Context(), event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		httputil.GetJSONResponse("event", result, nil).Write(w)
	})
}

func Delete(service event.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
			httputil.GetJSONResponse("event", nil, errors.New("missing or invalid event id"))
			return
		}

		_, err := service.Delete(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			httputil.GetJSONResponse("event", nil, err).Write(w)
			return
		}
	})
}

func GetEventHandlers(r *mux.Router, service event.Service) {
	r.Handle("/v1/events", Fetch(service)).Methods(http.MethodGet)
	r.Handle("/v1/events", Create(service)).Methods(http.MethodPost)
	r.Handle("/v1/events", Update(service)).Methods(http.MethodPut)
	r.Handle("/v1/events/{id}", Delete(service)).Methods(http.MethodDelete)
}
