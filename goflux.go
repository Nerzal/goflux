package goflux

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Nerzal/goflux/pkg/configmap"
	"github.com/Nerzal/goflux/pkg/deployment"
	"github.com/Nerzal/goflux/pkg/kustomize"
	"github.com/Nerzal/goflux/pkg/namespace"
	"github.com/Nerzal/goflux/pkg/service"
	"gopkg.in/yaml.v2"
)

// Goflux is used to create service, ingress, namespace files etc.
type Goflux interface {
	Initialize(component string) error
	CreateBase(component, namespace string) error

	CreateConfigMap(component, namespace string, data configmap.Data, path ...string) error
	CreateDeployment(component, namespace, imagePullSecret string, path ...string) error
	CreateNameSpace(component, namespace string, path ...string) error
	CreateService(component, namespace string, path ...string) error
	CreateKustomization(namespace string, path ...string) error
}

type goflux struct {
	service    service.Service
	kustomize  kustomize.Service
	namespace  namespace.Service
	deployment deployment.Service
	configmap  configmap.Service
}

// New creates a new instance of Goflux
func New() Goflux {
	service := service.NewService()
	kustomize := kustomize.NewService()
	namespace := namespace.NewService()
	deployment := deployment.NewService()
	configmap := configmap.NewService()

	return &goflux{
		service:    service,
		kustomize:  kustomize,
		namespace:  namespace,
		deployment: deployment,
		configmap:  configmap,
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

	err = ioutil.WriteFile(fmt.Sprintf("%s/service.yaml", basePath), binaryData, 0600)
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

func (goflux *goflux) CreateService(component, namespace string, path ...string) error {
	basePath := fmt.Sprintf("./%s/base", component)

	if len(path) != 0 {
		basePath = path[0]
	}

	return goflux.namespace.Create(namespace, basePath)
}

func (goflux *goflux) CreateDeployment(component, namespace, imagePullSecret string, path ...string) error {
	basePath := fmt.Sprintf("./%s/base", component)

	if len(path) != 0 {
		basePath = path[0]
	}

	return goflux.deployment.Create(component, namespace, imagePullSecret, basePath)
}

func (goflux *goflux) CreateConfigMap(component, namespace string, data configmap.Data, path ...string) error {
	basePath := fmt.Sprintf("./%s/base", component)

	if len(path) != 0 {
		basePath = path[0]
	}

	return goflux.configmap.Create(component, namespace, basePath, data)
}

func (goflux *goflux) CreateKustomization(namespace string, path ...string) error {
	localPath := "."

	if len(path) != 0 {
		localPath = path[0]
	}

	ressources, err := goflux.kustomize.FetchRessources(localPath)
	if err != nil {
		return err
	}

	patches, err := goflux.kustomize.FetchPatches(localPath)
	if err != nil {
		return err
	}

	return goflux.kustomize.Create(localPath, namespace, ressources, patches, nil)
}

func (goflux *goflux) createFolder(projectName string, folderName string) error {
	newpath := filepath.Join(".", projectName, folderName)
	return os.MkdirAll(newpath, 0750)
}
