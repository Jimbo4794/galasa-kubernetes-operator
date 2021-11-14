package api

import "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

type Api struct {
	Image            string
	ImagePullPolicy  string
	Storage          string
	StorageClassName string
	NodeSelector     map[string]string
}

func New(apiCrd *v2alpha1.GalasaApiComponent) *Api {
	return &Api{}
}

func (c *Api) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *Api) IsReady() bool {
	return true
}
