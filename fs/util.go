package fs

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

var ErrCreatingDirs = errors.New("generating directory structure")
var ErrRemovingFile = errors.New("cleaning the current file path")
var ErrGettingFileInfo = errors.New("getting fileInfo")
var ErrCreatingFile = errors.New("creating file")
var ErrCopyingData = errors.New("copying data")

func generateCleanFilePath(path string) (*os.File, func(), error) {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, os.ModePerm) // 0777
	if err != nil {
		return nil, nil, ErrCreatingDirs
	}

	// Return fileInfo to check for existence of a file
	_, err = os.Stat(path)
	if err == nil {
		err = os.Remove(path)
		if err != nil {
			return nil, nil, ErrRemovingFile
		}
	}

	if !errors.Is(err, os.ErrNotExist) {
		return nil, nil, ErrGettingFileInfo
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, nil, ErrCreatingFile
	}

	return file, func() {file.Close()}, nil

}

func writeFile(file *os.File, data io.Reader) error {
	_, err := io.Copy(file, data)
	if err != nil {
		return err
	}
	return nil
}