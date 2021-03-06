package validation

import (
	"strings"
	"errors"
	"regexp"
)

func ValidateNameUser(UserName string) (resultNameUser string, err error) {
	addUserName := prepareName(UserName)

	if addUserName == "" {
		return "", errors.New("Введите имя")
	}

	pattern := `^[A-Za-z]+$`
	matched, err := regexp.Match(pattern, []byte(addUserName))
	if matched == false || err != nil {
		return "", errors.New("Имя не может содержать цифры, знаки пунктуации, символы, пробелы")
	}
	return addUserName, nil
}

func prepareName(imputUserName string) (outputUserName string) {
	nameSpaceRemoval := strings.TrimSpace(imputUserName)
	nameLowerCase := strings.ToLower(nameSpaceRemoval)
	formattedUserName := strings.Title(nameLowerCase)
	return formattedUserName
}
