package secret

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Service is used to seal secrets
type Service interface {
	SealSecrets(path, certURL string) error
}

type service struct {
}

// NewService creates a new instance of Service
func NewService() Service {
	return &service{}
}

func (service *service) SealSecrets(path, certURL string) error {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "could not seal secrets")
	}

	for _, file := range fileInfos {
		if file.IsDir() {
			continue
		}

		if !strings.Contains(file.Name(), ".yaml") {
			continue
		}

		filePath := path + "/" + file.Name()

		err = service.Base64EncodeValues(filePath)
		if err != nil {
			return nil
		}

		err = service.SealSecret(filePath, certURL)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not seal secret in path: %s", filePath))
		}
	}

	return nil
}

func (service *service) SealSecret(filePath, certURL string) error {
	cmd := exec.Command("kubeseal", fmt.Sprintf("--format yaml --cert %s < %s > %s", certURL, filePath, filePath))
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (service *service) Base64EncodeValues(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(err, "could not base64EncodeValues")
	}

	var secret UnsealedData
	err = yaml.Unmarshal(data, &secret)
	if err != nil {
		return errors.Wrap(err, "could not unmarshal unsealed secret")
	}

	newMap := make(map[string]string)

	for key := range secret.Data {
		value := secret.Data[key]
		value = base64.StdEncoding.EncodeToString([]byte(value))

		newMap[key] = value
	}

	secret.Data = newMap

	outData, err := yaml.Marshal(secret)
	if err != nil {
		return errors.Wrap(err, "could not remarshal unsealed secret")
	}

	err = ioutil.WriteFile(filePath, outData, 0600)
	if err != nil {
		return errors.Wrap(err, "could not write base64 encoded unsealed secret")
	}

	return nil
}
