package main

import (
	"fmt"
	"log"
	"net/http"

	source "scraper/resources"

	"github.com/gorilla/mux"
)

func InitializeServer() {
	fmt.Println("Iniciando servidor.")
	router := generateRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}

func generateRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/infomoney", source.Search).Methods("GET")
	return router
}
