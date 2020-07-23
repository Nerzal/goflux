package configmap

// Config represents a configmap
type Config struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Data       Data     `yaml:"data"`
}

// Metadata are metadata
type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

// Data is used as data
type Data map[string]string
