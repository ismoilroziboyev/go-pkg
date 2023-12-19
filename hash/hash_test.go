package hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashMD5(t *testing.T) {
	testCases := []struct {
		Str  string
		Hash string
	}{
		{
			Str:  "this is the string",
			Hash: "fa77937febc2a2a9754d326bb88e0b16",
		},
		{
			Str:  "sample example",
			Hash: "28309b223bb9b0802a4068f97e21befd",
		},
		{
			Str:  "go-pkg utils package",
			Hash: "62ceb8d39ea1e323d2cbcf072b9ee7af",
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprintf("%d - subtest", index), func(t *testing.T) {
			hashedStr := HashMD5(testCase.Str)
			require.Equal(t, testCase.Hash, hashedStr)
		})
	}
}
