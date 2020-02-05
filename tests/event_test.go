package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/api"
	"github.com/klferreira/events-rest-api/internal/model"
	"github.com/klferreira/events-rest-api/pkg/mongo"
)

var db mongo.Client
var server *api.Server

var errorDb mongo.Client
var errorServer *api.Server

func getTestDBClient(url string) mongo.Client {
	db, err := mongo.NewMongoClient(url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getServer() *api.Server {

	router := mux.NewRouter()

	return api.NewServer(db, router)
}

func getErrorServer() *api.Server {
	testConfig := &api.Config{}
	db, err := mongo.NewMongoClient(testConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	return api.NewServer(db, router)
}

func setupDB(db mongo.Client) error {
	events := []interface{}{
		&model.Event{
			ID:    bson.NewObjectId(),
			Name:  "Random party",
			Place: "Random venue",
			Tags:  []string{"electronic", "dance"},
			Sessions: []time.Time{
				time.Now().Add(72 * time.Hour),
				time.Now().Add(48 * time.Hour),
			},
		},
		&model.Event{
			ID:    bson.NewObjectId(),
			Name:  "AC/DC Live Performance",
			Place: "Another random venue",
			Tags:  []string{"rock", "classic-rock", "live"},
			Sessions: []time.Time{
				time.Now().Add(24 * time.Hour),
			},
		},
	}

	return db.Insert("events", events...)
}

func tearDown(db mongo.Client) error {
	return db.DeleteAll("events", nil)
}

func init() {
	config := &api.Config{
		DatabaseURL: "mongodb://root:toor@localhost:27017/test?authSource=admin",
	}

	errConfig := &api.Config{}

	db = getTestDBClient(config.DatabaseURL)
	errorDb = getTestDBClient(errConfig.DatabaseURL)

	server = getServer()
	errorServer = getErrorServer()
}

func TestMain(m *testing.M) {
	setupDB(db)
	code := m.Run()
	tearDown(db)

	os.Exit(code)
}

func TestFetchEvents(t *testing.T) {
	t.Run("should successfully fetch all events", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/events", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should receive an error fetching events", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/events", nil)
		response := httptest.NewRecorder()

		errorServer.ServeHTTP(response, request)

		got := response.Result().StatusCode
		want := http.StatusInternalServerError

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestCreateEvent(t *testing.T) {
	t.Run("should successfully create a new event", func(t *testing.T) {

		event := model.Event{
			Name:  "Lollapalooza 2020",
			Place: "Autodromo de Interlagos",
			Tags:  []string{"festival", "pop", "rock"},
			Sessions: []time.Time{
				time.Now().Add(24 * time.Hour),
				time.Now().Add(48 * time.Hour),
			},
		}

		bs, _ := json.Marshal(event)

		request := httptest.NewRequest(http.MethodPost, "/v1/events", bytes.NewReader(bs))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusCreated {
			t.Errorf("status code expected was %d, but got %d", http.StatusCreated, response.Code)
			t.FailNow()
		}
	})

	t.Run("should fail at creating a new event", func(t *testing.T) {
		event := model.Event{
			Place: "Autodromo de Interlagos",
			Tags:  []string{"festival", "pop", "rock"},
			Sessions: []time.Time{
				time.Now().Add(24 * time.Hour),
				time.Now().Add(48 * time.Hour),
			},
		}

		bs, _ := json.Marshal(event)

		request := httptest.NewRequest(http.MethodPost, "/v1/events", bytes.NewReader(bs))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("status code expected was %d, but got %d", http.StatusBadRequest, response.Code)
			t.FailNow()
		}
	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("should update an existing event", func(t *testing.T) {
		event := &model.Event{}

		err := db.FindOne("events", bson.M{}, event)
		if err != nil {
			t.Error("couldn't find a documento to be updated")
		}

		event.Name = "Updated event"

		bs, _ := json.Marshal(event)

		request := httptest.NewRequest(http.MethodPut, "/v1/events", bytes.NewReader(bs))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("status code expected was %d, but got %d", http.StatusOK, response.Code)
			t.FailNow()
		}

		bresp, err := ioutil.ReadAll(response.Result().Body)
		if err != nil {
			t.Errorf("could not read response body")
			t.FailNow()
		}

		updateResponse := map[string]map[string]*model.Event{}
		err = json.Unmarshal(bresp, &updateResponse)
		if err != nil {
			t.Errorf("could not unmarshal response body")
			t.FailNow()
		}

		updatedEvent := updateResponse["data"]["event"]

		got := updatedEvent.Name
		want := event.Name

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("should fail at updating an event", func(t *testing.T) {
		event := &model.Event{
			Name:  "New Year Celebration",
			Place: "Madison Square Garden",
			Tags:  []string{"new-year"},
		}

		bs, _ := json.Marshal(event)

		request := httptest.NewRequest(http.MethodPut, "/v1/events", bytes.NewReader(bs))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("status code expected was %d, but got %d", http.StatusBadRequest, response.Code)
			t.FailNow()
		}
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("should successfully delete an event", func(t *testing.T) {
		event := &model.Event{}

		err := db.FindOne("events", bson.M{}, event)
		if err != nil {
			t.Error("couldn't find a documento to be updated")
		}

		url := fmt.Sprintf("/v1/events/%s", event.ID.Hex())
		request := httptest.NewRequest(http.MethodDelete, url, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("status code expected was %d, but got %d", http.StatusOK, response.Code)
			t.FailNow()
		}

		err = db.FindOne("events", bson.M{"_id": event.ID}, &model.Event{})
		if err != nil {
			t.Errorf("event with id %s was not expected to be found", event.ID.Hex())
		}
	})

	t.Run("should fail at deleting an event", func(t *testing.T) {
		event := &model.Event{}

		request := httptest.NewRequest(http.MethodDelete, "/v1/events/notvalid", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("status code expected was %d, but got %d", http.StatusOK, response.Code)
			t.FailNow()
		}

		err := db.FindOne("events", bson.M{"_id": event.ID}, &model.Event{})
		if err != nil {
			t.Errorf("event with id %s was not expected to be found", event.ID.Hex())
		}
	})
}
