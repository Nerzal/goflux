package ingress

import (
	"strconv"

	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

// Service is used to create ingress files
type Service interface {
	Create(component, namespace, host, tlsHost, secretName, endpoint string, tlsAcme bool, path string) error
}

type service struct{}

// NewService creates a new instance of Service
func NewService() Service {
	return &service{}
}

func (service *service) Create(component, namespace, host, tlsHost, secretName, endpoint string, tlsAcme bool, path string) error {
	data := Data{
		APIVersion: "networking.k8s.io/v1beta1",
		Kind:       "Ingress",
		Metadata: Metadata{
			Name:      component,
			Namespace: namespace,
			Annotations: Annotations{
				KubernetesIoTLSAcme:                          strconv.FormatBool(tlsAcme),
				KubernetesIoIngressClass:                     "traefik",
				TraefikIngressKubernetesIoRedirectEntryPoint: "https",
				TraefikIngressKubernetesIoRedirectPermanent:  "true",
			},
		},
		Spec: Spec{
			Rules: []Rules{
				{
					Host: host,
					HTTP: HTTP{
						Paths: []Path{
							{
								Path: endpoint,
								Backend: Backend{
									ServiceName: component,
									ServicePort: 80,
								},
							},
						},
					},
				},
			},
			TLS: []TLS{
				{
					Hosts:      []string{tlsHost},
					SecretName: secretName,
				},
			},
		},
	}

	err := files.WriteFile(data, path+"/ingress.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create ingress file")
	}

	return nil
}
