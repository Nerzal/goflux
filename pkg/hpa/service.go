package hpa

import (
	"fmt"

	"github.com/Nerzal/goflux/pkg/files"
	"github.com/pkg/errors"
)

// Service is used to create hpa files
type Service interface {
	Create(component, namespace string, minReplicas, maxReplicas int, path string) error
}

type service struct{}

// NewService creates a new instance of service
func NewService() Service {
	return &service{}
}

func (service *service) Create(component, namespace string, minReplicas, maxReplicas int, path string) error {
	if minReplicas == 0 {
		minReplicas = 2
	}

	if maxReplicas == 0 {
		maxReplicas = 4
	}

	query := fmt.Sprintf(`sum(irate(request_total{namespace="%s", deployment="%s", direction="inbound"}[2m]))`, namespace, component)

	data := Data{
		APIVersion: "autoscaling/v2beta1",
		Kind:       "HorizontalPodAutoscaler",
		Metadata: Metadata{
			Name:      component,
			Namespace: namespace,
			Annotations: Annotations{
				MetricConfigObjectRequestRatePrometheusPerReplica: "true",
				MetricConfigObjectRequestRatePrometheusQuery:      query,
			},
		},
		Spec: Spec{
			ScaleTargetRef: ScaleTargetRef{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       component,
			},
			MinReplicas: minReplicas,
			MaxReplicas: maxReplicas,
			Metrics: []Metrics{
				{
					Type: "Object",
					Object: Object{
						MetricName:  "request-rate",
						TargetValue: 5,
						Target: Target{
							APIVersion: "v1",
							Kind:       "Pod",
							Name:       "dummy-pod",
						},
					},
				},
			},
		},
	}

	err := files.WriteFile(data, path+"/hpa.yaml")
	if err != nil {
		return errors.Wrap(err, "could not create hpa file")
	}

	return nil
}
