package validation

import (
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"Backend_task_RF/validation"
	"testing"
)

func TestNameСolumnSigns(t *testing.T) {
	imput := []string{"&", "%", "*", "="}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр группировки"
		_, err := validation.ValidateAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
func TestNameСolumnStrings(t *testing.T) {
	imput := []string{" ", "DROP", "5"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр группировки"
		_, err := validation.ValidateAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
func TestSortingSigns(t *testing.T) {
	imput := []string{"&", "%", "*", "=", ";"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр сортировки"
		_, err := validation.ValidateOrder(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
func TestSortingStrings(t *testing.T) {
	imput := []string{" ", "DROP", "5"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр сортировки"
		_, err := validation.ValidateOrder(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}

func TestNegativeValueLimit(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imput := "-1"
	expectedResult := "Значение не может быть отрицательным"
	_, err = validation.ValidateLimit(imput, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}

func TestZeroValueLimit(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imput := "0"
	expectedResult := "Значение не может быть нулевым"
	_, err = validation.ValidateLimit(imput, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}

func TestIncorrectStringsLimit(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imput := []string{" ", "DROP", "&", ";"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Задано некорректное значение"
		_, err := validation.ValidateLimit(imputCurrentItem, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}

func TestNegativeValueOffset(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	
	imput := "-1"
	expectedResult := "Значение не может быть отрицательным"
	_, err = validation.ValidateOffset(imput, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}

func TestIncorrectStringsOffset(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	
	imput := []string{" ", "DROP", "&", ";"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Задано некорректное значение"
		_, err := validation.ValidateOffset(imputCurrentItem, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
