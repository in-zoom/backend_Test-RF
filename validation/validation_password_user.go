package validation

import (
	"github.com/joho/godotenv"
	"encoding/hex"
	"crypto/md5"
	"strings"
	"errors"
	"os"
)

func ValidatePasswordUsers(addPassword string) (resultPasswordUser string, err error) {
	addUserPassword := preparePassword(addPassword)
	if addUserPassword == "" {
		return "", errors.New("Не задан пароль")
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

func preparePassword(imputPassword string) (outputPassword string) {
	passwordSpaceRemoval := strings.TrimSpace(imputPassword)
	return passwordSpaceRemoval
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
