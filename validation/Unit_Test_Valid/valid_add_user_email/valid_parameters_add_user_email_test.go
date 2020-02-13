package validation

import (
	"github.com/DATA-DOG/go-sqlmock"
	"Backend_task_RF/validation"
	"testing"
)

func TestValidEmailUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := []string{"Alexey@mail.ru", "  Alexey@mail.ru", "   Alexey@mail.ru       "}
	for _, currentName := range input {
		rows := sqlmock.NewRows([]string{"name"})
		mock.ExpectQuery("SELECT email FROM the_users  WHERE email = " + " " + "'" + "Alexey@mail.ru" + "'").WillReturnRows(rows)
		_, err := validation.ValidateEmailUser(currentName, db)
		if err != nil {
			t.Error()
		}
	}
}
