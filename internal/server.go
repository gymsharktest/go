package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/packs", getPacksHandler).Methods("GET")
	router.HandleFunc("/packs", addPackHandler).Methods("POST")
	router.HandleFunc("/order", orderHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":1080", router))
}
