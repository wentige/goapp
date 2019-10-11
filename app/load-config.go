package app

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfig(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &Config); err != nil {
		return err
	}

	return nil
}
