package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Test start sleep", func(t *testing.T) {
		cmd := []string{"sleep", "2"}
		mapEnv := make(Environment)
		exitCode := RunCmd(cmd, mapEnv)
		require.Equal(t, 0, exitCode)
	})
	t.Run("Test start sleep with env", func(t *testing.T) {
		res, _ := ReadDir("./testdata/env")
		cmd := []string{"sleep", "1"}
		exitCode := RunCmd(cmd, res)
		require.Equal(t, 0, exitCode)
		testEnv, ok := os.LookupEnv("BAR")
		require.Equal(t, true, ok)
		require.Equal(t, "bar", testEnv)
	})
}
