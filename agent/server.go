package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var currentAgent *Agent

type StartRequest struct {
	Topic string `json:"topic"`
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set(
		"Access-Control-Allow-Origin",
		"*",
	)
	w.Header().Set(
		"Access-Control-Allow-Headers",
		"Content-Type",
	)
	w.Header().Set(
		"Access-Control-Allow-Methods",
		"GET, POST, OPTIONS",
	)
}

func StartResearchHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	enableCors(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var req StartRequest

	err := json.NewDecoder(r.Body).
		Decode(&req)

	if err != nil {
		http.Error(
			w,
			"Bad request",
			http.StatusBadRequest,
		)
		return
	}

	currentAgent = NewAgent(req.Topic)

	go currentAgent.Run()

	w.WriteHeader(http.StatusOK)
}

func StatusHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	enableCors(w)

	if currentAgent == nil {
		json.NewEncoder(w).Encode(map[string]any{
			"count":        0,
			"currentTopic": "",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"count":        len(currentAgent.memory.Research),
		"currentTopic": currentAgent.current,
	})
}

func StopResearchHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	enableCors(w)

	if currentAgent == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("STOP REQUEST RECEIVED")

	currentAgent.Stop()

	w.WriteHeader(http.StatusOK)
}
