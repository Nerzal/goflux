package kustomize

import (
	"io/ioutil"

	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

type Service interface {
	FetchRessources(path string) (error, []string)
	CreateBase(path string, ressources []string) error
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

		if fileInfo.Name() == "kustomize.yaml" {
			continue
		}

		result = append(result, fileInfo.Name())
	}

	return nil, result
}

func (service *service) CreateBase(path string, ressources []string) error {
	data := Data{
		Resources: ressources,
	}

	err := files.WriteFile(data, path+"/kustomize.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create base kustomize")
	}

	return nil
}
