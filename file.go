package telegroid

import (
	"io"
	"os"
	"bytes"
	"path"
)

type InputFile struct {
	Name   string
	Reader io.Reader
}

func FromFile(filepath string) (value InputFile, err error) {
	value = InputFile{}
	file, err := os.Open(filepath)
	if err == nil {
		value.Name = path.Base(file.Name())
		value.Reader = file
	}
	return
}

func FromData(name string, data []byte) InputFile {
	return InputFile{
		Name: name,
		Reader: bytes.NewReader(data),
	}
}
