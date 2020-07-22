package kustomize_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/kustomize"
)

func TestKustomize_FetchRessources(t *testing.T) {
	service := kustomize.NewService()
	err, files := service.FetchRessources("../../test/kustomize/base")
	if err != nil {
		t.Error(err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, found: %v", len(files))
	}
}
