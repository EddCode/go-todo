package main

import (
	"fmt"
	"net/http"

	"github.com/EddCode/go-todo/infraestructure"
)

func main() {
	router := infraestructure.Router()

	fmt.Println("Server listening on port 8080!")
	http.ListenAndServe(":8080", router)
}
