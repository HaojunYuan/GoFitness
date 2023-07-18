package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/workouts", CreateWorkout).Methods("POST")
	router.HandleFunc("/workouts", GetWorkouts).Methods("GET")
	router.HandleFunc("/workouts/{id}", GetWorkout).Methods("GET")
	router.HandleFunc("/workouts/{id}", UpdateWorkout).Methods("PUT")
	router.HandleFunc("/workouts/{id}", DeleteWorkout).Methods("DELETE")

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
