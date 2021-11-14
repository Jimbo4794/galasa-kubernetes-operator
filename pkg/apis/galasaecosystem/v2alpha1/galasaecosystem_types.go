package v2alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genreconciler:krshapedlogic=false
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GalasaEcosystem struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GalasaEcosystemSpec `json:"spec,omitempty"`

	// +optional
	Status GalasaEcosystemStatus `json:"status,omitempty"`
}

type GalasaEcosystemSpec struct {
	Hostname      string `json:"hostname"`
	GalasaVersion string `json:"galasaVersion"`
	// +optional
	BusyboxImage   string                   `json:"busyboxImage"`
	ComponentsSpec map[string]ComponentSpec `json:"componentsSpec"`
}

type ComponentSpec struct {
	Image string `json:"image"`
	// +optional
	ImagePullPolicy string `json:"imagePullPolicy"`
	// +optional
	Storage string `json:"storage"`
	// +optional
	StorageClassName string `json:"storageClassName"`
	// +optional
	NodeSelector map[string]string `json:"nodeSelector"`
}

type ComponentInterface interface {
	IsReady() bool
	HasChanged(spec ComponentSpec) bool
}

type GalasaEcosystemStatus struct {
	Ready        bool   `json:"ready"`
	BootstrapURL string `json:"bootstrapURL"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GalasaEcosystemList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GalasaEcosystem `json:"items"`
}
