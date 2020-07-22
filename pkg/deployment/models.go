package deployment

// Data is used as deployment file structure
type Data struct {
	APIVersion string   `yaml:"apiVersion,omitempty"`
	Kind       string   `yaml:"kind,omitempty"`
	Metadata   Metadata `yaml:"metadata,omitempty"`
	Spec       Spec     `yaml:"spec,omitempty"`
}

// Labels are labels
type Labels struct {
	App       string `yaml:"app,omitempty"`
	Component string `yaml:"component,omitempty"`
}

// Metadata is metadata
type Metadata struct {
	Name      string `yaml:"name,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	Labels    Labels `yaml:"labels,omitempty"`
}

// MatchLabels are matchLabels
type MatchLabels struct {
	App       string `yaml:"app,omitempty"`
	Component string `yaml:"component,omitempty"`
}

// Selector is a selector
type Selector struct {
	MatchLabels MatchLabels `yaml:"matchLabels,omitempty"`
}

// SecretKeyRef is used for secrets
type SecretKeyRef struct {
	Name string `yaml:"name,omitempty"`
	Key  string `yaml:"key,omitempty"`
}

// ValueFrom is used for secrets
type ValueFrom struct {
	SecretKeyRef SecretKeyRef `yaml:"secretKeyRef,omitempty"`
}

// Env is used for environmentVariables
type Env struct {
	Name      string    `yaml:"name,omitempty"`
	ValueFrom ValueFrom `yaml:"valueFrom,omitempty"`
}

// Ports are ports
type Ports struct {
	ContainerPort int `yaml:"containerPort,omitempty"`
}

// HTTPGet is used for liveness and readiness checks
type HTTPGet struct {
	Path string `yaml:"path,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

// Probe is a liveness or readiness probe
type Probe struct {
	HTTPGet             HTTPGet `yaml:"httpGet,omitempty"`
	InitialDelaySeconds int     `yaml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       int     `yaml:"periodSeconds,omitempty"`
	TimeoutSeconds      int     `yaml:"timeoutSeconds,omitempty"`
}

// Capabilities is used to drop capabilities
type Capabilities struct {
	Drop []string `yaml:"drop,omitempty"`
}

// SecurityContext is used as securityContext
type SecurityContext struct {
	RunAsNonRoot           bool         `yaml:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem bool         `yaml:"readOnlyRootFilesystem,omitempty"`
	RunAsUser              int          `yaml:"runAsUser,omitempty"`
	Capabilities           Capabilities `yaml:"capabilities,omitempty"`
}

// Container is used for containers section
type Container struct {
	Name            string          `yaml:"name,omitempty"`
	Env             []Env           `yaml:"env,omitempty"`
	Ports           []Ports         `yaml:"ports,omitempty"`
	LivenessProbe   Probe           `yaml:"livenessProbe,omitempty"`
	ReadinessProbe  Probe           `yaml:"readinessProbe,omitempty"`
	SecurityContext SecurityContext `yaml:"securityContext,omitempty"`
}

// ImagePullSecrets are used as imagePullSecrets
type ImagePullSecrets struct {
	Name string `yaml:"name,omitempty"`
}

// TemplateSpec is a spec
type TemplateSpec struct {
	Containers       []Container        `yaml:"containers,omitempty"`
	ImagePullSecrets []ImagePullSecrets `yaml:"imagePullSecrets,omitempty"`
}

// Template is a template
type Template struct {
	Metadata Metadata     `yaml:"metadata,omitempty"`
	Spec     TemplateSpec `yaml:"spec,omitempty"`
}

// Spec is a spec inside a template
type Spec struct {
	Selector Selector    `yaml:"selector,omitempty"`
	Template interface{} `yaml:"template,omitempty"` //Due to declarationCycle
}
