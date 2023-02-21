package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"strings"
)

var (
	ErrIsNotFolder    = errors.New("is not folder")
	ErrFolderNotExist = errors.New("folder not exist")
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func checkDir(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return ErrFolderNotExist
	}
	if !stat.IsDir() {
		return ErrIsNotFolder
	}
	return nil
}

func ReadDir(dir string) (Environment, error) {
	err := checkDir(dir)
	if err != nil {
		return nil, err
	}

	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envMap := make(Environment)

	for _, tmpValue := range readDir {
		if strings.Contains(tmpValue.Name(), "=") {
			break
		}
		targetPath := path.Join(dir, tmpValue.Name())

		fileStat, err := os.Stat(targetPath)
		if err != nil {
			return nil, err
		}
		if fileStat.Size() != 0 {
			file, _ := os.Open(targetPath)
			br := bufio.NewReader(file)
			line, err := br.ReadBytes(byte('\n'))
			if err != nil && err != io.EOF {
				return nil, err
			}
			file.Close()

			line = bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))

			mekeString := strings.Builder{}
			mekeString.Write(line)
			finalValue := mekeString.String()

			finalValue = strings.TrimRight(finalValue, " 	\n")

			envMap[tmpValue.Name()] = EnvValue{
				Value:      finalValue,
				NeedRemove: false,
			}
		} else {
			envMap[tmpValue.Name()] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
		}
	}

	return envMap, nil
}
