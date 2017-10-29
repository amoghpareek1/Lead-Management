package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func postSalesforceConnectionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		log.Println(err)
		return
	}

	var connection Connection

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	err = json.Unmarshal(body, &connection)
	if err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	if connection.SalesforceAPIVersion == "" || connection.SalesforceName == "" || connection.SalesforcePassword == "" || connection.SalesforceSecurityToken == "" || connection.SalesforceUsername == "" {
		sendResponse(w, false, "All fields are required.")
		return
	}

	connection.OwnerID = session.Values["userID"].(uint)

	gormDatabase.Create(&connection)

	sendResponse(w, true, "Connection is created successfully.")
}

func postMySQLConnectionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		log.Println(err)
		return
	}

	var connection Connection

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	err = json.Unmarshal(body, &connection)
	if err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	if connection.MySQLDatabase == "" || connection.MySQLName == "" || connection.MySQLPassword == "" || connection.MySQLPort == "" || connection.MySQLServer == "" || connection.MySQLUserID == "" {
		sendResponse(w, false, "All fields are required.")
		return
	}

	connection.OwnerID = session.Values["userID"].(uint)

	gormDatabase.Create(&connection)

	sendResponse(w, true, "Connection is created successfully.")
}

func getConnectionsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		log.Println(err)
		return
	}

	var connections []Connection
	gormDatabase.Where("owner_id = ?", session.Values["userID"]).Find(&connections)

	sendResponse(w, true, connections)
}
