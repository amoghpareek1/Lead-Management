package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")

	router.HandleFunc("/api/v1/sign-up", signUpHandler).Methods("POST")
	router.HandleFunc("/api/v1/sign-in", signInHandler).Methods("POST")
	router.HandleFunc("/api/v1/sign-out", signOutHandler).Methods("GET")

	router.HandleFunc("/api/v1/job", postJobHandler).Methods("POST")

	router.HandleFunc("/api/v1/salesforce-connections", postSalesforceConnectionHandler).Methods("POST")
	router.HandleFunc("/api/v1/mysql-connections", postMySQLConnectionHandler).Methods("POST")
	router.HandleFunc("/api/v1/connections", getConnectionsHandler).Methods("GET")

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))).Methods("GET")

	if err := http.ListenAndServe(":15000", router); err != nil {
		log.Println(err)
		return
	}
}
