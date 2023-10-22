package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	var min int64 = 10
	var max int64 = 5000

	randomInt := RandomInt(min, max)

	require.GreaterOrEqual(t, randomInt, min)
	require.LessOrEqual(t, randomInt, max)
}
