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

const trimString = " \t\n"

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
	if err := checkDir(dir); err != nil {
		return nil, err
	}

	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envMap := make(Environment)

	for _, tmpValue := range readDir {
		if strings.Contains(tmpValue.Name(), "=") {
			continue
		}
		targetPath := path.Join(dir, tmpValue.Name())

		fileStat, err := os.Stat(targetPath)
		if err != nil {
			return nil, err
		}
		if fileStat.Size() == 0 {
			envMap[tmpValue.Name()] = EnvValue{
				NeedRemove: true,
			}
			continue
		}
		file, _ := os.Open(targetPath)
		br := bufio.NewReader(file)
		line, err := br.ReadBytes(byte('\n'))
		if err != nil && err != io.EOF {
			return nil, err
		}
		file.Close()

		line = bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))

		envMap[tmpValue.Name()] = EnvValue{
			Value:      strings.TrimRight(string(line), trimString),
			NeedRemove: false,
		}
	}

	return envMap, nil
}
