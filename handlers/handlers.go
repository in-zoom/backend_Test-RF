package handlers

import (
	"Backend_task_RF/login"
	"Backend_task_RF/validation"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type user struct {
	Name     string `json:"name"`
	e_mail   string `json:"e_mail"`
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
	resultNameUser, err := validation.ValidateNameUser(addedUser.Name, login.Init())
	if err != nil {
		responseError(w, 400, err)
		return
	}
	err = addNewUser(resultNameUser)
	if err != nil {
		responseError(w, 500, err)
		return
	}

	/*resultPasswordUser, err := validation.ValidatePasswordUser(addedUser.Password)
	if err != nil {
		responseError(w, 400, err)
		return
	}*/

}
func responseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
