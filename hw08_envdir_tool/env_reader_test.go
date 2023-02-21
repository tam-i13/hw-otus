package main

import (
	"errors"
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
	t.Run("Is not folder", func(t *testing.T) {
		_, res := ReadDir("./main.go")
		require.Equal(t, ErrIsNotFolderTest, res)
	})

	t.Run("Is not folder", func(t *testing.T) {
		_, res := ReadDir("./folder_not_exist")
		require.Equal(t, ErrFolderNotExistTest, res)
	})

	t.Run("Set test env", func(t *testing.T) {
		content := []byte("test_value")
		strContent := strings.Builder{}
		strContent.Write(content)

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
		resVal := res[nameFile].Value
		require.Equal(t, strContent.String(), resVal)
	})

	t.Run("Set test env with space", func(t *testing.T) {
		content := []byte("test_value ")
		check := []byte("test_value")
		strContent := strings.Builder{}
		strContent.Write(check)

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
		resVal := res[nameFile].Value
		require.Equal(t, strContent.String(), resVal)
	})

	t.Run("Set test env with tab", func(t *testing.T) {
		content := []byte("test_value	")
		check := []byte("test_value")
		strContent := strings.Builder{}
		strContent.Write(check)

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
		resVal := res[nameFile].Value
		require.Equal(t, strContent.String(), resVal)
	})

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
