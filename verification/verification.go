package verification

import (
	"Backend_task_RF/validation"
	"Backend_task_RF/hashing"
	"Backend_task_RF/data"
	"database/sql"
	"errors"
)

func VerificationLogin(loginUser, passUser string, db *sql.DB) (int, error) {

	prepareLoginUser := validation.PrepareEmailUser(loginUser)
	preparePassUser := validation.PreparePassword(passUser)

	if prepareLoginUser == "" {
		return 0, errors.New("Введите логин")
	} else if preparePassUser == "" {
		return 0, errors.New("Введите пароль")
	}
	
	row := db.QueryRow("SELECT id, email, password  FROM the_users WHERE email = $1", prepareLoginUser)
	var dataUser data.RegisterUser
	err := row.Scan(&dataUser.Id, &dataUser.Email, &dataUser.Password)
	if err != nil {
		return 0, errors.New("Введен неверный логин или пароль")
	}

	hashPass, err := hashing.HashPasswordUser(preparePassUser)
	if err != nil {
		return 0, err
	}

	idUser := dataUser.Id
	if dataUser.Password != hashPass {
		return 0, errors.New("Введен неверный логин или пароль")
	}
	return idUser, nil
}
