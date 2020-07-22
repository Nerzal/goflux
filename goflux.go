package goflux

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Nerzal/goflux/pkg/kustomize"
	"github.com/Nerzal/goflux/pkg/namespace"
	"github.com/Nerzal/goflux/pkg/service"
	"gopkg.in/yaml.v2"
)

// Goflux is used to create service, ingress, namespace files etc.
type Goflux interface {
	Initialize(component string) error
	CreateBase(component, namespace string) error
	CreateNameSpace(component, namespace string, path ...string) error
}

type goflux struct {
	service   service.Service
	kustomize kustomize.Service
	namespace namespace.Service
}

// New creates a new instance of Goflux
func New() Goflux {
	service := service.NewService()
	kustomize := kustomize.NewService()
	namespace := namespace.NewService()

	return &goflux{
		service:   service,
		kustomize: kustomize,
		namespace: namespace,
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

	basePath := fmt.Sprintf("./%s/base", component)

	err = ioutil.WriteFile(fmt.Sprintf("%s/service.yaml", basePath), binaryData, 0644)
	if err != nil {
		return err
	}

	ressources, err := goflux.kustomize.FetchRessources(basePath)
	if err != nil {
		return err
	}

	err = goflux.kustomize.Create(basePath, "", ressources, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (goflux *goflux) CreateNameSpace(component, namespace string, path ...string) error {
	basePath := fmt.Sprintf("./%s/base", component)

	if len(path) != 0 {
		basePath = path[0]
	}

	return goflux.namespace.Create(namespace, basePath)
}

func (goflux *goflux) createFolder(projectName string, folderName string) error {
	newpath := filepath.Join(".", projectName, folderName)
	return os.MkdirAll(newpath, 0777)
}
