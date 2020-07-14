package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type SomeResponse struct {
	Msg string `json:"msg"`
}

func main() {
	response := SomeResponse{Msg: "Benky jede"}
	http.HandleFunc("/mock", func(writer http.ResponseWriter, request *http.Request) {
		rand := rand.Intn(50)
		time.Sleep(time.Duration(rand) * time.Millisecond)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(response)
	})

	http.ListenAndServe(":8081", nil)
}