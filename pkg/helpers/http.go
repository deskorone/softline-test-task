package helpers

import (
	"encoding/json"
	"net/http"
)

/*
	Вспомогательный пакет, с обертками для ответов на запросы
*/

type Response struct {
	Message string `json:"message"`
}

const (
	MethodNotAllowed = "метод недоступен"
)

func Failed(w http.ResponseWriter, text string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Message: text})
}

func Success(w http.ResponseWriter, b interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}
