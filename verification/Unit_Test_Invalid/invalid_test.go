package validation

import (
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"Backend_task_RF/verification"
	"Backend_task_RF/hashing"
	"testing"
)

func TestInvalidIfThereIsNoUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imputEmail := "Alexey@mail.ru"
	imputPass := "1234567"
	expectedResult := "Введен неверный логин или пароль"

	rows := sqlmock.NewRows([]string{"id", "email", "password"})
	mock.ExpectQuery("SELECT id, email, password FROM the_users").WillReturnRows(rows)
	_, err = verification.VerificationLogin(imputEmail, imputPass, db)

	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfTheUserIsThereButThePasswordIsIncorrect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imputEmail := "Alexey@mail.ru"
	imputPass := "1234567"
	hashPass, _ := hashing.HashPasswordUser("123456")
	expectedResult := "Введен неверный логин или пароль"

	rows := sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(1, "Alexey@mail.ru", hashPass)
	mock.ExpectQuery("SELECT id, email, password FROM the_users").WillReturnRows(rows)
	_, err = verification.VerificationLogin(imputEmail, imputPass, db)

	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfEmailLoginIsNotEntered(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imputEmail := "       "
	imputPass := "1234567"
	expectedResult := "Введите логин"
	_, err = verification.VerificationLogin(imputEmail, imputPass, db)

	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfNoPasswordIsEntered(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	imputEmail := "Alexey@mail.ru"
	imputPass := "     "
	expectedResult := "Введите пароль"
	_, err = verification.VerificationLogin(imputEmail, imputPass, db)

	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}
