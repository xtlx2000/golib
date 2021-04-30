package file

import (
	"io/ioutil"
)

func WriteFile(filename, text string) error {
	return ioutil.WriteFile(filename, []byte(text), 0666)
}

func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	return string(content), err
}
