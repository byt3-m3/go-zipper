package server

import (
	"github.com/gorilla/mux"
	"go-zip/internal/models"
	"go-zip/internal/vars"
	"log"
	"net/http"
)

type WebServer interface {
	Run() error
}

type Opt func(c *Config)

func WithListenAddress(address string) Opt {
	return func(c *Config) {
		c.ListenAddress = address
	}
}

type Config struct {
	ListenAddress string
}

type server struct {
	router  *mux.Router
	config  *Config
	fetcher models.ZipFetcher
}

func NewServer(opts ...Opt) WebServer {
	cfg := &Config{}

	for _, opt := range opts {
		opt(cfg)
	}

	s := server{
		router:  mux.NewRouter(),
		config:  cfg,
		fetcher: models.NewZipFetcher(models.WithZipFileLocation(vars.ZipFileLocation)),
	}

	s.registerRoutes()

	return s
}

func (s server) registerRoutes() {
	s.router.HandleFunc("/api/v1/zip", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("received request:", request.RemoteAddr, request.Method, request.URL)

		zipCode := request.URL.Query().Get("zip")

		data := s.fetcher.GetAddressData(zipCode)

		BuildJSONResponseStatusOkWBody(writer, data)

	})
}

func (s server) Run() error {
	log.Println("starting server:", s.config.ListenAddress)
	return http.ListenAndServe(s.config.ListenAddress, s.router)
}
