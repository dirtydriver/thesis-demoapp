package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	clearMessages()

	t.Run("Test POST", func(t *testing.T) {
		message := Message{
			Text:     "Hello, world!",
			User:     "Akos",
			DateTime: time.Now(),
		}

		payload, _ := json.Marshal(message)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()

		handler(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
		}

		expectedResponse := "Message received and stored.\n"
		if rr.Body.String() != expectedResponse {
			t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
		}
	})

	t.Run("Test GET", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		handler(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
		}

		var response []Message
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}

		if len(response) != 1 {
			t.Errorf("Handler returned unexpected number of messages: got %v, want %v", len(response), 1)
		}

		expectedMessage := "Hello, world!"
		if response[0].Text != expectedMessage {
			t.Errorf("Handler returned unexpected message text: got %v, want %v", response[0].Text, expectedMessage)
		}
	})
}

func clearMessages() {
	messages = []Message{}
}

func TestGenerateID(t *testing.T) {
	id1 := generateID()
	id2 := generateID()

	if id1 == id2 {
		t.Errorf("Generated IDs are not unique: %v, %v", id1, id2)
	}
}