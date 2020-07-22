package namespace_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/namespace"
)

func TestNamespace_Create(t *testing.T) {
	service := namespace.NewService()
	err := service.Create("clarilab", "../../test")
	if err != nil {
		t.Error(err)
	}
}
