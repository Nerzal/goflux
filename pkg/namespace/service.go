package namespace

import (
	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

// Service provides functionality to create namespace files
type Service interface {
	Create(name string, path string) error
}

type service struct{}

// NewService creates a new instance of service
func NewService() Service {
	return &service{}
}

func (service *service) Create(name string, path string) error {
	data := Data{
		APIVersion: "v1",
		Kind:       "Namespace",
		Metadata: Metadata{
			Name: name,
			Annotations: Annotations{
				LinkerdIoInject: "enabled",
			},
		},
	}

	err := files.WriteFile(data, path+"/namespace.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create namespace file")
	}

	return nil
}
