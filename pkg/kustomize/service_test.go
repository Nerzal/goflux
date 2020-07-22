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

func TestKustomize_CreateBase(t *testing.T) {
	service := kustomize.NewService()

	var ressources []string
	var err error

	t.Run("fetch ressources", func(t *testing.T) {
		err, ressources = service.FetchRessources("../../test/kustomize/base")
		if err != nil {
			t.Error(err)
		}

		if len(ressources) != 2 {
			t.Errorf("Expected 2 files, found: %v", len(ressources))
		}
	})

	t.Run("create base", func(t *testing.T) {
		err = service.CreateBase("../../test/kustomize/base", ressources)
		if err != nil {
			t.Error(err)
		}
	})

}
