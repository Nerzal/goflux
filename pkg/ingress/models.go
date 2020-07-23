package ingress

// Data respresents an ingressFile structure
type Data struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

// Annotations are annotations
type Annotations struct {
	KubernetesIoIngressClass                     string `yaml:"kubernetes.io/ingress.class"`
	TraefikIngressKubernetesIoRedirectEntryPoint string `yaml:"traefik.ingress.kubernetes.io/redirect-entry-point"`
	TraefikIngressKubernetesIoRedirectPermanent  string `yaml:"traefik.ingress.kubernetes.io/redirect-permanent"`
	KubernetesIoTLSAcme                          string `yaml:"kubernetes.io/tls-acme"`
}

// Metadata are metadata
type Metadata struct {
	Name        string      `yaml:"name"`
	Namespace   string      `yaml:"namespace"`
	Annotations Annotations `yaml:"annotations"`
}

// Backend is backend
type Backend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort int    `yaml:"servicePort"`
}

// Path is a path
type Path struct {
	Path    string  `yaml:"path"`
	Backend Backend `yaml:"backend"`
}

// HTTP is http
type HTTP struct {
	Paths []Path `yaml:"paths"`
}

// Rules are rules
type Rules struct {
	Host string `yaml:"host"`
	HTTP HTTP   `yaml:"http"`
}

// TLS is the tls config
type TLS struct {
	Hosts      []string `yaml:"hosts"`
	SecretName string   `yaml:"secretName"`
}

// Spec is a spec
type Spec struct {
	Rules []Rules `yaml:"rules"`
	TLS   []TLS   `yaml:"tls"`
}
