package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	User     string    `json:"user"`
	DateTime time.Time `json:"datetime"`
}

var messages []Message

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		message.ID = generateID()
		message.DateTime = time.Now()
		messages = append(messages, message)
		fmt.Fprintln(w, "Message received and stored.")
	} else if r.Method == "GET" {
		response, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}