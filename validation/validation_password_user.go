package validation

import (
	"Backend_task_RF/DB"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func ValidatePasswordUsers(addPassword string) (resultPasswordUser string, err error) {
	addUserPassword, err := preparePassword(addPassword)
	if err != nil {
		return "", err
	}
	if len(addUserPassword) < 5 {
		return "", errors.New("Пароль должен состоять хотя бы из 6 символов")
	}
	hashPassword, err := hashPasswordUser(addUserPassword)
	if err != nil {
		return "", err
	}
	return hashPassword, nil
}

func AuthorizationPass(passwordEntered string) (err error) {
	db, err := DB.Open()
	if err != nil {
		return err
	}
	passwordPrepare, err := preparePassword(passwordEntered)
	if err != nil {
		return err
	}
	hashPassword, err := hashPasswordUser(passwordPrepare)
	if err != nil {
		return err
	}

	query := "SELECT password FROM the_users WHERE password = " + " " + "'" + hashPassword + "'"
	rows, err := db.Query(query)

	if err != nil {
		return err
	}
	defer rows.Close()
	var userPassFromExisting string
	for rows.Next() {
		if err = rows.Scan(&userPassFromExisting); err != nil {
			return err
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}
	if userPassFromExisting == "" {
		return errors.New("Введен неверный логин или пароль")
	}
	return nil
}

func preparePassword(imputPassword string) (outputPassword string, err error) {
	if imputPassword == "" {
		return "", errors.New("Не задан пароль")
	}
	passwordSpaceRemoval := strings.TrimSpace(imputPassword)
	return passwordSpaceRemoval, nil
}

func hashPasswordUser(password string) (hashPassword string, err error) {
	err = godotenv.Load(".env") //Загрузить файл .env
	if err != nil {
		return "", err
	}
	salt := md5.Sum([]byte(os.Getenv("salt")))
	saltHash := hex.EncodeToString(salt[:])
	hashPasswordSalt := password + saltHash

	hash := md5.Sum([]byte(hashPasswordSalt))
	return hex.EncodeToString(hash[:]), nil
}
