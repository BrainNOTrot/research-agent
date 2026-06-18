package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc(
		"/research/start",
		StartResearchHandler,
	)

	http.HandleFunc(
		"/research/status",
		StatusHandler,
	)

	fmt.Println(
		"Server running on :8080",
	)

	http.HandleFunc(
		"/research/stop",
		StopResearchHandler,
	)

	fs := http.FileServer(
		http.Dir("dashboard/dist"),
	)

	http.Handle("/", fs)

	http.ListenAndServe(
		":8080",
		nil,
	)

}
