# GoFitness

The GoFitness is a Go-based web application that allows users to track their workout sessions. Users can create new workout records, retrieve existing workout records, update workout details, and delete workout entries.

## Features

- Create a new workout record with details such as date, name, type, duration, and calories burned.
- Retrieve a list of all workout records.
- Retrieve a specific workout record by its ID.
- Update the details of an existing workout record.
- Delete a workout record.

## Prerequisites

Before running the project, ensure that you have the following installed:

- Go programming language (version 1.16 or above)
- SQLite database engine

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Go to the project directory:
    ```bash
    cd GoFitness
    ```
3. Initialize the SQLite database:
    ```bash
    go run db.go
    ```
    This will create a workouts.db file in the project directory, which will serve as the database for storing workout records.
4. Install dependencies:
    ```bash
    go mod download
    ```
5. Build the project:
    ```bash
    go build
    ```
6. Run the executable:
    ```bash
    ./GoFitness
    ```
    The server will start running on "http://localhost:8080".

## API Endpoints

- `POST /workouts`: Create a new workout record.
- `GET /workouts`: Retrieve a list of all workout records.
- `GET /workouts/{id}`: Retrieve a specific workout record by ID.
- `PUT /workouts/{id}`: Update the details of a workout record.
- `DELETE /workouts/{id}`: Delete a workout record by ID.

## Future Plans

- Implement persistent storage to retain workout records even after server restart.
- Build a frontend UI using React.
- Deploy the app.

These future plans aim to enhance functionality, improve user experience, and make the GoFitness application accessible to a wider audience.

