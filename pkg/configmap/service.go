package configmap

import (
	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

// Service is used to create configmaps
type Service interface {
	Create(component, namespace, path string, data Data) error
}

type service struct{}

// NewService creates a new instance of Service
func NewService() Service {
	return &service{}
}

func (service *service) Create(component, namespace, path string, data Data) error {
	config := Config{
		APIVersion: "v1",
		Kind:       "ConfigMap",
		Metadata: Metadata{
			Name:      component,
			Namespace: namespace,
		},
		Data: data,
	}

	err := files.WriteFile(config, path+"/configmap.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create deployment file")
	}

	return nil
}
