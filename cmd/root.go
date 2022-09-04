package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/tantalor93/httpmockie/internal/mockspec"
)

var router = mux.NewRouter()

const (
	portArgumentName      = "port"
	portArgumentShortName = "p"

	dirArgumentName      = "dir"
	dirArgumentShortName = "d"
)

var (
	port *int
	dir  *string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RootCmd root of command.
var RootCmd = cobra.Command{
	Use: "httpmockie",
	RunE: func(cmd *cobra.Command, args []string) error {
		specs, err := mockspec.CollectFromDirectory(*dir)
		if err != nil {
			return err
		}

		for _, spec := range specs {
			fmt.Printf("Registering handler for path '%s'\n", spec.Path)
			createHandler(spec)
		}

		addr := ":" + strconv.Itoa(*port)
		fmt.Printf("Starting server on address %s\n", addr)
		if err := http.ListenAndServe(addr, router); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	port = RootCmd.PersistentFlags().IntP(portArgumentName, portArgumentShortName, 8081, "port to run mock server on")
	dir = RootCmd.PersistentFlags().StringP(dirArgumentName, dirArgumentShortName, ".", "directory with mock specifications")
	if err := RootCmd.MarkPersistentFlagRequired(dirArgumentName); err != nil {
		panic(err)
	}
}

func createHandler(mock mockspec.MockSpecification) *mux.Route {
	return router.HandleFunc(mock.Path, func(writer http.ResponseWriter, request *http.Request) {
		if mock.Delay != nil {
			dur := mock.Delay.DurationMs
			deviation := 0
			if mock.Delay.DeviationMs != 0 {
				dur -= mock.Delay.DeviationMs
				deviation = rand.Intn(mock.Delay.DeviationMs + mock.Delay.DeviationMs)
			}
			time.Sleep(time.Duration(dur+deviation) * time.Millisecond)
		}
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
		if mock.JSONBody != nil {
			json.NewEncoder(writer).Encode(mock.JSONBody)
		}
		if mock.Base64Body != nil {
			writer.Write(mock.Base64Body)
		}
		if mock.Body != nil {
			writer.Write([]byte(*mock.Body))
		}
	})
}

// Execute executes root command.
func Execute() error {
	return RootCmd.Execute()
}
