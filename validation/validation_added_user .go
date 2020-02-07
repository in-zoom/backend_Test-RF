package validation

import (
	"Backend_task_RF/CreateTable"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"regexp"
	"strings"
)

func ValidateNameUser(addedUserName string, db *sql.DB) (nameCat string, err error) {
	var rows *sql.Rows
	addUserName := prepareName(addedUserName)

	if addUserName == "" {
		return "", errors.New("Введите имя")
	}

	pattern := `^[A-Za-z]+$`
	matched, err := regexp.Match(pattern, []byte(addUserName))
	if matched == false || err != nil {
		return "", errors.New("Имя не может содержать цифры, знаки пунктуации, символы ")
	}

	query := "SELECT -- FROM -- WHERE -- = " + " " + "'" + addUserName + "'"
	rows, err = db.Query(query)

	if err != nil {
		return "", err
	}
	defer rows.Close()
	var userNameFromExisting string
	for rows.Next() {
		if err = rows.Scan(&userNameFromExisting); err != nil {
			return "", err
		}
	}
	if err = rows.Err(); err != nil {
		return "", err
	}
	if userNameFromExisting == "" {
		table.CreateTable()
		return addUserName, nil
	} else {
		return "", errors.New("Имя" + " " + userNameFromExisting + " " + "уже используется")
	}

}

func prepareName(imputUserName string) (outputUserName string) {
	nameSpaceRemoval := strings.TrimSpace(imputUserName)
	nameLowerCase := strings.ToLower(nameSpaceRemoval)
	formattedUserName := strings.Title(nameLowerCase)
	return formattedUserName
}
