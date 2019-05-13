package main

import (
    "log"
    "./routes"
    "net/http"
)

func main() {
    router := routes.SetRoutes()

    log.Fatal(http.ListenAndServe(":8080", router))
}