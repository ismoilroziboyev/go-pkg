package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidPhoneNumber(t *testing.T) {
	testCases := []struct {
		Phone   string
		IsValid bool
	}{
		{
			Phone:   "+998939444321",
			IsValid: false,
		},
		{
			Phone:   "939444321",
			IsValid: false,
		},
		{
			Phone:   "998939444321",
			IsValid: true,
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("%d - subtest", index), func(t *testing.T) {
			isValid := IsValidPhone(testCase.Phone)
			require.Equal(t, testCase.IsValid, isValid)
		})
	}
}
