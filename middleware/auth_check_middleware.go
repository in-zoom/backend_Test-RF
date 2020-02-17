package middleware

import (
	"github.com/julienschmidt/httprouter"
	jwt "github.com/dgrijalva/jwt-go"
	"Backend_task_RF/handlers"
	"Backend_task_RF/data"
	"net/http"
    "strings"
	"context"
	"errors"
	"os"
)

func AuthCheckMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		
		w.Header().Add("Content-Type", "application/json")
		tokenHeader := r.Header.Get("Authorization") 

		if tokenHeader == "" { 
			err := errors.New("Отсутствует токен авторизации")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		splitted := strings.Split(tokenHeader, " ") 
		if len(splitted) != 2 {
			err := errors.New("Неверный токен авторизации")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		tokenPart := splitted[1] 
		tk := &data.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { 
			err := errors.New("Неверный токен аутентификации")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		if !token.Valid { 
			err := errors.New("Токен недействителен")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next(w, r, ps) 
	}
}

func UserHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.Login(w, r, ps)
	}
}

func UsersListHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.GetListUsers(w, r, ps)
	}
}

func UpdateUserHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.UpdateUserPhoneNumber(w, r, ps)
	}
}
