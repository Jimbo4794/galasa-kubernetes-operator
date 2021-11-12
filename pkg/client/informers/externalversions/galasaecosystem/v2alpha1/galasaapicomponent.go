/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	time "time"

	galasaecosystemv2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	versioned "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/listers/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GalasaApiComponentInformer provides access to a shared informer and lister for
// GalasaApiComponents.
type GalasaApiComponentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.GalasaApiComponentLister
}

type galasaApiComponentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGalasaApiComponentInformer constructs a new informer for GalasaApiComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGalasaApiComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGalasaApiComponentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGalasaApiComponentInformer constructs a new informer for GalasaApiComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGalasaApiComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GalasaV2alpha1().GalasaApiComponents(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GalasaV2alpha1().GalasaApiComponents(namespace).Watch(context.TODO(), options)
			},
		},
		&galasaecosystemv2alpha1.GalasaApiComponent{},
		resyncPeriod,
		indexers,
	)
}

func (f *galasaApiComponentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGalasaApiComponentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *galasaApiComponentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&galasaecosystemv2alpha1.GalasaApiComponent{}, f.defaultInformer)
}

func (f *galasaApiComponentInformer) Lister() v2alpha1.GalasaApiComponentLister {
	return v2alpha1.NewGalasaApiComponentLister(f.Informer().GetIndexer())
}
