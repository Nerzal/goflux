package config_test

import (
	"io/ioutil"
	"testing"

	"github.com/Nerzal/goflux/pkg/config"
	"gopkg.in/yaml.v2"
)

func TestConfig_Load(t *testing.T) {
	annotations := config.Annotations{}
	annotations["fluxcd.io/automated"] = "true"
	annotations["fluxcd.io/tag.foo-bar"] = "glob:dev-*"

	defaultConfig := config.Config{
		Deployment: config.Deployment{
			ImagePullSecret: "mySecret",
			Annotations:     annotations,
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
		HPA: config.HPA{
			MaxReplicas: 12,
			MinReplicas: 4,
		},
		Secrets: config.Secrets{
			DevCertURL:       "foo.dev.cert",
			TestCertURL:      "foo.test.cert",
			ProdCertURL:      "foo.prod.cert",
			SecretFolderName: "_secrets",
		},
	}

	defaultConfigData, err := yaml.Marshal(defaultConfig)
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile("../../test/config/goflux.yaml", defaultConfigData, 0644)

	_, err = config.LoadConfig("../../test/config/goflux.yaml")
	if err != nil {
		t.Error(err)
	}
}
