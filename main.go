package main

import (
	"encoding/json"
	"net/http"
)

type SomeResponse struct {
	Msg string `json:"msg"`
}

func main() {
	response := SomeResponse{Msg: "Benky jede"}
	http.HandleFunc("/mock", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(response)
	})

	http.ListenAndServe(":8081", nil)
}