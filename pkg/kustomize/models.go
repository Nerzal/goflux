package kustomize

type Data struct {
	APIVersion string   `yaml:"apiVersion,omitempty"`
	Kind       string   `yaml:"kind,omitempty"`
	Namespace  string   `yaml:"namespace,omitempty"`
	Bases      []string `yaml:"bases,omitempty"`
	Resources  []string `yaml:"resources,omitempty"`
	Patches    []string `yaml:"patches,omitempty"`
}
