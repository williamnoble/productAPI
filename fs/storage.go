package fs

import "io"

type FileStore interface{
	Save(path string, file io.Reader) error
}