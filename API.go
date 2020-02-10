package main

import (
	"github.com/julienschmidt/httprouter"
	"Backend_task_RF/handlers"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/users", handlers.GetListUsers)
	router.POST("/registration", handlers.AddNewUser)
	http.ListenAndServe(":8080", router)
}
