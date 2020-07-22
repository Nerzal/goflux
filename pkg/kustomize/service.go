package kustomize

type Service interface {
	FetchRessources()
}

type service struct{}

func NewService() Service {
	return &service{}
}
