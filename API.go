package main

import (
	"github.com/julienschmidt/httprouter"
	"Backend_task_RF/handlers"
	"github.com/joho/godotenv"
	"net/http"
	"log"
	"os"
)

func init(){
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
	}
}

func main() {
	router := httprouter.New()
	router.POST("/index/registration", handlers.AddNewUser)
	router.GET("/users", handlers.GetListUsers)
	port, _ := os.LookupEnv("PORT")
	log.Fatal(http.ListenAndServe(port, router))
}
