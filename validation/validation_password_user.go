package validation

import (
	"github.com/joho/godotenv"
	"Backend_task_RF/DB"
	"encoding/hex"
	"crypto/md5"
	"strings"
	"errors"
	"os"
)

func ValidatePasswordUsers(addPassword string) (resultPasswordUser string, err error) {
	addUserPassword, err := preparePassword(addPassword)
	if err != nil {
		return "", err
	}
	if len(addUserPassword) <= 5 {
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
	passwordSpaceRemoval := strings.TrimSpace(imputPassword)

	if passwordSpaceRemoval == "" {
		return "", errors.New("Не задан пароль")
	}
	
	/*pattern := `[a-zA-Z0-9_\-\.]+\@[a-z0-9\.\-]+`
	matched, err := regexp.Match(pattern, []byte(passwordSpaceRemoval))
	if matched == false || err != nil {
		return "", errors.New("Email введен неверно")
	}*/
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