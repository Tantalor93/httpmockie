package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-http-mock/internal/mockspec"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var router = mux.NewRouter()

var (
	port *int
	dir  *string
)

var RootCmd = cobra.Command{
	Use: "go-http-mock",
	RunE: func(cmd *cobra.Command, args []string) error {
		specs, err := mockspec.CollectFromDirectory(*dir)
		if err != nil {
			return err
		}

		for _, spec := range specs {

			fmt.Printf("Registering handler for path '%s'\n", spec.Path)
			createHandler(spec)
		}

		itoa := strconv.Itoa(*port)
		if err := http.ListenAndServe(":"+itoa, router); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	port = RootCmd.PersistentFlags().Int("port", 8081, "port to run mock server on")
	dir = RootCmd.PersistentFlags().String("dir", ".", "directory with mock specification")
}

func createHandler(mock mockspec.MockSpecification) *mux.Route {
	return router.HandleFunc(mock.Path, func(writer http.ResponseWriter, request *http.Request) {
		if mock.Headers != nil {
			for k, v := range mock.Headers {
				for _, hv := range v {
					writer.Header().Add(k, hv)
				}
			}
		}
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
