package service

// Data is used as structure for service files
type Data struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

// Spec holds spec options
type Spec struct {
	Ports    []Port   `yaml:"ports"`
	Selector Selector `yaml:"selector"`
}

// Selector holds selector options
type Selector struct {
	App       string `yaml:"app"`
	Component string `yaml:"component"`
}

// Port holds port data
type Port struct {
	Port       int    `yaml:"port"`
	TargetPort int    `yaml:"targetPort"`
	Protocol   string `yaml:"protocol"`
}

// Metadata holds metadata
type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Labels    Labels `yaml:"labels"`
}

// Labels holds label data
type Labels struct {
	App       string `yaml:"app"`
	Component string `yaml:"component"`
}
