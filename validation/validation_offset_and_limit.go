package validation

import (
	"errors"
	"strconv"
)

func ValidateLimit(limit string) (resultLimit string, err error) {
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return "", errors.New("Задано некорректное значение")
		}
		if limitInt >= 0 {
			if limitInt != 0 {
				if limitInt > 27 {
					return "limit" + " " + strconv.Itoa(27), nil
				} else {
					return "limit" + " " + limit, nil
				}
			}
			return "", errors.New("Значение не может быть нулевым")
		}
		return "", errors.New("Значение не может быть отрицательным")
	}
	return "", nil
}

func ValidateOffset(offset string) (resultOffset string, err error) {
	if offset != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return "", errors.New("Задано некорректное значение")
		}
		if offsetInt >= 0 {
			if offsetInt > 26 {
				return "offset" + " " + strconv.Itoa(26), nil
			} else {
				return "offset" + " " + offset, nil
			}
		}
		return "", errors.New("Значение не может быть отрицательным")
	}
	return "", nil
}
