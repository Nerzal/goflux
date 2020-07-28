package config

// Config is used to configure goflux
type Config struct {
	Deployment Deployment `yaml:"deployment,omitempty"`
	HPA        HPA        `yaml:"hpa,omitempty"`
	Secrets    Secrets    `yaml:"secrets,omitempty"`
}

// Secrets is used to seal and find secrets
type Secrets struct {
	SecretFolderName string `yaml:"secretFolderName,omitempty"`
	DevCertURL       string `yaml:"devCertURL,omitempty"`
	TestCertURL      string `yaml:"testCertURL,omitempty"`
	ProdCertURL      string `yaml:"prodCertURL,omitempty"`
}

// HPA is used to configure the creation of hpa files
type HPA struct {
	MinReplicas int `yaml:"minReplicas,omitempty"`
	MaxReplicas int `yaml:"maxReplicas,omitempty"`
}

// Deployment is used to configure the creation of deployment files
type Deployment struct {
	ImagePullSecret string      `yaml:"imagePullSecret,omitempty"`
	Annotations     Annotations `yaml:"annotations,omitempty"`
	Ressources      Ressources  `yaml:"ressources,omitempty"`
}

// Ressources are used as ressources
type Ressources struct {
	Limits   Ressource `yaml:"limits,omitempty"`
	Requests Ressource `yaml:"requests,omitempty"`
}

// Ressource is used inside ressources
type Ressource struct {
	CPU    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}

// Annotations are used as annotations
type Annotations map[string]string
