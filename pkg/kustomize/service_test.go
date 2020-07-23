package kustomize_test

import (
	"testing"

	"github.com/Nerzal/goflux/pkg/kustomize"
)

func TestKustomize_FetchRessources(t *testing.T) {
	service := kustomize.NewService()
	files, err := service.FetchRessources("../../test/kustomize/base")
	if err != nil {
		t.Error(err)
	}

	if len(files) != 3 {
		t.Errorf("Expected 3 files, found: %v", len(files))
	}
}

func TestKustomize_CreateBase(t *testing.T) {
	service := kustomize.NewService()

	var ressources []string
	var err error

	t.Run("fetch ressources", func(t *testing.T) {
		ressources, err = service.FetchRessources("../../test/kustomize/base")
		if err != nil {
			t.Error(err)
		}

		if len(ressources) != 3 {
			t.Errorf("Expected 3 files, found: %v", len(ressources))
		}
	})

	t.Run("create base", func(t *testing.T) {
		err = service.Create("../../test/kustomize/base", "", ressources, nil, nil)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestKustomize_CreateDevWithSecrets(t *testing.T) {
	service := kustomize.NewService()

	var ressources []string
	var err error

	t.Run("fetch ressources with secrets", func(t *testing.T) {
		ressources, err = service.FetchRessources("../../test/kustomize/dev")
		if err != nil {
			t.Error(err)
		}

		if len(ressources) != 4 {
			t.Errorf("Expected 4 files, found: %v", len(ressources))
		}
	})

	t.Run("create dev with secrets", func(t *testing.T) {
		err = service.Create("../../test/kustomize/dev", "testNameSpace", ressources, nil, nil)
		if err != nil {
			t.Error(err)
		}
	})
}
