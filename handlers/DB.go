package handlers

import (
	"Backend_task_RF/login"
    _ "github.com/lib/pq"
)

var err error

func addNewUser(userName string) (err error) {
	db := login.Init()
	ins := "INSERT INTO cats (name) VALUES ($1)"
	_, err = db.Exec(ins, userName)
	if err != nil {
		return err
	}
	return nil
}
