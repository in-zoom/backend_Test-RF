package validation

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func ValidatePasswordUsers(addPassword string) (resultPasswordUser string, err error) {
	addUserPassword := preparePassword(addPassword)
	if addUserPassword == "" {
		return "", errors.New("Не задан пароль")
	}
	/*pattern := `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?!.*\s).*$`
	matched, err := regexp.Match(pattern, []byte(addUserPassword))
	if matched == false || err != nil {
		return "", errors.New("Пароль должен соcтоять из цифр и букв ")
	}*/
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
