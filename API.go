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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	router := httprouter.New()
	router.POST("/index/registration", handlers.AddNewUser)
	router.GET("/index/login", middleware.UsersMiddleware(middleware.UserHandler()))
	router.GET("/index/users", middleware.UsersMiddleware(middleware.UserHandler()))
	port, _ := os.LookupEnv("PORT")
	log.Fatal(http.ListenAndServe(port, router))
}