package resmon

import (
	"context"

	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	galasaecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Resmon struct {
	Ecosystemclient galasaecosystem.Interface
	Name            string
	Namespace       string

	Image            string
	ImagePullPolicy  string
	Storage          string
	StorageClassName string
	NodeSelector     map[string]string
	Owner            []v1.OwnerReference
	Status           v2alpha1.ComponentStatus
}

func New(resmonCrd *v2alpha1.GalasaResmonComponent, k galasaecosystem.Interface) *Resmon {
	return &Resmon{}
}

func (c *Resmon) HasChanged(spec v2alpha1.ComponentSpec) bool {
	return false
}
func (c *Resmon) IsReady(ctx context.Context) bool {
	return true
}
func (c *Resmon) GetObjects() []runtime.Object {
	return nil
}
