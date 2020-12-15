package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/persistence"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// creates a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	id, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Location", "/users/" + string(id))
	w.Write([]byte(fmt.Sprintf("Created a new user with id %d", id)))
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