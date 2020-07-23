package hpa

// Data is used for hpa files
type Data struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

// Annotations are annotations
type Annotations struct {
	MetricConfigObjectRequestRatePrometheusQuery      string `yaml:"metric-config.object.request-rate.prometheus/query,omitempty"`
	MetricConfigObjectRequestRatePrometheusPerReplica string `yaml:"metric-config.object.request-rate.prometheus/per-replica,omitempty"`
}

// Metadata are metadata
type Metadata struct {
	Name        string      `yaml:"name"`
	Namespace   string      `yaml:"namespace"`
	Annotations Annotations `yaml:"annotations"`
}

// ScaleTargetRef is scaleTargetRef
type ScaleTargetRef struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Name       string `yaml:"name"`
}

// Target is target
type Target struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Name       string `yaml:"name"`
}

// Object is object
type Object struct {
	MetricName  string `yaml:"metricName"`
	Target      Target `yaml:"target"`
	TargetValue int    `yaml:"targetValue"`
}

// Metrics are metrics
type Metrics struct {
	Type   string `yaml:"type"`
	Object Object `yaml:"object"`
}

// Spec is spec
type Spec struct {
	ScaleTargetRef ScaleTargetRef `yaml:"scaleTargetRef"`
	MinReplicas    int            `yaml:"minReplicas"`
	MaxReplicas    int            `yaml:"maxReplicas"`
	Metrics        []Metrics      `yaml:"metrics"`
}
