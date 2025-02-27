package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ResolveAmounts(t *testing.T) {
	in := []string{"wootether", "10000", "1ether", "0.1ether", "23948ether"}
	out := ResolveAmounts(in)
	require.ElementsMatch(t, out, []string{"wootether", "10000", "1000000000000000000", "100000000000000000", "23948000000000000000000"})
}
