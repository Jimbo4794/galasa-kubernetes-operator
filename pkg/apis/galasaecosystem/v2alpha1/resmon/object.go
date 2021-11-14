package resmon

import "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

type Resmon struct {
	Image           string
	ImagePullPolicy string
	NodeSelector    map[string]string
}

func New(resmonCrd *v2alpha1.GalasaResmonComponent) *Resmon {
	return &Resmon{}
}

func (c *Resmon) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *Resmon) IsReady() bool {
	return true
}
