package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_RF/validation"
	"testing"
)

func TestValidUserName(t *testing.T) {
	inputName := []string{"Alexey", "Alexey    ", "    Alexey", "      Alexey          "}
	expectedResult := "Alexey"
	for _, currentName := range inputName {
		resultNameUser, err := validation.ValidateNameUser(currentName)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, resultNameUser, expectedResult)

	}
}