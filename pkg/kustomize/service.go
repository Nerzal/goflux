package kustomize

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

// Service is used to create kustomization files
type Service interface {
	FetchRessources(path string) ([]string, error)
	FetchPatches(path string) ([]string, error)
	Create(path, namespace string, ressources, patches, bases []string) error
}

type service struct{}

// NewService creates a new instance of Service
func NewService() Service {
	return &service{}
}

func (service *service) FetchRessources(path string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch resources")
	}

	var result []string
	var secrets []string

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if strings.Contains(fileName, "patch") {
			continue
		}

		if fileInfo.IsDir() && fileName == "_secrets" {
			fetchResult, err := service.FetchRessources(path + "/" + fileName)
			if err != nil {
				return nil, errors.Wrap(err, "could not fetch secret ressources")
			}

			secrets = fetchResult
			continue
		}

		if fileName == "kustomization.yaml" {
			continue
		}

		result = append(result, fileName)
	}

	for _, secret := range secrets {
		result = append(result, fmt.Sprintf("_secrets/%s", secret))
	}

	return result, nil
}

func (service *service) FetchPatches(path string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch patches")
	}

	var result []string

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if !strings.Contains(fileName, "patch") {
			continue
		}

		result = append(result, fileName)
	}

	return result, nil
}

func (service *service) Create(path, namespace string, ressources, patches, bases []string) error {
	data := Data{
		Namespace: namespace,
		Resources: ressources,
		Bases:     bases,
		Patches:   patches,
	}

	err := files.WriteFile(data, path+"/kustomization.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create base kustomize")
	}

	return nil
}
