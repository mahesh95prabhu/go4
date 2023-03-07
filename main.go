package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Create the struct for user
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func main() {
	http.HandleFunc("/users", handleUsers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// this funtion routes the request to the correct function depending on request type
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	case "POST":
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
//This is the function for GET requests, which will serve the users details mentioned below
func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Name: "Alice", Age: 23},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	json.NewEncoder(w).Encode(users)
}
//This is the function for POST requests, which will create a user as per the provided details
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Do something with the user data, e.g. save to database
	w.WriteHeader(http.StatusCreated)
}
