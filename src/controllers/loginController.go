package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/persistence"
	"devbook/src/responses"
	"devbook/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Login authenticates a user
func Login(w http.ResponseWriter, r *http.Request) {

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, error)
		return
	}

	var credential models.Credential
	if error = json.Unmarshal(requestBody, &credential); error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	if error = credential.ValidateAndNormalizeCredential(); error != nil {
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
	user, error := repository.GetUserByEmail(credential.Email)
	if error != nil {
		responses.ErrorResponse(w, http.StatusBadRequest, error)
		return
	}

	if error = security.CheckPassword(user.Password, credential.Password); error != nil {
		responses.ErrorResponse(w, http.StatusUnauthorized, error)
		return
	}

	token, error := security.GetToken(user.ID, user.Email)
	if error != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, error)
		return
	}

	authToken := models.Token{
		UserId: strconv.FormatUint(user.ID, 10),
		Token: token}

	responses.JsonResponse(w, http.StatusOK, authToken)
}
