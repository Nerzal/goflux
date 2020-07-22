package goflux

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Nerzal/goflux/pkg/service"
	"gopkg.in/yaml.v2"
)

// Goflux is used to create service, ingress, namespace files etc.
type Goflux interface {
	Initialize(component string) error
	CreateBase(component, namespace string) error
}

type goflux struct {
	service service.Service
}

// New creates a new instance of Goflux
func New() Goflux {
	service := service.NewService()
	return &goflux{
		service: service,
	}
}

func (goflux *goflux) Initialize(projectName string) error {
	err := goflux.createFolder(projectName, "base")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = goflux.createFolder(projectName, "dev")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = goflux.createFolder(projectName, "test")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = goflux.createFolder(projectName, "prod")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (goflux *goflux) CreateBase(component, namespace string) error {
	serviceData := goflux.service.New(component, namespace)
	binaryData, err := yaml.Marshal(serviceData)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("./%s/base/service.yaml", component), binaryData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (goflux *goflux) createFolder(projectName string, folderName string) error {
	newpath := filepath.Join(".", projectName, folderName)
	return os.MkdirAll(newpath, 0777)
}
