package ingress_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/ingress"
)

func TestIngress_Create(t *testing.T) {
	service := ingress.NewService()

	err := service.Create("kycnow-api", "clarilab", "app.foo.kycnow.de", "*.foo.kycnow.de", "my-cert-de-crt", "/api", true, "../../test/ingress")
	if err != nil {
		t.Error(err)
	}
}
