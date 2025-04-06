package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("fuzzing code")

	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(fmt.Errorf("http.ListenAndServe: %v", err))
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)

		return
	}

	if req.Age < 0 {
		http.Error(w, "Invalid age", http.StatusBadRequest)

		return
	}

	// Create user logic goes here.

	_, err = fmt.Fprintf(w, "User %s created successfully", req.Username)
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
