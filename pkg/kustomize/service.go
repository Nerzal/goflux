package kustomize

type Service interface {
	FetchRessources(path string)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (service *service) FetchRessources(path string) {

}
