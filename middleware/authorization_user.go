package middleware

import (
	"Backend_task_RF/handlers"
	"Backend_task_RF/validation"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UsersMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		login, password, hasAuth := r.BasicAuth()
		errr := validation.AuthorizationLogin(login)
		err := validation.AuthorizationPass(password)

		if hasAuth && errr == nil && err == nil {
			next(w, r, ps)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func UserHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.Protected(w, r, ps)
		//GetListUsers(w, r, ps)
	}
}
