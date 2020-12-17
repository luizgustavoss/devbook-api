package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/persistence"
	"devbook/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// CreateUser creates a user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	if error = user.PrepareCreate(); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusCreated, user)
}

// ListUsers lists all users
func ListUsers(w http.ResponseWriter, r *http.Request) {

	description := strings.ToLower(r.URL.Query().Get("desc"))

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	users, error := repository.ListUsers(description)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusOK, users)
}

// FindUserById find a user by id
func FindUserById(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	id, error := strconv.ParseUint(pathParameters["id"], 10, 64)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	user, error := repository.GetUserById(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	if user.ID != 0 {
		responses.JsonResponse(w, http.StatusOK, user)
	} else {
		responses.JsonResponse(w, http.StatusNotFound, nil)
	}

}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	id, error := strconv.ParseUint(pathParameters["id"], 10, 64)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	if error = user.PrepareUpdate(); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	error = repository.Update(id, user)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusNoContent, nil)


}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	id, error := strconv.ParseUint(pathParameters["id"], 10, 64)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	error = repository.Delete(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	responses.JsonResponse(w, http.StatusNoContent, nil)
}