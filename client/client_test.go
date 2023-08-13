package client

import (
	"fmt"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient(WithHost("https://zip-api.baxterhome.net"))

	data, err := client.GetAddressData("34761")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
