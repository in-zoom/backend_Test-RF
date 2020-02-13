package validation

import (
	"github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"Backend_task_RF/validation"
	"testing"
)

func TestInvalidIfEmailUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var input = []string{
		"Alexey@mail.ru", 
		"Alexey@mail.ru", 
		"  Alexey@mail.ru", 
		"   Alexey@mail.ru       "}

	expectedResult := "Пользователь с Email - Alexey@mail.ru уже зарегистрирован"

	for _, currentName := range input {
		rows := sqlmock.NewRows([]string{"email"}).AddRow("Alexey@mail.ru")
		mock.ExpectQuery("SELECT email FROM the_users WHERE email = " + " " + "'" + "Alexey@mail.ru" + "'").WillReturnRows(rows)
		_, err := validation.ValidateEmailUser(currentName, db)
		
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}

func TestInvalidIfEmailUserThereAreCharactersAndNumbers(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var inputInvalidEmail = []string{
		"  Ale-xey@mail.ru",
		"   Alexey;@mail.ru       ",
		"Alexey2mail.ru",
		"   Ale@xey@mail.ru       ",
		"Ale#xey@mail.ru",
		"Alexey&mail.ru",
		"Alexey-mail.ru"}

	expectedResult := "Email введен неверно"
	for _, currentInvalidEmail := range inputInvalidEmail {
		_, err := validation.ValidateEmailUser(currentInvalidEmail, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}

func TestInvalidIfEmptyLine(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	inputInvalidName := []string{"", " ", "      "}
	expectedResult := "Введите E-mail"
	for _, currentInvalidEmailEmptyLine := range inputInvalidName {
		_, err := validation.ValidateEmailUser(currentInvalidEmailEmptyLine, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}
