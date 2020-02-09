package DB

import (
	"database/sql"
	"fmt"
	"log"
	"os"
    "github.com/joho/godotenv"
)

var db *sql.DB
var err error

func Open() (*sql.DB, error) {
	db = connect()
	err = createTable()
	return db, err
}

func connect() *sql.DB {
	e := godotenv.Load(".env") //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func createTable() (err error) {
	ins := "CREATE TABLE IF NOT EXISTS the_users (id SERIAL, user_name VARCHAR, email VARCHAR, login VARCHAR, password VARCHAR)"
	_, err = db.Exec(ins)
	if err != nil {
		return err
	}
	return nil
}

func AddNewUser(userName, EmailUser, PasswordUser string) (err error) {
	ins := "INSERT INTO the_users (user_name, email, password) VALUES ($1, $2, $3)"
	_, err = db.Exec(ins, userName, EmailUser, PasswordUser)
	if err != nil {
		return err
	}
	return nil
}