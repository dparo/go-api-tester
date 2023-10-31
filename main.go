package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	port := 0
	flag.IntVar(&port, "p", 42069, "Provide a port number")
	flag.Parse()

	fmt.Printf("Server listening on port %d\n", port)

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	delayParam := r.URL.Query().Get("delay")
	statusParam := r.URL.Query().Get("status")

	delay, err := strconv.Atoi(delayParam)
	if err != nil || delay < 0 {
		delay = 0
	}

	status, err := strconv.Atoi(statusParam)
	if err != nil || status < 0 {
		status = 200
	}

	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	type Message struct {
		Status  int                 `json:"status"`
		Text    string              `json:"message"`
		Headers map[string][]string `json:"headers"`
		Body    string              `json:"body"`
	}

	body, err := json.Marshal(
		Message{Status: status, Text: "Hello, World!", Headers: r.Header, Body: string(bodyBytes)},
	)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}
