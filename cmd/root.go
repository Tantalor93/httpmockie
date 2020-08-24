package cmd

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go-http-mock/internal/log"
	"go-http-mock/internal/mockspec"
	"net/http"
	"strconv"
)

var router = mux.NewRouter()

var port *int
var dir *string

var RootCmd = cobra.Command{
	Use: "go-http-mock",
	Run: func(cmd *cobra.Command, args []string) {
		defer log.Logger.Sync()

		specs := mockspec.CollectFromDirectory(*dir)

		for _, spec := range specs {
			log.Logger.Infof("Registering handler for path '%s'", spec.Path)
			createHandler(spec)
		}

		itoa := strconv.Itoa(*port)
		if err := http.ListenAndServe(":"+itoa, router); err != nil {
			panic(err)
		}
	},
}

func init() {
	port = RootCmd.PersistentFlags().Int("port", 8081, "port to run mock server on")
	dir = RootCmd.PersistentFlags().String("dir", ".", "directory with mock specification")
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

func Execute() {
	RootCmd.Execute()
}