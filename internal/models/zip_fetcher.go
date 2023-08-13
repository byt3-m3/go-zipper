package models

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

type opt func(c *Config)

var (
	WithZipFileLocation = func(location string) opt {

		return func(c *Config) {
			c.ZipFileLocation = location
		}
	}
)

type Config struct {
	ZipFileLocation string
}

type AddressData struct {
	Zip       string `csv:"zip"`
	City      string `csv:"city"`
	StateID   string `csv:"state_id"`
	StateName string `csv:"state_name"`
}

type ZipFetcher interface {
	GetAddressData(zipCode string) AddressData
}

type zipFetcher struct {
	ZipRecords []AddressData
	*Config
}

func NewZipFetcher(opts ...opt) ZipFetcher {
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	f := &zipFetcher{Config: cfg}

	f.loadData()

	return f
}

func (f *zipFetcher) loadData() {
	if f.ZipFileLocation == "" {
		log.Fatalln("must set ZipFileLocation")
	}

	file, err := os.Open(f.ZipFileLocation)
	if err != nil {
		log.Fatalln(err)
	}

	err = gocsv.Unmarshal(file, &f.ZipRecords)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("zip data loaded")
}

func (f *zipFetcher) GetAddressData(zipCode string) AddressData {

	for _, record := range f.ZipRecords {
		if record.Zip == zipCode {
			log.Println("found", record)
			return record
		}

	}

	log.Println(zipCode, "not found")

	return AddressData{}

}
