package handlers

import (
	"github.com/julienschmidt/httprouter"
	jwt "github.com/dgrijalva/jwt-go"
	"Backend_task_RF/verification"
	"Backend_task_RF/validation"
	"Backend_task_RF/hashing"
	"Backend_task_RF/data"
	"Backend_task_RF/DB"
	"encoding/json"
	"net/http"
	"os"
)

type errMessage struct {
	Message string `json:"message"`
}

func AddNewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	addedUser := data.RegisterUser{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&addedUser)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	db, err := DB.Open()
	if err != nil {
		ResponseError(w, 500, err)
		return
	}
	resultNameUser, err := validation.ValidateNameUser(addedUser.Name)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultEmailUser, err := validation.ValidateEmailUser(addedUser.Email, db)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultPasswordUser, err := validation.ValidatePasswordUsers(addedUser.Password)
	hashPasswordUser, _ := hashing.HashPasswordUser(resultPasswordUser)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultPhoneNumber, err := validation.ValidatePhoneNumber(addedUser.PhoneNumber)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	err = DB.AddNewUser(resultNameUser, resultEmailUser, hashPasswordUser, resultPhoneNumber)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth := data.Auth{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}
	
	db := DB.Connect()
	userid, err := verification.VerificationLogin(auth.Username, auth.Password, db)
    if err != nil {
		ResponseError(w, 400, err)
		return
	}
	
	tk := &data.Token{UserId: userid}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	w.Header().Set("Authorization", "Bearer " + tokenString)
	w.WriteHeader(200)
}

func GetListUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var attribute, order, offset, limit string
	url := r.URL.Query()

	if len(r.URL.RawQuery) > 0 {
		attribute = url.Get("attribute")
		order = url.Get("order")
		offset = url.Get("offset")
		limit = url.Get("limit")

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resultAttribute, err := validation.ValidateAttribute(attribute)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}
	resultOrder, err := validation.ValidateOrder(order)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultOffset, err := validation.ValidateOffset(offset)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}
	
	resultLimit, err := validation.ValidateLimit(limit)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}
	resultListUsers, err := DB.ListUsers(resultAttribute, resultOrder, resultOffset, resultLimit)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}
	if err = json.NewEncoder(w).Encode(resultListUsers); err != nil {
		ResponseError(w, 500, err)
	}
}

func UpdateUserPhoneNumber(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := r.Context().Value("user").(int)
	userUpdate := data.RegisterUser{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}
	
	resultPhoneNumber, err := validation.ValidatePhoneNumber(userUpdate.PhoneNumber)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	err = DB.UpdatePhoneNumber(resultPhoneNumber, id)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}

}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
