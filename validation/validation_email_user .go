package validation

import (
  _ "github.com/lib/pq"
	"database/sql"
	"strings"
	"errors"
	"regexp"
)

func ValidateEmailUser(EmailUser string, db *sql.DB) (resultEmailUser string, err error) {
	var rows *sql.Rows
	addEmailUser := PrepareEmailUser(EmailUser)

	if addEmailUser == "" {
		return "", errors.New("Введите E-mail")
	}

	pattern := `^\w+@\w+\.\w+$`
	matched, err := regexp.Match(pattern, []byte(addEmailUser))
	if matched == false || err != nil {
		return "", errors.New("Email введен неверно")
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

func PrepareEmailUser(imputEmailUser string) (outputEmailUser string) {
	emailSpaceRemoval := strings.TrimSpace(imputEmailUser)
	return emailSpaceRemoval
}
