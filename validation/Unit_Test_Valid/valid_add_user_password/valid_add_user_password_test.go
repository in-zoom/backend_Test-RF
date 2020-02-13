package validation

import (
	"github.com/stretchr/testify/assert"
	"Backend_task_5/validation"
	"testing"
)

func TestValidUserPassword(t *testing.T) {
	inputPassword := []string{"123456", "1a2b3c4@5&6", "    A1x2345y6", "      s1w23456sdf34!          "}
	expectedResult := ""
	for _, currentPassword := range inputPassword {
		resultPassword, err := validation.ValidateColor(currentPassword)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, resultPassword, expectedResult)
    }
}
