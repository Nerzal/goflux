package kustomize

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

type Service interface {
	FetchRessources(path string) (error, []string)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (service *service) FetchRessources(path string) (error, []string) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "could not fetch resourrces"), nil
	}

	var result []string

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		result = append(result, fileInfo.Name())
	}

	return nil, result
}
