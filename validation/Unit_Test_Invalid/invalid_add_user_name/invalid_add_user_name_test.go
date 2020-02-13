package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestInvalidUserName(t *testing.T) {
	inputName := []string{"Ale xey", "1Alexey", "-Alexey", "    @Alexey", "      Alexey!          "}
	expectedResult := "Имя не может содержать цифры, знаки пунктуации, символы, пробелы"
	for _, currentName := range inputName {
		_, err := validation.ValidateNameUser(currentName)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}

func TestInvalidUserNameIfEmpty(t *testing.T) {
	inputName := []string{"", "   ", "     "}
	expectedResult := "Введите имя"
	for _, currentName := range inputName {
		_, err := validation.ValidateNameUser(currentName)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}
