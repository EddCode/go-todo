package infraestructure

import (
	"github.com/EddCode/go-todo/useCase"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", useCase.Home).Methods("GET")
	router.HandleFunc("/api/task", useCase.GetAllTask).Methods("GET")

	return router
}
