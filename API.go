package main

import (
	"github.com/julienschmidt/httprouter"
	"Backend_task_RF/middleware"
	"Backend_task_RF/handlers"
	"github.com/joho/godotenv"
	"net/http"
	"log"
	"os"
)

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	initialize()
	router := httprouter.New()
	router.POST("/registration", handlers.AddNewUser)
	router.POST("/login", handlers.Login)
	router.GET("/users", middleware.AuthCheckMiddleware(middleware.UsersListHandler()))
	router.PUT("/update", middleware.AuthCheckMiddleware(middleware.UpdateUserHandler()))
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe( ":" + port, router))
}