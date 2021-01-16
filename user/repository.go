package user

import (
	"errors"
	"io/ioutil"
	"os"
)

type connector struct {
	path string
}

func (s connector) readFile() ([]byte, error) {
	jsonFile, err := os.OpenFile("user.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, errors.New("Open file error")
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if len(data) == 0 {
		data = []byte("[]")
	}

	if err != nil {
		return nil, errors.New("Read file error")
	}
	return data, nil
}

func (s connector) writeFile(data []byte) error {
	err := ioutil.WriteFile("user.json", data, 0644)
	if err != nil {
		return errors.New("Covert file json to string error")
	}
	return nil
}
