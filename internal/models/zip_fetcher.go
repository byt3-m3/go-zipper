package models

import (
	"github.com/gocarina/gocsv"
	"go-zip/internal/vars"
	"log"
	"os"
)

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
}

func NewZipFetcher() ZipFetcher {
	f := &zipFetcher{}

	f.loadData()

	return f
}

func (f *zipFetcher) loadData() {

	file, err := os.Open(vars.ZipFileLocation)
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
