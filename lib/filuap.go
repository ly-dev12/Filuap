package filuap

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	path      string
	ReadBytes int
}

func (f *File) DataBuf() []byte {
	file, _ := os.Open(f.path)
	result, _ := ioutil.ReadAll(file)
	return result
}

func Data(f FileInter) []byte {
	b := f.DataBuf()
	return b
}

func WriteDataFile(data []byte, nameFile string) (string, error) {
	err := ioutil.WriteFile(nameFile, data, 0777)
	if err != nil {
		err := errors.New("A problem has ocurred when you tryed create file!")
		return "", err
	}

	return "Success create file", nil
}

type FileInter interface {
	DataBuf() []byte
}

func ReadFile(path string) *File {

	path = filepath.Join(path)
	return &File{path: path}
}
