package metrics

import "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

type Metrics struct {
	Image           string
	ImagePullPolicy string
	NodeSelector    map[string]string
}

func New(metricsCrd *v2alpha1.GalasaMetricsComponent) *Metrics {
	return &Metrics{}
}

func (c *Metrics) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *Metrics) IsReady() bool {
	return true
}
