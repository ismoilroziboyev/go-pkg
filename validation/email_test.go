package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		Email   string
		IsValid bool
	}{
		{
			Email:   "email@example.com",
			IsValid: true,
		},
		{
			Email:   "password",
			IsValid: false,
		},
		{
			Email:   "test@gmail.com",
			IsValid: true,
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("%d - subtest", index), func(t *testing.T) {
			isValid := IsValidEmail(testCase.Email)
			require.Equal(t, testCase.IsValid, isValid)
		})
	}
}
