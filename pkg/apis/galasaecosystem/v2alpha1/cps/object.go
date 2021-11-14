package cps

import "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

type Cps struct {
	Image            string
	ImagePullPolicy  string
	Storage          string
	StorageClassName string
	NodeSelector     map[string]string
}

func New(cpsCrd *v2alpha1.GalasaCpsComponent) *Cps {
	return &Cps{}
}

func (c *Cps) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *Cps) IsReady() bool {
	return true
}
