package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func postJobHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionStore.Get(r, "user-session")
	if err != nil {
		log.Println(err)
		return
	}

	var job Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	gormDatabase.Create(&job)

	sendResponse(w, true, "Job created successfully.")
}

// func getJobsHandler(w http.ResponseWriter, r *http.Request) {
// 	session, err := sessionStore.Get(r, "user-session")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	var jobs []Job

// 	var currentUserJobs []Jobs

// 	gormDatabase.Find(&jobs)

// 	for i := 0; i < len(jobs); i++ {
// 		var connection SalesforceConnection
// 		gormDatabase.Where("id = ?", jobs[i].)
// 	}

// 	sendResponse(w, true, "Job created successfully.")
// }
