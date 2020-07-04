package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"go-http-mock/internal/log"
	"go-http-mock/internal/mockspec"

	"net/http"
	"strconv"
)

var router = mux.NewRouter()

func main() {
	defer log.Logger.Sync()

	port, dir := parseFlags()

	specs := mockspec.CollectFromDirectory(*dir)

	for _, spec := range specs {
		log.Logger.Infof("Registering handler for path '%s'", spec.Path)
		createHandler(spec)
	}

	itoa := strconv.Itoa(*port)
	if err := http.ListenAndServe(":"+itoa, router); err != nil {
		panic(err)
	}
}

func parseFlags() (*int, *string) {
	port := flag.Int("port", 8081, "port")
	dir := flag.String("dir", ".", "directory with mock specification")

	flag.Parse()

	if port == nil || dir == nil {
		log.Logger.Errorf("--port and --dir flags have to be specified")
		panic("Required arguments not provided")
	}

	return port, dir
}

func createHandler(mock mockspec.MockSpecification) *mux.Route {
	return router.HandleFunc(mock.Path, func(writer http.ResponseWriter, request *http.Request) {
		if mock.Status != 0 {
			writer.WriteHeader(mock.Status)
		}
		if mock.JsonBody != nil {
			json.NewEncoder(writer).Encode(mock.JsonBody)
		}
		if mock.Body != nil {
			writer.Write([]byte(*mock.Body))
		}
	})
}
