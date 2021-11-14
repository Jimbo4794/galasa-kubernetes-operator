package enginecontroller

import "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

type EngineController struct {
	Image           string
	ImagePullPolicy string
	NodeSelector    map[string]string
}

func New(engineControllerCrd *v2alpha1.GalasaEngineControllerComponent) *EngineController {
	return &EngineController{}
}

func (c *EngineController) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *EngineController) IsReady() bool {
	return true
}
