package validation

import (
	"errors"
	"regexp"
	"strings"
)

func ValidatePhoneNumber(PhoneNumber string) (prepareNumberPhone string, err error) {
	preparePhoneNumber := preparePhoneNumber(PhoneNumber)

	if preparePhoneNumber == "" {
		return "", errors.New("Введите номер телефона")
	}

	pattern := `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
	matched, err := regexp.Match(pattern, []byte(preparePhoneNumber))
	if matched == false || err != nil {
		return "", errors.New("Номер телефона не может содержать буквы, знаки пунктуации")
	}
	return preparePhoneNumber, nil
}

func preparePhoneNumber(imputPhoneNumber string) (outputPhoneNumber string) {
	PhoneNumberSpaceRemoval := strings.TrimSpace(imputPhoneNumber)
	return PhoneNumberSpaceRemoval
}
