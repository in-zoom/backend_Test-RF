package validation

import "errors"

func ValidateAttribute(attribute string) (resultParam string, errr error) {
	whiteAttribute := []string{"id", "user_name", "email"}
	var resultAttribute string
	if attribute != "" {
		for _, paramAttribute := range whiteAttribute {
			if attribute == paramAttribute {
				resultAttribute += attribute
			}
		}
		if resultAttribute == "" {
			return "", errors.New("Неверный параметр группировки")
		} else {
			return "ORDER BY" + " " + resultAttribute, nil
		}
	}
	return "", nil
}
func ValidateOrder(order string) (resultParam string, errr error) {
	whiteOrder := []string{"asc", "desc"}
	var resultOrder string
	if order != "" {
		for _, paramOrder := range whiteOrder {
			if order == paramOrder {
				resultOrder += order
			}
		}
		if resultOrder == "" {
			return "", errors.New("Неверный параметр сортировки")
		} else {
			return resultOrder, nil
		}
	}
	return "", nil
}
