package controllers

import "net/http"

// creates a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Creating a user"))
}

// lists all users
func ListUsers(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Listing uses"))
}

// find a user by id
func FindUserById(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Finding user by id"))
}

// updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updating a user"))
}

// deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleting a user"))
}