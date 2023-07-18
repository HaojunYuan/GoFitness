package main

type Workout struct {
	ID             int    `json:"id"`
	Date           string `json:"date"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Duration       int    `json:"duration"`
	CaloriesBurned int    `json:"calories_burned"`
}

type WorkoutList []Workout
