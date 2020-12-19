package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/persistence"
	"devbook/src/responses"
	"devbook/src/security"
	"encoding/json"
	"errors"
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

	tokenUserId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	if id != tokenUserId {
		responses.ErrorResponse(w, http.StatusForbidden,
			errors.New("Cannot change other user's data"))
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

	tokenUserId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	if id != tokenUserId {
		responses.ErrorResponse(w, http.StatusForbidden,
			errors.New("Cannot change other user's data"))
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


// FollowUser follows a user
func FollowUser(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	followedId, error := strconv.ParseUint(pathParameters["id"], 10, 64)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	followerId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	if followedId == followerId {
		responses.ErrorResponse(w, http.StatusBadRequest,
			errors.New("User cannot follow itself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	error = repository.FollowUser(followedId, followerId)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, nil)
}


// UnfollowUser unfollows a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	followedId, error := strconv.ParseUint(pathParameters["id"], 10, 64)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	followerId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	if followedId == followerId {
		responses.ErrorResponse(w, http.StatusBadRequest,
			errors.New("User cannot unfollow itself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewUserRepository(db)
	error = repository.UnfollowUser(followedId, followerId)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, nil)
}

// GetUserFollowers lists a user's followers
func GetUserFollowers(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	followedId, error := strconv.ParseUint(pathParameters["id"], 10, 64)
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
	users, error := repository.GetFollowersForUserId(followedId)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, users)
}

// GetFollowedUsers lists users a user follows
func GetFollowedUsers(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	followerId, error := strconv.ParseUint(pathParameters["id"], 10, 64)
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
	users, error := repository.GetFollowedUsersForUserId(followerId)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, users)
}

// UpdatePassword updates a user's password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	loggedUserId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	pathParameters := mux.Vars(r)
	userId, error := strconv.ParseUint(pathParameters["id"], 10, 64)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	if loggedUserId != userId {
		responses.ErrorResponse(w, http.StatusUnauthorized,
			errors.New("A user can update its own password, only"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var passwordUpdate models.PasswordUpdate
	if error = json.Unmarshal(requestBody, &passwordUpdate); error != nil {
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
	currentPassword, error := repository.GetUserPasswordById(userId)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	if error = security.CheckPassword(
		currentPassword, passwordUpdate.PreviousPassword); error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized,
			errors.New("Invalid previous password"))
		return
	}

	hashedPassword, error := security.Hash(passwordUpdate.NewPassword)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest,
			errors.New("Failed to update password"))
		return
	}

	error = repository.UpdateUserPassword(
		userId, string(hashedPassword))
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, nil)
}