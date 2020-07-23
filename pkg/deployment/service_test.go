package deployment_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/deployment"
)

func TestDeployment_Create(t *testing.T) {
	service := deployment.NewService()
	err := service.Create("kycnow-api", "clarilab", "myCoolSecret", "../../test/deployment")
	if err != nil {
		t.Error(err)
	}
}
