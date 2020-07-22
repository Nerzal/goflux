package kustomize

import (
	"fmt"
	"io/ioutil"

	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

type Service interface {
	FetchRessources(path string) (error, []string)
	Create(path string, ressources []string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (service *service) FetchRessources(path string) (error, []string) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "could not fetch resources"), nil
	}

	var result []string
	var secrets []string

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() && fileInfo.Name() == "_secrets" {
			err, fetchResult := service.FetchRessources(path + "/" + fileInfo.Name())
			if err != nil {
				return errors.Wrap(err, "could not fetch secret ressources"), nil
			}

			secrets = fetchResult
			continue
		}

		if fileInfo.Name() == "kustomize.yaml" {
			continue
		}

		result = append(result, fileInfo.Name())
	}

	for _, secret := range secrets {
		result = append(result, fmt.Sprintf("_secrets/%s", secret))
	}

	return nil, result
}

func (service *service) Create(path string, ressources []string) error {
	data := Data{
		Resources: ressources,
	}

	err := files.WriteFile(data, path+"/kustomize.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create base kustomize")
	}

	return nil
}
