/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by injection-gen. DO NOT EDIT.

package galasaresmoncomponent

import (
	context "context"

	apisgalasaecosystemv2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	versioned "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/informers/externalversions/galasaecosystem/v2alpha1"
	client "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/injection/client"
	factory "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/injection/informers/factory"
	galasaecosystemv2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/listers/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Galasa().V2alpha1().GalasaResmonComponents()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

func withDynamicInformer(ctx context.Context) context.Context {
	inf := &wrapper{client: client.Get(ctx), resourceVersion: injection.GetResourceVersion(ctx)}
	return context.WithValue(ctx, Key{}, inf)
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v2alpha1.GalasaResmonComponentInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/informers/externalversions/galasaecosystem/v2alpha1.GalasaResmonComponentInformer from context.")
	}
	return untyped.(v2alpha1.GalasaResmonComponentInformer)
}

type wrapper struct {
	client versioned.Interface

	namespace string

	resourceVersion string
}

var _ v2alpha1.GalasaResmonComponentInformer = (*wrapper)(nil)
var _ galasaecosystemv2alpha1.GalasaResmonComponentLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apisgalasaecosystemv2alpha1.GalasaResmonComponent{}, 0, nil)
}

func (w *wrapper) Lister() galasaecosystemv2alpha1.GalasaResmonComponentLister {
	return w
}

func (w *wrapper) GalasaResmonComponents(namespace string) galasaecosystemv2alpha1.GalasaResmonComponentNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace, resourceVersion: w.resourceVersion}
}

// SetResourceVersion allows consumers to adjust the minimum resourceVersion
// used by the underlying client.  It is not accessible via the standard
// lister interface, but can be accessed through a user-defined interface and
// an implementation check e.g. rvs, ok := foo.(ResourceVersionSetter)
func (w *wrapper) SetResourceVersion(resourceVersion string) {
	w.resourceVersion = resourceVersion
}

func (w *wrapper) List(selector labels.Selector) (ret []*apisgalasaecosystemv2alpha1.GalasaResmonComponent, err error) {
	lo, err := w.client.GalasaV2alpha1().GalasaResmonComponents(w.namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector:   selector.String(),
		ResourceVersion: w.resourceVersion,
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apisgalasaecosystemv2alpha1.GalasaResmonComponent, error) {
	return w.client.GalasaV2alpha1().GalasaResmonComponents(w.namespace).Get(context.TODO(), name, v1.GetOptions{
		ResourceVersion: w.resourceVersion,
	})
}
