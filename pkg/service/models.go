package service

type Data struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Spec struct {
	Ports    []Port  `yaml:"ports"`
	Selector Selector `yaml:"selector"`
}

type Selector struct {
	App       string `yaml:"app"`
	Component string `yaml:"component"`
}

type Port struct {
	Port       int    `yaml:"port"`
	TargetPort int    `yaml:"targetPort"`
	Protocol   string `yaml:"protocol"`
}

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Labels    Labels `yaml:"labels"`
}

type Labels struct {
	App       string `yaml:"app"`
	Component string `yaml:"component"`
}
