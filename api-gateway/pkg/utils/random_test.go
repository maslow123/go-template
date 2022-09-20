package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandom(t *testing.T) {
	randomStr := RandomString("", 5)
	randomInt := RandomInt(10, 100)

	require.Equal(t, 5, len(randomStr))
	require.LessOrEqual(t, randomInt, int64(100))
}
