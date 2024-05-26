package main

import (
	"fmt"
	"log"
	"net/http"

	resource "news-scrapper/resource"

	"github.com/gorilla/mux"
)

func main() {
	InitializeServer()
}

func InitializeServer() {
	fmt.Println("Iniciando servidor.")
	router := generateRoutes()
	fmt.Println("Rotas geradas")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func generateRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/infomoney", resource.Search).Methods("GET")
	return router
}
