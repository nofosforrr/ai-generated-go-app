package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type EventRequest struct {
	ID      uuid.UUID   `json:"id"`
	Context interface{} `json:"context"`
}

var eventStore = make(map[uuid.UUID]EventRequest)

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePost(w, r)
	case "GET":
		handleGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var req EventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	}

	fmt.Printf("Received event with ID: %s and context: %v\n", req.ID, req.Context)
	eventStore[req.ID] = req
	w.WriteHeader(http.StatusOK)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/event/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	event, exists := eventStore[id]
	if !exists {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
