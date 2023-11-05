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
	http.HandleFunc("/status", serverStatus)
	log.Println("Application has Started ....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func serverStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("/status endpoint was called")

		msg := map[string]string{"Message": "A szerver működik"}
		status, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(status)
	}
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
		log.Println("New message received to the / endpoint ")
	} else if r.Method == "GET" {
		response, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		log.Println("Message was queried from the / endpoint ")

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func generateID() string {
	time.Sleep(1)
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
