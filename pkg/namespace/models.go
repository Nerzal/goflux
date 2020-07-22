package namespace

// Data are namespace data
type Data struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
}

// Annotations are annotations
type Annotations struct {
	LinkerdIoInject string `yaml:"linkerd.io/inject"`
}

// Metadata are metadata
type Metadata struct {
	Name        string      `yaml:"name"`
	Annotations Annotations `yaml:"annotations"`
}
