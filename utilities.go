package main

import (
	"encoding/json"
	"log"

	"net/http"
)

func contains(slice []string, needle string) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}

func sendResponse(w http.ResponseWriter, success bool, data interface{}) {
	var response Response
	response.Success = success
	response.Data = data

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}
