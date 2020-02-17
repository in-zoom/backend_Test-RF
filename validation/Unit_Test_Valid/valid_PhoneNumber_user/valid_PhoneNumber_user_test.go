package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestInvalidNumberFhoneIfEmpty(t *testing.T) {
	imput := ""
	expectedResult := ""
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidNumberFhoneIfStringOfMultipleSpaces(t *testing.T) {
	imput := "     "
	expectedResult := ""
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestInvalidNumberFhoneWithAPlus(t *testing.T) {
	imput := "+77777777777"
	expectedResult := "+77777777777"
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestInvalidNumberFhoneWithSpace(t *testing.T) {
	imput := "8 777 777 77 77"
	expectedResult := "8 777 777 77 77"
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestInvalidNumberFhoneWithSpaceAndWithAPlus(t *testing.T) {
	imput := "+7 777 777 77 77"
	expectedResult := "+7 777 777 77 77"
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestInvalidNumberWithDash(t *testing.T) {
	imput := "8-777-777-77-77"
	expectedResult := "8-777-777-77-77"
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestInvalidNumberWithDashAndWithAPlu(t *testing.T) {
	imput := "+7-777-777-77-77"
	expectedResult := "+7-777-777-77-77"
	actualResult, err := validation.ValidatePhoneNumber(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
