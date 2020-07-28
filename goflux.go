package goflux

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Nerzal/goflux/pkg/config"
	"github.com/Nerzal/goflux/pkg/configmap"
	"github.com/Nerzal/goflux/pkg/deployment"
	"github.com/Nerzal/goflux/pkg/hpa"
	"github.com/Nerzal/goflux/pkg/ingress"
	"github.com/Nerzal/goflux/pkg/kustomize"
	"github.com/Nerzal/goflux/pkg/namespace"
	"github.com/Nerzal/goflux/pkg/service"
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
	CreateHpa(component, namespace string, minReplicas, maxReplicas int, path ...string) error
	CreateIngress(component, namespace, host, tlsHost, secretName, endpoint string, tlsAcme bool, path ...string) error
}

type goflux struct {
	service    service.Service
	kustomize  kustomize.Service
	namespace  namespace.Service
	deployment deployment.Service
	configmap  configmap.Service
	hpa        hpa.Service
	ingress    ingress.Service
	config     *config.Config
}

// New creates a new instance of Goflux
func New() Goflux {
	service := service.NewService()
	kustomize := kustomize.NewService()
	namespace := namespace.NewService()
	deployment := deployment.NewService()
	configmap := configmap.NewService()
	hpa := hpa.NewService()
	ingress := ingress.NewService()

	myConfig, err := config.LoadConfig("./goflux.yaml")
	if err != nil {
		fmt.Println("no goflux.yaml config found in current directory. continue without config")
	}

	return &goflux{
		service:    service,
		kustomize:  kustomize,
		namespace:  namespace,
		deployment: deployment,
		configmap:  configmap,
		hpa:        hpa,
		ingress:    ingress,
		config:     myConfig,
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
	basePath := fmt.Sprintf("./%s/base", component)

	err := goflux.service.Create(component, namespace, basePath)
	if err != nil {
		return err
	}

	var imagePullSecrets string
	if goflux.config != nil {
		imagePullSecrets = goflux.config.Deployment.ImagePullSecret
	}

	err = goflux.deployment.Create(component, namespace, imagePullSecrets, basePath)
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

func (goflux *goflux) CreateHpa(component, namespace string, minReplicas, maxReplicas int, path ...string) error {
	basePath := "."

	if len(path) != 0 {
		basePath = path[0]
	}

	return goflux.hpa.Create(component, namespace, minReplicas, maxReplicas, basePath)
}

func (goflux *goflux) CreateIngress(component, namespace, host, tlsHost, secretName, endpoint string, tlsAcme bool, path ...string) error {
	basePath := "."

	if len(path) != 0 {
		basePath = path[0]
	}

	return goflux.ingress.Create(component, namespace, host, tlsHost, secretName, endpoint, tlsAcme, basePath)
}

func (goflux *goflux) createFolder(projectName string, folderName string) error {
	newpath := filepath.Join(".", projectName, folderName)
	return os.MkdirAll(newpath, 0750)
}
