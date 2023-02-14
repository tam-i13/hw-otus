package main

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func eq(f1, f2 string) bool {
	file1, _ := os.ReadFile(f1)
	file2, _ := os.ReadFile(f2)
	return bytes.Equal(file1, file2)
}

func TestCopy(t *testing.T) {
	t.Run("File not exist", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		ErrFileNotExist := errors.New("file not exist")
		err := Copy("/not_exist/test", "/tmp/test.txt", 0, 10)
		require.Equal(t, ErrFileNotExist, err)
	})
	t.Run("Offset exceeds file size", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
		err := Copy("./testdata/input.txt", "/tmp/test.txt", 1000000000000, 10)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})
	t.Run("Unsupported file", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		ErrUnsupportedFile = errors.New("unsupported file")
		err := Copy("/dev/urandom", "/tmp/test.txt", 0, 10)
		require.Equal(t, ErrUnsupportedFile, err)
	})
	t.Run("Vegativ limit value", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		ErrNegativLimit = errors.New("negativ limit value")
		err := Copy("./testdata/input.txt", "/tmp/test.txt", 0, -10)
		require.Equal(t, ErrNegativLimit, err)
	})

	t.Run("Copy test offset 0 limit 0", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		Copy("./testdata/input.txt", "/tmp/test.txt", 0, 0)
		res := eq("testdata/out_offset0_limit0.txt", "/tmp/test.txt")
		require.Equal(t, true, res)
	})
	t.Run("Copy test offset 0 limit 10", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		Copy("./testdata/input.txt", "/tmp/test.txt", 0, 10)
		res := eq("testdata/out_offset0_limit10.txt", "/tmp/test.txt")
		require.Equal(t, true, res)
	})
	t.Run("Copy test offset 100 limit 1000", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		Copy("./testdata/input.txt", "/tmp/test.txt", 100, 1000)
		res := eq("testdata/out_offset100_limit1000.txt", "/tmp/test.txt")
		require.Equal(t, true, res)
	})

	t.Run("Copy test offset -100 limit 100", func(t *testing.T) {
		defer os.Remove("/tmp/test.txt")
		Copy("./testdata/input.txt", "/tmp/test.txt", -100, 100)
		res := eq("testdata/out_offset-100_limit100.txt", "/tmp/test.txt")
		require.Equal(t, true, res)
	})

	t.Run("Copy test offset 0 limit 10000000000", func(t *testing.T) {
		defer os.Remove("/tmp/test1.txt")
		Copy("./testdata/input.txt", "/tmp/test1.txt", 0, 10000000000)
		res := eq("testdata/out_offset0_limit0.txt", "/tmp/test1.txt")
		require.Equal(t, true, res)
	})
}
