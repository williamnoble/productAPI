package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalFileSystem struct {
	maximumFileSize int    //1mb
	basePath        string // e.g. /.images
}

func NewLocalFileSystem(basePath string, size int) *LocalFileSystem {
	path, _ := filepath.Abs(basePath)

	fmt.Println(path)
	return &LocalFileSystem{
		maximumFileSize: size,
		basePath:        path,
	}
}

// Conforms to our storage interface

func (local *LocalFileSystem) Save(path string, contents io.Reader) error {
	filePath := local.generateFullPath(path) // takes image path + base path
	fmt.Println("FILEPATH: " + filePath)
	file, closer, err := generateCleanFilePath(filePath)
	if err != nil {
		return err
	}
	defer closer()

	err = writeFile(file, contents)
	if err != nil {
		return err
	}
	return nil
}

func (local *LocalFileSystem) Get(path string) (*os.File, func(), error) {
	file, closer, err := generateCleanFilePath(path)
	if err != nil {
		return nil, nil, err
	}
	return file, closer, nil
}

func (local *LocalFileSystem) generateFullPath(path string) string {
	return filepath.Join(local.basePath, path)
}