package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klferreira/events-rest-api/api/handler"
	"github.com/klferreira/events-rest-api/pkg/mongo"
)

type Server struct {
	db     mongo.Client
	router *mux.Router
}

type Config struct {
	DatabaseURL string `cfg:"DATABASE_URL" cfgDefault:"mongodb://root:toor@localhost:27017/events?authSource=admin" cfgRequired:"true"`
	APIPort     string `cfg:"API_PORT" cfgDefault:"3000" cfgRequired:"true"`
}

func NewServer(config *Config, r *mux.Router) *Server {
	db, err := mongo.NewMongoClient(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	handler.GetEventHandlers(r)

	return &Server{db, r}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
