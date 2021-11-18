/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by injection-gen. DO NOT EDIT.

package client

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"

	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	versioned "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	typedgalasav2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned/typed/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	discovery "k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/dynamic"
	rest "k8s.io/client-go/rest"
	injection "knative.dev/pkg/injection"
	dynamicclient "knative.dev/pkg/injection/clients/dynamicclient"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClientFromConfig)
	injection.Default.RegisterClientFetcher(func(ctx context.Context) interface{} {
		return Get(ctx)
	})
	injection.Dynamic.RegisterDynamicClient(withClientFromDynamic)
}

// Key is used as the key for associating information with a context.Context.
type Key struct{}

func withClientFromConfig(ctx context.Context, cfg *rest.Config) context.Context {
	return context.WithValue(ctx, Key{}, versioned.NewForConfigOrDie(cfg))
}

func withClientFromDynamic(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key{}, &wrapClient{dyn: dynamicclient.Get(ctx)})
}

// Get extracts the versioned.Interface client from the context.
func Get(ctx context.Context) versioned.Interface {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		if injection.GetConfig(ctx) == nil {
			logging.FromContext(ctx).Panic(
				"Unable to fetch github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned.Interface from context. This context is not the application context (which is typically given to constructors via sharedmain).")
		} else {
			logging.FromContext(ctx).Panic(
				"Unable to fetch github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned.Interface from context.")
		}
	}
	return untyped.(versioned.Interface)
}

type wrapClient struct {
	dyn dynamic.Interface
}

var _ versioned.Interface = (*wrapClient)(nil)

func (w *wrapClient) Discovery() discovery.DiscoveryInterface {
	panic("Discovery called on dynamic client!")
}

func convert(from interface{}, to runtime.Object) error {
	bs, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("Marshal() = %w", err)
	}
	if err := json.Unmarshal(bs, to); err != nil {
		return fmt.Errorf("Unmarshal() = %w", err)
	}
	return nil
}

// GalasaV2alpha1 retrieves the GalasaV2alpha1Client
func (w *wrapClient) GalasaV2alpha1() typedgalasav2alpha1.GalasaV2alpha1Interface {
	return &wrapGalasaV2alpha1{
		dyn: w.dyn,
	}
}

type wrapGalasaV2alpha1 struct {
	dyn dynamic.Interface
}

func (w *wrapGalasaV2alpha1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapGalasaV2alpha1) GalasaApiComponents(namespace string) typedgalasav2alpha1.GalasaApiComponentInterface {
	return &wrapGalasaV2alpha1GalasaApiComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasaapicomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaApiComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaApiComponentInterface = (*wrapGalasaV2alpha1GalasaApiComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaApiComponent, opts v1.CreateOptions) (*v2alpha1.GalasaApiComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaApiComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaApiComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaApiComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaApiComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaApiComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaApiComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaApiComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaApiComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaApiComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaApiComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaApiComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaApiComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaApiComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaApiComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaApiComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaApiComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaApiComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaCpsComponents(namespace string) typedgalasav2alpha1.GalasaCpsComponentInterface {
	return &wrapGalasaV2alpha1GalasaCpsComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasacpscomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaCpsComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaCpsComponentInterface = (*wrapGalasaV2alpha1GalasaCpsComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaCpsComponent, opts v1.CreateOptions) (*v2alpha1.GalasaCpsComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaCpsComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaCpsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaCpsComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaCpsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaCpsComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaCpsComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaCpsComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaCpsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaCpsComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaCpsComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaCpsComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaCpsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaCpsComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaCpsComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaCpsComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaCpsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaCpsComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaEcosystems(namespace string) typedgalasav2alpha1.GalasaEcosystemInterface {
	return &wrapGalasaV2alpha1GalasaEcosystemImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasaecosystems",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaEcosystemImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaEcosystemInterface = (*wrapGalasaV2alpha1GalasaEcosystemImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) Create(ctx context.Context, in *v2alpha1.GalasaEcosystem, opts v1.CreateOptions) (*v2alpha1.GalasaEcosystem, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaEcosystem",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEcosystem{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaEcosystem, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEcosystem{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaEcosystemList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEcosystemList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaEcosystem, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEcosystem{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) Update(ctx context.Context, in *v2alpha1.GalasaEcosystem, opts v1.UpdateOptions) (*v2alpha1.GalasaEcosystem, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaEcosystem",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEcosystem{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaEcosystem, opts v1.UpdateOptions) (*v2alpha1.GalasaEcosystem, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaEcosystem",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEcosystem{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEcosystemImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaEngineControllerComponents(namespace string) typedgalasav2alpha1.GalasaEngineControllerComponentInterface {
	return &wrapGalasaV2alpha1GalasaEngineControllerComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasaenginecontrollercomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaEngineControllerComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaEngineControllerComponentInterface = (*wrapGalasaV2alpha1GalasaEngineControllerComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaEngineControllerComponent, opts v1.CreateOptions) (*v2alpha1.GalasaEngineControllerComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaEngineControllerComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEngineControllerComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaEngineControllerComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEngineControllerComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaEngineControllerComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEngineControllerComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaEngineControllerComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEngineControllerComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaEngineControllerComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaEngineControllerComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaEngineControllerComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEngineControllerComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaEngineControllerComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaEngineControllerComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaEngineControllerComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaEngineControllerComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaEngineControllerComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaMetricsComponents(namespace string) typedgalasav2alpha1.GalasaMetricsComponentInterface {
	return &wrapGalasaV2alpha1GalasaMetricsComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasametricscomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaMetricsComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaMetricsComponentInterface = (*wrapGalasaV2alpha1GalasaMetricsComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaMetricsComponent, opts v1.CreateOptions) (*v2alpha1.GalasaMetricsComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaMetricsComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaMetricsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaMetricsComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaMetricsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaMetricsComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaMetricsComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaMetricsComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaMetricsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaMetricsComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaMetricsComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaMetricsComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaMetricsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaMetricsComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaMetricsComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaMetricsComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaMetricsComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaMetricsComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaRasComponents(namespace string) typedgalasav2alpha1.GalasaRasComponentInterface {
	return &wrapGalasaV2alpha1GalasaRasComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasarascomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaRasComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaRasComponentInterface = (*wrapGalasaV2alpha1GalasaRasComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaRasComponent, opts v1.CreateOptions) (*v2alpha1.GalasaRasComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaRasComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaRasComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaRasComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaRasComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaRasComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaRasComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaRasComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaRasComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaRasComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaRasComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaRasComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaRasComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaRasComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaRasComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaRasComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaRasComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaRasComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaResmonComponents(namespace string) typedgalasav2alpha1.GalasaResmonComponentInterface {
	return &wrapGalasaV2alpha1GalasaResmonComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasaresmoncomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaResmonComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaResmonComponentInterface = (*wrapGalasaV2alpha1GalasaResmonComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaResmonComponent, opts v1.CreateOptions) (*v2alpha1.GalasaResmonComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaResmonComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaResmonComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaResmonComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaResmonComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaResmonComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaResmonComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaResmonComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaResmonComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaResmonComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaResmonComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaResmonComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaResmonComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaResmonComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaResmonComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaResmonComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaResmonComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaResmonComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapGalasaV2alpha1) GalasaToolboxComponents(namespace string) typedgalasav2alpha1.GalasaToolboxComponentInterface {
	return &wrapGalasaV2alpha1GalasaToolboxComponentImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "galasa.dev",
			Version:  "v2alpha1",
			Resource: "galasatoolboxcomponents",
		}),

		namespace: namespace,
	}
}

type wrapGalasaV2alpha1GalasaToolboxComponentImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedgalasav2alpha1.GalasaToolboxComponentInterface = (*wrapGalasaV2alpha1GalasaToolboxComponentImpl)(nil)

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) Create(ctx context.Context, in *v2alpha1.GalasaToolboxComponent, opts v1.CreateOptions) (*v2alpha1.GalasaToolboxComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaToolboxComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaToolboxComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaToolboxComponent, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaToolboxComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaToolboxComponentList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaToolboxComponentList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaToolboxComponent, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaToolboxComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) Update(ctx context.Context, in *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaToolboxComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaToolboxComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaToolboxComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) UpdateStatus(ctx context.Context, in *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaToolboxComponent, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "galasa.dev",
		Version: "v2alpha1",
		Kind:    "GalasaToolboxComponent",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v2alpha1.GalasaToolboxComponent{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapGalasaV2alpha1GalasaToolboxComponentImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}
