package deployment_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/config"
	"github.com/Nerzal/goflux/pkg/deployment"
)

func TestDeployment_Create(t *testing.T) {
	service := deployment.NewService()
	err := service.Create("kycnow-api", "clarilab", "myCoolSecret", "../../test/deployment")
	if err != nil {
		t.Error(err)
	}
}

func TestDeployment_CreatePatch(t *testing.T) {
	service := deployment.NewService()
	err := service.CreatePatch("kycnow-api", "clarilab", "myCoolSecret", "../../test/deployment", &config.Config{
		Deployment: config.Deployment{
			Ressources: config.Ressources{
				Limits: config.Ressource{
					CPU:    "60m",
					Memory: "60Mi",
				},
				Requests: config.Ressource{
					CPU:    "40m",
					Memory: "40Mi",
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
}
