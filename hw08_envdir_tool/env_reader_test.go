package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	ErrIsNotFolderTest    = errors.New("is not folder")
	ErrFolderNotExistTest = errors.New("folder not exist")
	nameFile              = "testvalue"
)

func TestReadDir(t *testing.T) {
	tests := []struct {
		start        string
		expectedErr1 error
	}{
		{"./main.go", ErrIsNotFolderTest},
		{"./folder_not_exist", ErrFolderNotExistTest},
	}

	tests2 := []struct {
		content []byte
		res     string
	}{
		{[]byte("test_value"), "test_value"},
		{[]byte("test_value 	"), "test_value"},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("case from test %d", i), func(t *testing.T) {
			t.Parallel()
			_, err := ReadDir(tt.start)
			require.Equal(t, tt.expectedErr1, err)
		})
	}

	for i, tt2 := range tests2 {
		tt2 := tt2
		t.Run(fmt.Sprintf("case from test2 %d", i), func(t *testing.T) {
			t.Parallel()

			strContent := strings.Builder{}
			strContent.Write(tt2.content)

			tmpDir, _ := os.MkdirTemp("", "tmpdir")
			defer os.RemoveAll(tmpDir)
			filePath := path.Join(tmpDir, nameFile)
			tmpFile, _ := os.Create(filePath)
			defer os.Remove(tmpFile.Name())
			if _, err := tmpFile.Write(tt2.content); err != nil {
				log.Fatal(err)
			}
			if err := tmpFile.Close(); err != nil {
				log.Fatal(err)
			}
			res, err := ReadDir(tmpDir)
			require.Equal(t, nil, err)
			resVal := res[nameFile].Value
			require.Equal(t, tt2.res, resVal)
		})
	}

	t.Run("BAR test", func(t *testing.T) {
		res, err := ReadDir("./testdata/env")
		require.Equal(t, nil, err)
		resVal := res["BAR"].Value
		require.Equal(t, "bar", resVal)
	})

	t.Run("Set test empty env value", func(t *testing.T) {
		content := []byte("")

		tmpDir, _ := os.MkdirTemp("", "tmpdir")
		defer os.RemoveAll(tmpDir)

		filePath := path.Join(tmpDir, nameFile)
		tmpFile, _ := os.Create(filePath)
		defer os.Remove(tmpFile.Name())
		if _, err := tmpFile.Write(content); err != nil {
			log.Fatal(err)
		}
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}

		res, err := ReadDir(tmpDir)
		require.Equal(t, nil, err)
		resVal := res[nameFile].NeedRemove
		require.Equal(t, true, resVal)
	})
}
