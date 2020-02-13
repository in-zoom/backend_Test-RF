package DB

import (
	"Backend_task_RF/data"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB
var err error
var rows *sql.Rows

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
	ins := "CREATE TABLE IF NOT EXISTS the_users (id SERIAL, user_name VARCHAR, email VARCHAR, password VARCHAR)"
	_, err = db.Exec(ins)
	if err != nil {
		return err
	}
	counterSequenceId()
	return nil

}

func counterSequenceId() error {
	query := "Select last_value from the_users_id_seq"
	rows, err = db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var valueId int
	for rows.Next() {
		if err = rows.Scan(&valueId); err != nil {
			return err
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}
	if valueId == 1 {
		ins := "ALTER SEQUENCE the_users_id_seq RESTART WITH 123"
		_, err = db.Exec(ins)
		if err != nil {
			return err
		}
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

func ListUsers(attribute, order, offset, limit string) ([]data.User, error) {
	db = connect()
	query := "SELECT id, user_name, email FROM the_users" + " " + attribute + " " + order + " " + offset + " " + limit
	rows, err = db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]data.User, 0)
	var user data.User
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		list = append(list, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
