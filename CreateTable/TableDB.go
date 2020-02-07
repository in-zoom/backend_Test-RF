package table

import (
	"Backend_task_RF/login"
	_ "github.com/lib/pq"
)

var err error

func CreateTable() (err error) {
	db := login.Init()
	ins := "CREATE TABLE IF NOT EXISTS the_users (id SERIAL, user_name varchar, e_mail varchar, password varchar)"
	_, err = db.Exec(ins)
	if err != nil {
		return err
	}
	return nil
}
