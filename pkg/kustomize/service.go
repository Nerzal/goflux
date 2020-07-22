package kustomize

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

type Service interface {
	FetchRessources(path string) (error, []string)
	FetchPatches(path string) (error, []string)
	Create(path, namespace string, ressources, patches, bases []string) error
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
		fileName := fileInfo.Name()
		if strings.Contains(fileName, "patch") {
			continue
		}

		if fileInfo.IsDir() && fileName == "_secrets" {
			err, fetchResult := service.FetchRessources(path + "/" + fileName)
			if err != nil {
				return errors.Wrap(err, "could not fetch secret ressources"), nil
			}

			secrets = fetchResult
			continue
		}

		if fileName == "kustomize.yaml" {
			continue
		}

		result = append(result, fileName)
	}

	for _, secret := range secrets {
		result = append(result, fmt.Sprintf("_secrets/%s", secret))
	}

	return nil, result
}

func (service *service) FetchPatches(path string) (error, []string) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "could not fetch patches"), nil
	}

	var result []string

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if !strings.Contains(fileName, "patch") {
			continue
		}

		result = append(result, fileName)
	}

	return nil, result
}

func (service *service) Create(path, namespace string, ressources, patches, bases []string) error {
	data := Data{
		Namespace: namespace,
		Resources: ressources,
		Bases:     bases,
		Patches:   patches,
	}

	err := files.WriteFile(data, path+"/kustomize.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create base kustomize")
	}

	return nil
}
