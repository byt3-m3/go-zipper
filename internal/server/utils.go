package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func buildJSONResponse(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("content-type", "application/json")
	return w
}

func marshallInterface(data interface{}) []byte {
	respBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return respBytes
}

func BuildJSONResponseStatusOkWBody(w http.ResponseWriter, body interface{}) http.ResponseWriter {
	w = buildJSONResponse(w)
	w.WriteHeader(http.StatusOK)
	respBytes := marshallInterface(body)
	_, _ = w.Write(respBytes)

	return w
}
