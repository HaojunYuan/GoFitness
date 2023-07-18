package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "workouts.db")
	if err != nil {
		return err
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS workouts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date TEXT,
			name TEXT,
			type TEXT,
			duration INTEGER,
			calories_burned INTEGER
		);`
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	log.Println("Connected to SQLite database")

	return nil
}

func AddWorkout(workout Workout) (int64, error) {
	result, err := db.Exec("INSERT INTO workouts (date, name, type, duration, calories_burned) VALUES (?, ?, ?, ?, ?)",
		workout.Date, workout.Name, workout.Type, workout.Duration, workout.CaloriesBurned)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetWorkoutList() (WorkoutList, error) {
	rows, err := db.Query("SELECT id, date, name, type, duration, calories_burned FROM workouts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := WorkoutList{}

	for rows.Next() {
		var workout Workout
		err := rows.Scan(&workout.ID, &workout.Date, &workout.Name, &workout.Type, &workout.Duration, &workout.CaloriesBurned)
		if err != nil {
			return nil, err
		}

		workouts = append(workouts, workout)
	}

	return workouts, nil
}

func GetWorkoutByID(id int) (Workout, error) {
	var workout Workout

	err := db.QueryRow("SELECT id, date, name, type, duration, calories_burned FROM workouts WHERE id = ?", id).
		Scan(&workout.ID, &workout.Date, &workout.Name, &workout.Type, &workout.Duration, &workout.CaloriesBurned)
	if err != nil {
		return workout, err
	}

	return workout, nil
}

func UpdateWorkoutByID(id int, updatedWorkout Workout) error {
	_, err := db.Exec("UPDATE workouts SET date = ?,name = ?, type = ?, duration = ?, calories_burned = ? WHERE id = ?",
		updatedWorkout.Date, updatedWorkout.Name, updatedWorkout.Type, updatedWorkout.Duration, updatedWorkout.CaloriesBurned, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteWorkoutByID(id int) error {
	_, err := db.Exec("DELETE FROM workouts WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
