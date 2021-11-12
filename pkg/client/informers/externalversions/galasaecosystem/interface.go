/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by informer-gen. DO NOT EDIT.

package galasaecosystem

import (
	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/informers/externalversions/galasaecosystem/v2alpha1"
	internalinterfaces "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V2alpha1 provides access to shared informers for resources in V2alpha1.
	V2alpha1() v2alpha1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V2alpha1 returns a new v2alpha1.Interface.
func (g *group) V2alpha1() v2alpha1.Interface {
	return v2alpha1.New(g.factory, g.namespace, g.tweakListOptions)
}