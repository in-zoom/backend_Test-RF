package verification

import (
	"Backend_task_RF/validation"
	"Backend_task_RF/hashing"
	"Backend_task_RF/data"
	_ "github.com/lib/pq"
	 "Backend_task_RF/DB"
	"errors"
	"regexp"
)

func VerificationLogin(loginUser, passUser string) (int, error) {

	db, err := DB.Open()
	if err != nil {
		return 0, err
	}

	if passUser == "" {
		return 0, errors.New("Введите логин")
	}

	prepareLoginUser := validation.PrepareEmailUser(loginUser)
	hashPass, err := hashing.HashPasswordUser(passUser)
	if err != nil {
		return 0, err
	}

	if prepareLoginUser == "" {
		return 0, errors.New("Введите логин")
	}

	pattern := `^\w+@\w+\.\w+$`
	matched, err := regexp.Match(pattern, []byte(prepareLoginUser))
	if matched == false || err != nil {
		return 0, errors.New("Введен неверный логин или пароль")
	}

	row := db.QueryRow("SELECT id, email, password  FROM the_users WHERE email = $1", loginUser)
	var dataUser data.User
	err = row.Scan(&dataUser.Id, &dataUser.Email, &dataUser.Password)
	if err != nil {
		return 0, err
	}
	idUser := dataUser.Id
	if len(dataUser.Email) <= 0 && dataUser.Password != hashPass {
		return 0, errors.New("Введен неверный логин или пароль")
	}
	return idUser, nil
}
