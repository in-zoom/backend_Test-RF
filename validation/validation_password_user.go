package validation

import (
	"strings"
	"errors"
	"regexp"
)

func ValidatePasswordUsers(addPassword string) (resultPasswordUser string, err error) {
	PasswordPrepare, err := PreparePassword(addPassword)
	if err != nil {
		return "", err
	}

	if PasswordPrepare == "" {
		return "", errors.New("Не задан пароль")
	}
	
	pattern := `^[A-Za-z0-9_-]{6,25}$`
	matched, err := regexp.Match(pattern, []byte(PasswordPrepare))
	if matched == false || err != nil {
		return "", errors.New("Пароль должен состоять хотя бы из 6 символов, может содержать буквы, цифры, знаки -, _")
	}
    return PasswordPrepare, nil
}

func PreparePassword(imputPassword string) (outputPassword string, err error) {
	passwordSpaceRemoval := strings.TrimSpace(imputPassword)
    return passwordSpaceRemoval, nil
}
