package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestСolumnId(t *testing.T) {
	input := "id"
	expectedResult := "ORDER BY id"
	actualResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestСolumnUserName(t *testing.T) {
	input := "user_name"
	expectedResult := "ORDER BY user_name"
	actualResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestСolumnEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	expectedResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, expectedResult)
}
func TestSortingAscending(t *testing.T) {
	input := "asc"
	expectedResult := "asc"
	actualResult, err := validation.ValidateOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestSortingDescending(t *testing.T) {
	input := "desc"
	expectedResult := "desc"
	actualResult, err := validation.ValidateOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestSortingEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfLimitEqualEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestIfOffsetEqualEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
