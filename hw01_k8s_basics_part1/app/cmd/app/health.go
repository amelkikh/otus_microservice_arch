package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

var j, _ = json.Marshal(StatusResponse{
	Status: "OK",
})

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if _, err := w.Write(j); err != nil {
		log.Panic(err)
	}
}
