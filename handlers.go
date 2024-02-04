package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	openai "github.com/sashabaranov/go-openai"
)

func CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	id, err := AddWorkout(workout)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := GetWorkoutByID(int(id))
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get created workout", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdWorkout)
}

func GetWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := GetWorkoutList()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get workouts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workouts)
}

func GetWorkout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workoutID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	workout, err := GetWorkoutByID(workoutID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get workout", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(workout)
}

func UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workoutID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	var updatedWorkout Workout
	err = json.NewDecoder(r.Body).Decode(&updatedWorkout)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = UpdateWorkoutByID(workoutID, updatedWorkout)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update workout", http.StatusInternalServerError)
		return
	}

	// Fetch the updated workout from the database
	updatedWorkout, err = GetWorkoutByID(workoutID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch updated workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedWorkout)
}

func DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workoutID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	// Fetch the workout to be deleted from the database
	deletedWorkout, err := GetWorkoutByID(workoutID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch workout to be deleted", http.StatusInternalServerError)
		return
	}

	err = DeleteWorkoutByID(workoutID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedWorkout)
}

func ChatGPT(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Text string `json:"text"`
    }

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        log.Println(err)
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Load OpenAI API key from environment variable
    openaiAPIKey := os.Getenv("OPENAI_API_KEY")

    if openaiAPIKey == "" {
        log.Fatal("OpenAI API key not set")
    }

    client := openai.NewClient(openaiAPIKey)
    result, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input.Text,
				},
			},
		},
	)
    if err != nil {
        log.Println(err)
        http.Error(w, "Failed to generate response from ChatGPT", http.StatusInternalServerError)
        return
    }

    response := result.Choices[0].Message.Content

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"response": response})
}