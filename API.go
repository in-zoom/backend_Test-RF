package main

import (
	"Backend_task_RF/handlers"
	"net/http"
    "github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/registration", handlers.AddNewUser)
	http.ListenAndServe(":8080", router)
}
