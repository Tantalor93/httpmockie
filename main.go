package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"go-http-mock/log"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type MockSpecification struct {
	JsonBody map[string]interface{} `json:"jsonBody"`
	Path     string                 `json:"path"`
	Status   int                    `json:"status"`
}

var router = mux.NewRouter()

func main() {
	defer log.Logger.Sync()

	port, dir := parseFlags()

	if err := filepath.Walk(*dir, walkDir); err != nil {
		panic(err)
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

func walkDir(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		log.Logger.Infof("Reading mock specification '%s'", path)

		file, e := os.Open(path)
		if e != nil {
			log.Logger.Errorf("Error opening while '%s'", path)
			return nil
		}
		read, err := ioutil.ReadAll(file)
		if err != nil {
			log.Logger.Errorf("Error while reading file '%s'", path)
			return nil
		}

		var mock MockSpecification
		e = json.Unmarshal(read, &mock)
		if e != nil {
			log.Logger.Errorf("Error while parsing mock specification in file '%s'", path)
			return nil
		}
		if len(mock.Path) > 0 {
			log.Logger.Infof("Registering handler for path '%s'", mock.Path)
			createHandler(mock)
		}
		return nil
	}
	return nil
}

func createHandler(mock MockSpecification) *mux.Route {
	return router.HandleFunc(mock.Path, func(writer http.ResponseWriter, request *http.Request) {
		if mock.Status != 0 {
			writer.WriteHeader(mock.Status)
		}
		if mock.JsonBody != nil {
			json.NewEncoder(writer).Encode(mock.JsonBody)
		}
	})
}
