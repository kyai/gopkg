package file

import (
	"io/ioutil"
	"os"
)

func Reader(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(text), err
}

func Writer(path, text string) error {
	return ioutil.WriteFile(path, []byte(text), 0666)
}

func Remove(path string) error {
	return os.Remove(path)
}
