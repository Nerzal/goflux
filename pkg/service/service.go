package service

// Service is a service service, lol
type Service interface {
	New(name, namespace string) Data
}

type service struct{}

// NewService creates a new instance of Service
func NewService() Service {
	return &service{}
}

func (service *service) New(name, namespace string) Data {
	data := Data{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name:      name,
			Namespace: namespace,
			Labels: Labels{
				App:       namespace,
				Component: name,
			},
		},
		Spec: Spec{
			Ports: []Port{
				{
					Port:       80,
					TargetPort: 8080,
					Protocol:   "TCP",
				},
			},
			Selector: Selector{
				App:       namespace,
				Component: name,
			},
		},
	}

	return data
}
