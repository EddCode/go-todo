package useCase

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/EddCode/go-todo/domain"
)

func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload := domain.GetAll()
	json.NewEncoder(w).Encode(payload)
}
