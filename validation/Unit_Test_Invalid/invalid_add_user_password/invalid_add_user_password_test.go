package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestInvalidUserPassword(t *testing.T) {
	inputPassword := []string{"1", "1a2b3", "    A123", "      s1w!          "}
	expectedResult := "Пароль должен состоять хотя бы из 6 символов"
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
