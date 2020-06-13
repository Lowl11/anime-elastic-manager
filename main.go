package main

import (
	"encoding/json"
	"fmt"
	"lazy-owl/elastic-manager/src/elastic"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var ElasticUrl string = "http://127.0.0.1:9200/"

func main() {
	mux := mux.NewRouter()

	// Работа с индексами
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/api/v1/getIndices", getIndicesHandler).Methods("GET")
	mux.HandleFunc("/api/v1/createIndex", createIndexHandler).Methods("GET")
	mux.HandleFunc("/api/v1/deleteIndex", deleteIndexHandler).Methods("GET")

	// Работа с данными
	mux.HandleFunc("/api/v1/indexData", indexDataHandler).Methods("POST")

	const address string = ":8080"

	server := http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		"src/templates/index.html",
	)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	templ.Execute(w, nil)
}

func indexDataHandler(w http.ResponseWriter, r *http.Request) {
	result := elastic.IndexData(&w, ElasticUrl, r)
	json.NewEncoder(w).Encode(result)
}

func getIndicesHandler(w http.ResponseWriter, r *http.Request) {
	indices := elastic.GetIndices(&w, ElasticUrl)
	json.NewEncoder(w).Encode(indices)
}

func createIndexHandler(w http.ResponseWriter, r *http.Request) {
	result := elastic.CreateIndex(&w, ElasticUrl)
	json.NewEncoder(w).Encode(result)
}

func deleteIndexHandler(w http.ResponseWriter, r *http.Request) {
	result := elastic.DeleteIndex(&w, ElasticUrl)
	json.NewEncoder(w).Encode(result)
}
