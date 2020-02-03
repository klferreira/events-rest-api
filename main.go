package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crgimenes/goconfig"
	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/api"
)

func main() {
	config := &api.Config{}
	if err := goconfig.Parse(config); err != nil {
		log.Fatal("Could not parse environment vars")
	}

	router := mux.NewRouter()
	server := api.NewServer(config, router)

	fmt.Printf("Events api listening on port %s", config.APIPort)

	log.Fatal(http.ListenAndServe(":3000", server))
}
