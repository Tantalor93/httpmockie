package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"strconv"
)

type SomeResponse struct {
	Msg string `json:"msg"`
}

func main() {
	port := flag.Int("port", 8081, "port")

	flag.Parse()


	http.HandleFunc("/mock", func(writer http.ResponseWriter, request *http.Request) {
		response := SomeResponse{Msg: "Benky jede"}
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(response)
	})

	itoa := strconv.Itoa(*port)
	http.ListenAndServe(":"+itoa, nil)
}