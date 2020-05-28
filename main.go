package main

import (
	"elastic-manager/manager"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var ElasticUrl string = "http://127.0.0.1:9200/"

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/v1/getIndices", getIndicesHandler).Methods("GET")
	mux.HandleFunc("/api/v1/createIndex", createIndexHandler).Methods("GET")
	mux.HandleFunc("/api/v1/deleteIndex", de)

	const address string = ":8080"

	server := http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func getIndicesHandler(w http.ResponseWriter, r *http.Request) {
	manager := &manager.ElasticManager{Url: ElasticUrl}
	indices := manager.GetIndices(&w)
	json.NewEncoder(w).Encode(indices)
}

func createIndexHandler(w http.ResponseWriter, r *http.Request) {
	manager := &manager.ElasticManager{Url: ElasticUrl}
	result := manager.CreateIndex(&w)
	json.NewEncoder(w).Encode(result)
}

func deleteIndexHandler(w http.ResponseWriter, r *http.Request) {
	manager := &manager.ElasticManager{Url: ElasticUrl}
	result := manager.CreateIndex(&w)
	json.NewEncoder(w).Encode(result)
}
