package main

import (
	"./routes"
	"log"
	"net/http"
	//	"github.com/pkg/errors"
)

func main() {
	router := routes.SetRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
