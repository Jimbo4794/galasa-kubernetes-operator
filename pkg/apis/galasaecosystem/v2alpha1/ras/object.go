package ras

import "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

type Ras struct {
	Image            string
	ImagePullPolicy  string
	Storage          string
	StorageClassName string
	NodeSelector     map[string]string
}

func New(rasCrd *v2alpha1.GalasaRasComponent) *Ras {
	return &Ras{}
}

func (c *Ras) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *Ras) IsReady() bool {
	return true
}
