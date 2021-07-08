package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func WriteFile(filename, text string) error {
	return ioutil.WriteFile(filename, []byte(text), 0666)
}

func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	return string(content), err
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
