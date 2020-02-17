package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestInvalidNumberFhone(t *testing.T) {
	imput := []string{
		"qwertyuio",
		"8 777 777 7f7 77",
		"8-777-777-7f7-77",
		"8-777-777-7f-77",
		"+7-777-77-7f-77"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Номер телефона не может содержать буквы, знаки пунктуации"
		_, err := validation.ValidatePhoneNumber(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}


