package secret

// UnsealedData represents a unsealed secret
type UnsealedData struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Data       Data     `yaml:"data"`
	Type       string   `yaml:"type"`
}

// Metadata is metadata
type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

// Data represents the unsealed NON base 64 encoded key value pairs
type Data map[string]string
