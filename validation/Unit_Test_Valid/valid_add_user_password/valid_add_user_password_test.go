package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestValidUserPasswordOnlyNumbers(t *testing.T) {
	inputPassword := "1234561234"
	expectedResult := "1234561234"
	resultPassword, err := validation.ValidatePasswordUsers(inputPassword)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultPassword, expectedResult)
}

func TestValidUserPasswordOnlyLetters(t *testing.T) {
	inputPassword := "QwErTyuioP"
	expectedResult := "QwErTyuioP"
	resultPassword, err := validation.ValidatePasswordUsers(inputPassword)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultPassword, expectedResult)
}

func TestValidUserPasswordNumbersAndLettersDashSign(t *testing.T) {
	inputPassword := "-1x2345y6"
	expectedResult := "-1x2345y6"
	resultPassword, err := validation.ValidatePasswordUsers(inputPassword)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultPassword, expectedResult)
}

func TestValidUserPasswordNumbersAndLetters(t *testing.T) {
	inputPassword := "      s1w23456sdf34          "
	expectedResult := "s1w23456sdf34"

	resultPassword, err := validation.ValidatePasswordUsers(inputPassword)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultPassword, expectedResult)
}
