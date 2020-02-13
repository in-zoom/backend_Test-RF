package validation

import (
	"Backend_task_RF/DB"
  _ "github.com/lib/pq"
	"database/sql"
	"errors"
	"regexp"
	"strings"
)

func ValidateEmailUser(EmailUser string, db *sql.DB) (resultEmailUser string, err error) {
	var rows *sql.Rows
	addEmailUser, err := prepareEmailUsre(EmailUser)
	if err != nil {
		return "", err
	}

	query := "SELECT email FROM the_users WHERE email = " + " " + "'" + addEmailUser + "'"
	rows, err = db.Query(query)

	if err != nil {
		return "", err
	}
	defer rows.Close()
	var userEmailFromExisting string
	for rows.Next() {
		if err = rows.Scan(&userEmailFromExisting); err != nil {
			return "", err
		}
	}
	if err = rows.Err(); err != nil {
		return "", err
	}
	if userEmailFromExisting == "" {
		return addEmailUser, nil
	} else {
		return "", errors.New("Пользователь с Email -" + " " + userEmailFromExisting + " " + "уже зарегистрирован")
	}

}

func AuthorizationLogin(email string) (err error) {
	db, err := DB.Open()
	if err != nil {
		return err
	}

	PrepareAuthorization, err := prepareEmailUsre(email)
	if err != nil {
		return err
	}

	query := "SELECT email FROM the_users WHERE email = " + " " + "'" + PrepareAuthorization + "'"
	rows, err := db.Query(query)

	if err != nil {
		return err
	}
	defer rows.Close()

	var userEmailFromExisting string
	for rows.Next() {
		if err = rows.Scan(&userEmailFromExisting); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if userEmailFromExisting == "" {
		return errors.New("Введен неверный логин или пароль")
	}
	return nil
}

func prepareEmailUsre(imputEmailUser string) (outputEmailUser string, err error) {

	emailSpaceRemoval := strings.TrimSpace(imputEmailUser)

	if emailSpaceRemoval == "" {
		return "", errors.New("Введите E-mail")
	}

	pattern := `^\w+@\w+\.\w+$`
	matched, err := regexp.Match(pattern, []byte(emailSpaceRemoval))
	if matched == false || err != nil {
		return "", errors.New("Email введен неверно")
	}
	return emailSpaceRemoval, nil
}