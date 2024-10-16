package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"event-router/src/app/handlers/event"

	"github.com/google/uuid"
)

func TestHandleEventPOST(t *testing.T) {
	// Создаем тестовый запрос с данными события
	eventData := event.EventRequest{Context: map[string]interface{}{"key": "value"}}
	jsonData, err := json.Marshal(eventData)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем запрос с телом в виде JSON-данных
	req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(event.HandleEvent)

	// Обрабатываем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем код статуса
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус %v, получен: %v", http.StatusOK, status)
	}
}

func TestHandleEventGET(t *testing.T) {
	// Создаем тестовый запрос с данными события и сохраняем его в хранилище
	eventData := event.EventRequest{ID: uuid.New(), Context: map[string]interface{}{"key": "value"}}
	jsonData, err := json.Marshal(eventData)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(event.HandleEvent)
	handler.ServeHTTP(rr, req)

	// Создаем запрос на получение события
	reqGet, err := http.NewRequest("GET", "/event/"+eventData.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	rrGet := httptest.NewRecorder()
	handler.ServeHTTP(rrGet, reqGet)

	// Проверяем код статуса
	if status := rrGet.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус %v, получен: %v", http.StatusOK, status)
	}

	// Декодируем ответ и проверяем содержимое
	var respEvent event.EventRequest
	err = json.NewDecoder(rrGet.Body).Decode(&respEvent)
	if err != nil {
		t.Fatal(err)
	}
	if respEvent.ID != eventData.ID || len(respEvent.Context) == 0 {
		t.Errorf("Ожидается событие с ID %v и контекстом %v, получено: %v", eventData.ID, eventData.Context, respEvent)
	}
}
