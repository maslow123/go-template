package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	randEmail := RandomString("email", 10)
	randName := RandomString("name", 10)

	testCases := []struct {
		NameOfTest string
		Email      string `json:"email" validate:"required,email"`
		Name       string `json:"name" validate:"required,min=2,max=100"`
		IsError    bool
	}{
		{
			"Passed all of input",
			fmt.Sprintf("%s@gmail.com", randEmail),
			randName,
			false,
		},
		{
			"Invalid format email",
			randEmail,
			randName,
			true,
		},
		{
			"Invalid length of name",
			fmt.Sprintf("%s@gmail.com", randEmail),
			"",
			true,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		valid := IsValid(tc)

		t.Run(tc.NameOfTest, func(t *testing.T) {
			if tc.IsError {
				require.Error(t, valid)
			} else {
				require.NoError(t, valid)
			}
		})
	}
}
