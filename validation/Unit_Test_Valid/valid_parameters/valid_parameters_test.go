package validation

import (
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
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

func TestValidIfTheLimitIsGreaterThanTheNumberOfEntriesInTheTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "100"
	expectedResult := "limit 3"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateLimit(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfTheLimitIsLessThanOrEqualToTheNumberOfEntriesInTheTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "2"
	expectedResult := "limit 2"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateLimit(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestIfLimitEqualEmpty(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateLimit(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfTheOffsetIsGreaterThanTheNumberOfEntriesInTheTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "100"
	expectedResult := "offset 2"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfTheOffsetIsEqualToTheNumberOfEntriesInTheTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "3"
	expectedResult := "offset 2"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfUserEnteredOffsetZero(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "0"
	expectedResult := "offset 0"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(2)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestIfOffsetEqualEmpty(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
