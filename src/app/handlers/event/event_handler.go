package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type EventRequest struct {
	ID      uuid.UUID              `json:"event_id,omitempty"`
	Context map[string]interface{} `json:"context"`
}

var eventStore = make(map[uuid.UUID]EventRequest)

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req EventRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if req.ID == uuid.Nil {
			req.ID = uuid.New()
		}
		fmt.Printf("Received event with ID: %s and context: %v\n", req.ID, req.Context)
		eventStore[req.ID] = req
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == "GET" {
		// Extract the event_id from the URL parameters
		idStr := r.URL.Path[len("/event/"):]
		id, err := uuid.Parse(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		event, exists := eventStore[id]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Return the event as JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(event)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
