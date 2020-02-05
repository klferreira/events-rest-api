package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crgimenes/goconfig"
	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/api"
	"github.com/klferreira/events-rest-api/pkg/mongo"
)

func main() {
	config := &api.Config{}
	if err := goconfig.Parse(config); err != nil {
		log.Fatal("Could not parse environment vars")
	}

	db, err := mongo.NewMongoClient(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	server := api.NewServer(db, router)

	fmt.Printf("Events api listening on port %s", config.APIPort)

	log.Fatal(http.ListenAndServe(":3000", server))
}
