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
)


// CreatePublication creates a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var publication models.Publication
	if error = json.Unmarshal(requestBody, &publication); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	userId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	publication.AuthorId = userId

	if error = publication.Prepare(); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewPublicationRepository(db)
	publication.ID, error = repository.CreatePublication(publication)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusCreated, publication)
}

// UpdatePublication updates a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

	userId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

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

	repository := persistence.NewPublicationRepository(db)
	storedPublication, error := repository.GetPublicationById(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	if storedPublication.AuthorId != userId {
		responses.ErrorResponse(w, http.StatusUnauthorized,
			errors.New("A user can edit its own publications, only"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)

	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var publication models.Publication
	if error = json.Unmarshal(requestBody, &publication); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	publication.AuthorId = userId

	if error = publication.Prepare(); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	error = repository.UpdatePublication(id, publication)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusNoContent, nil)
}

// DeletePublication deletes a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {

	userId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

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

	repository := persistence.NewPublicationRepository(db)
	storedPublication, error := repository.GetPublicationById(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	if storedPublication.AuthorId != userId {
		responses.ErrorResponse(w, http.StatusUnauthorized,
			errors.New("A user can delete its own publications, only"))
		return
	}

	error = repository.DeletePublication(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusNoContent, nil)
}

// GetPublication get a publication by id
func GetPublication(w http.ResponseWriter, r *http.Request) {

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

	repository := persistence.NewPublicationRepository(db)
	publication, error := repository.GetPublicationById(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	if publication.ID != 0 {
		responses.JsonResponse(w, http.StatusOK, publication)
	} else {
		responses.JsonResponse(w, http.StatusNotFound, nil)
	}
}

// GetPublications get all publications of a user and its followers
func GetPublications(w http.ResponseWriter, r *http.Request) {

	userId, error := security.ExtractUserId(r)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := persistence.NewPublicationRepository(db)
	publications, error := repository.GetPublicationsForUserId(userId)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, publications)
}

// GetUserPublications get publications of a user
func GetUserPublications(w http.ResponseWriter, r *http.Request) {

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

	repository := persistence.NewPublicationRepository(db)
	publications, error := repository.GetUserPublicationById(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, publications)
}

// LikePublication registers a like in publication
func LikePublication(w http.ResponseWriter, r *http.Request) {

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

	repository := persistence.NewPublicationRepository(db)
	error = repository.RegisterPublicationLike(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, nil)
}

// LikePublication registers an unlike in publication
func UnlikePublication(w http.ResponseWriter, r *http.Request) {

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

	repository := persistence.NewPublicationRepository(db)
	error = repository.RegisterPublicationUnlike(id)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}
	responses.JsonResponse(w, http.StatusOK, nil)
}