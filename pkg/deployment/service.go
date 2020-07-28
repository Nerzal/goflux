package deployment

import (
	"github.com/Nerzal/goflux/pkg/config"
	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

// Service is used to create deployment files
type Service interface {
	Create(component, namespace, imagePullSecret, path string, configs ...*config.Config) error
	CreatePatch(component, namespace, imagePullSecret, path string, configs ...*config.Config) error
}

type service struct{}

// NewService creates a new instance of Service
func NewService() Service {
	return &service{}
}

func (service *service) Create(component, namespace, imagePullSecret, path string, configs ...*config.Config) error {
	var currentConfig *config.Config
	if len(configs) != 0 {
		currentConfig = configs[0]
	}

	ressources := &Ressources{
		Limits:   Ressource{},
		Requests: Ressource{},
	}

	if currentConfig != nil {
		imagePullSecret = currentConfig.Deployment.ImagePullSecret

		ressources.Limits.CPU = currentConfig.Deployment.Ressources.Limits.CPU
		ressources.Limits.Memory = currentConfig.Deployment.Ressources.Limits.Memory

		ressources.Requests.CPU = currentConfig.Deployment.Ressources.Requests.CPU
		ressources.Requests.Memory = currentConfig.Deployment.Ressources.Requests.Memory
	}

	data := Data{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      component,
			Namespace: namespace,
			Labels: Labels{
				App:       namespace,
				Component: component,
			},
		},
		Spec: Spec{
			Selector: Selector{
				MatchLabels: MatchLabels{
					App:       namespace,
					Component: component,
				},
			},
			Template: Template{
				Metadata: Metadata{
					Labels: Labels{
						App:       namespace,
						Component: component,
					},
				},
				Spec: TemplateSpec{
					Containers: []Container{
						{
							Name: component,
							Ports: []Ports{
								{
									ContainerPort: 8080,
								},
							},
							LivenessProbe: Probe{
								HTTPGet: HTTPGet{
									Path: "live",
									Port: 8086,
								},
								InitialDelaySeconds: 10,
								PeriodSeconds:       3,
								TimeoutSeconds:      5,
							},
							ReadinessProbe: Probe{
								HTTPGet: HTTPGet{
									Path: "live",
									Port: 8086,
								},
								InitialDelaySeconds: 10,
								PeriodSeconds:       3,
								TimeoutSeconds:      5,
							},
							Ressources: ressources,
						},
					},
					ImagePullSecrets: []ImagePullSecrets{
						{
							Name: imagePullSecret,
						},
					},
				},
			},
		},
	}

	err := files.WriteFile(data, path+"/deployment.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create deployment file")
	}

	return nil
}

func (service *service) CreatePatch(component, namespace, imagePullSecret, path string, configs ...*config.Config) error {
	var currentConfig *config.Config
	if len(configs) != 0 {
		currentConfig = configs[0]
	}

	ressources := &Ressources{
		Limits:   Ressource{},
		Requests: Ressource{},
	}

	if currentConfig != nil {
		imagePullSecret = currentConfig.Deployment.ImagePullSecret

		ressources.Limits.CPU = currentConfig.Deployment.Ressources.Limits.CPU
		ressources.Limits.Memory = currentConfig.Deployment.Ressources.Limits.Memory

		ressources.Requests.CPU = currentConfig.Deployment.Ressources.Requests.CPU
		ressources.Requests.Memory = currentConfig.Deployment.Ressources.Requests.Memory
	}

	data := Data{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      component,
			Namespace: namespace,
			Labels: Labels{
				App:       namespace,
				Component: component,
			},
		},
		Spec: Spec{
			Selector: Selector{
				MatchLabels: MatchLabels{
					App:       namespace,
					Component: component,
				},
			},
			Template: Template{
				Metadata: Metadata{
					Labels: Labels{
						App:       namespace,
						Component: component,
					},
				},
				Spec: TemplateSpec{
					Containers: []Container{
						{
							Name:       component,
							Ressources: ressources,
						},
					},
				},
			},
		},
	}

	err := files.WriteFile(data, path+"/deployment-patch.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create deployment file")
	}

	return nil
}
