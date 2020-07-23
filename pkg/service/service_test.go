package service_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/service"
)

func TestService_Create(t *testing.T) {
	service := service.NewService()

	err := service.Create("kycnow-api", "clarilab", "../../test/service")
	if err != nil {
		t.Error(err)
	}
}
