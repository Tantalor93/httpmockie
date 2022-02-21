package main

import (
	"go-http-mock/cmd"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	cmd.Execute()
}
