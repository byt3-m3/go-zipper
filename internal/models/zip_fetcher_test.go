package models

import "testing"

func TestZipFetcher(t *testing.T) {
	fetcher := NewZipFetcher()

	fetcher.GetAddressData("34761")
}
