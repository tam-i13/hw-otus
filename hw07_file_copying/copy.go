package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrFileNotExist          = errors.New("file not exist")
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrOpenFile              = errors.New("open file error")
	ErrCreateeFile           = errors.New("create file error")
	ErrNegativLimit          = errors.New("negativ limit value")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileStat, err := os.Stat(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileNotExist
		}
	}

	if fileStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if fileStat.Size() == 0 {
		return ErrUnsupportedFile
	}

	fileFrom, errOpen := os.Open(fromPath)

	if errOpen != nil {
		return ErrOpenFile
	}
	defer fileFrom.Close()

	fileDist, errCreateeFile := os.Create(toPath)

	if errCreateeFile != nil {
		return ErrCreateeFile
	}
	defer fileDist.Close()

	if limit == 0 {
		limit = fileStat.Size()
	}
	if limit < 0 {
		return ErrNegativLimit
	}

	if offset > 0 {
		_, errFileSeek := fileFrom.Seek(offset, io.SeekStart)
		if errFileSeek != nil {
			return errFileSeek
		}
	}
	if offset < 0 {
		_, errFileSeek := fileFrom.Seek(offset, io.SeekEnd)
		if errFileSeek != nil {
			return errFileSeek
		}
	}
	_, errCopy := io.CopyN(fileDist, fileFrom, limit)
	if errCopy != nil {
		if errors.Is(errCopy, io.EOF) {
			return nil
		}
		return errCopy
	}

	return nil
}
