package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	password := RandomString("pass-", 5)
	hashedPassword := HashPassword(password)
	valid := CheckPasswordHash(password, hashedPassword)

	require.NotEmpty(t, hashedPassword)
	require.True(t, valid)
}
