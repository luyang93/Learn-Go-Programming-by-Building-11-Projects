package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/pkg/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
