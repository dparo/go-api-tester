package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	port := 42069
	fmt.Printf("Server listening on port %d\n", port)

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	delayParam := r.URL.Query().Get("delay")
	responseStatusParam := r.URL.Query().Get("status")

	delay, err := strconv.Atoi(delayParam)
	if err != nil || delay < 0 {
		delay = 0
	}

	responseStatus, err := strconv.Atoi(responseStatusParam)
	if err != nil || responseStatus < 0 {
		responseStatus = 200
	}

	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}

	type Message struct {
		Status int    `json:"status"`
		Text   string `json:"message"`
	}

	message := Message{Status: responseStatus, Text: "Hello, World!"}
	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(jsonData)
}
