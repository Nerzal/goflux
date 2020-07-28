package secret_test

import (
	"io/ioutil"
	"testing"

	"github.com/Nerzal/goflux/pkg/secret"
	"gopkg.in/yaml.v2"
)

func TestSecret_SealSecrets(t *testing.T) {
	values := make(map[string]string)
	values["User"] = "Prinzessin"
	values["Passwort"] = "Vong Theelinger"

	// Arrange
	unsealedSecret := secret.UnsealedData{
		APIVersion: "v1",
		Kind:       "Secret",
		Type:       "Opaque",
		Metadata: secret.Metadata{
			Name:      "kycnow-api",
			Namespace: "clarilab",
		},
		Data: values,
	}

	bytes, err := yaml.Marshal(unsealedSecret)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("../../test/secret/testSecret.yaml", bytes, 0644)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("../../test/secret/testSecret2.yaml", bytes, 0644)
	if err != nil {
		t.Error(err)
	}

	service := secret.NewService()

	// Act
	err = service.SealSecrets("../../test/secret", "")
	if err != nil {
		t.Error(err)
	}

	// Assert
}
