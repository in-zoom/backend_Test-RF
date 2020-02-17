package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestInvalidUserPassword(t *testing.T) {
	inputPassword := []string{
		"      eqrtreyeys1w!  ",
		"ADCE&WUewqi",
		"jshfaj;lda",
		" Qw@rtuYoiM ", 
		"    A123g", 
		"1d"}
	expectedResult := "Пароль должен состоять хотя бы из 6 символов, может содержать буквы, цифры, знаки -, _"
	for _, currentPassword := range inputPassword {
		_, err := validation.ValidatePasswordUsers(currentPassword)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}

func TestInvalidUserPasswordIfEmpty(t *testing.T) {
	inputPassword := []string{"", "   ", "     "}
	expectedResult := "Не задан пароль"
	for _, currentPassword := range inputPassword {
		_, err := validation.ValidatePasswordUsers(currentPassword)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}
