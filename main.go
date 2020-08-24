package main

import (
	"github.com/gorilla/mux"
	"go-http-mock/cmd"
)

var router = mux.NewRouter()

func main() {
	cmd.Execute()
}
