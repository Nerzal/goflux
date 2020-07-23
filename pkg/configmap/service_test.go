package configmap_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/configmap"
)

func TestConfigMap_Create(t *testing.T) {
	service := configmap.NewService()

	data := configmap.Data{}
	data["KEYCLOAK_REALM"] = "myKeyCloakRealm"
	data["SERVICE_NAME"] = "kycnow-api"

	err := service.Create("kycnow-api", "clarilab", "../../test/configmap", data)
	if err != nil {
		t.Error(err)
	}
}
