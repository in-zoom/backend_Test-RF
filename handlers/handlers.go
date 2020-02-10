package handlers

import (
	"github.com/julienschmidt/httprouter"
	"Backend_task_RF/validation"
	"Backend_task_RF/data"
	"Backend_task_RF/DB"
	"encoding/json"
	"net/http"
	"fmt"
)

type errMessage struct {
	Message string `json:"message"`
}

func GetListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var offset, limit string
	url := r.URL.Query()

	if len(r.URL.RawQuery) > 0 {
		offset = url.Get("offset")
		limit = url.Get("limit")

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resultOffset, err := validation.ValidateOffset(offset)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultLimit, err := validation.ValidateLimit(limit)
	if err != nil {
		responseError(w, 400, err)
		return
	}
	resultListUsers, err := DB.ListUsers(resultOffset, resultLimit)
	if err != nil {
		responseError(w, 500, err)
		return
	}
	if err = json.NewEncoder(w).Encode(resultListUsers); err != nil {
		responseError(w, 500, err)
	}
}

func AddNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//var addedUser User
	addedUser := data.User{}
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
