package client

import (
	"encoding/json"
	"fmt"
	"github.com/byt3-m3/go-zipper/internal/models"

	"log"
	"net/http"
)

type Config struct {
	zipFetcherHost string
}

type opt func(cfg *Config)

var (
	WithHost = func(host string) opt {
		return func(cfg *Config) {
			cfg.zipFetcherHost = host
		}
	}
)

type ZipFetcherClient interface {
	GetAddressData(zip string) models.AddressData
}

type client struct {
	httpClient http.Client
	*Config
}

func NewClient(opts ...opt) ZipFetcherClient {
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	return client{
		httpClient: http.Client{},
		Config:     cfg,
	}
}

func (c client) GetAddressData(zip string) models.AddressData {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.zipFetcherHost, fmt.Sprintf("/api/v1/zip?=%s", zip)), nil)
	if err != nil {

	}

	resp, err := c.httpClient.Do(req)

	var data models.AddressData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalln(err)
	}

	return data

}
