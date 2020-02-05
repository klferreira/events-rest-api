package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/api"
	"github.com/klferreira/events-rest-api/pkg/mongo"
)

var server *api.Server
var errorServer *api.Server

func getServer() *api.Server {
	testConfig := &api.Config{
		DatabaseURL: "mongodb://root:toor@localhost:27017/test?authSource=admin",
	}

	db, err := mongo.NewMongoClient(testConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

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

func TestMain(t *testing.M) {

	server = getServer()
	errorServer = getErrorServer()

	os.Exit(t.Run())
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
