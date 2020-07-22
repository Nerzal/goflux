package files

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func WriteFile(data interface{}, path string) error {
	binaryData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, binaryData, 0644)
	if err != nil {
		return err
	}

	return nil
}
