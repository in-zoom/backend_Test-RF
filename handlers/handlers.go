package handlers

import (
	"Backend_task_RF/OpenDB"
	"Backend_task_RF/validation"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type errMessage struct {
	Message string `json:"message"`
}

func AddNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var addedUser user
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&addedUser)
	if err != nil {
		responseError(w, 400, err)
		return
	}

	db, err := DB.Open()
	if err != nil {
		responseError(w, 500, err)
		return
	}
	resultNameUser, err := validation.ValidateNameUser(addedUser.Name)
	if err != nil {
		responseError(w, 400, err)
		return
	}

	resultEmailUser, err := validation.ValidateEmailUser(addedUser.Email, db)
	if err != nil {
		responseError(w, 400, err)
		return
	}

	resultPasswordUser, err := validation.ValidatePasswordUsers(addedUser.Password)
	fmt.Println(resultPasswordUser)
	if err != nil {
		responseError(w, 400, err)
		return
	}

	err = DB.AddNewUser(resultNameUser, resultEmailUser, resultPasswordUser)
	if err != nil {
		responseError(w, 500, err)
		return
	}

}
func responseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}