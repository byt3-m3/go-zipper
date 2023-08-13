package models

import (
	"testing"
)

func TestZipFetcher(t *testing.T) {

	fetcher := NewZipFetcher(WithZipFileLocation("uszips.csv"))

	fetcher.GetAddressData("34761")
}
