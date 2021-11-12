/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	"time"

	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	scheme "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GalasaToolboxComponentsGetter has a method to return a GalasaToolboxComponentInterface.
// A group's client should implement this interface.
type GalasaToolboxComponentsGetter interface {
	GalasaToolboxComponents(namespace string) GalasaToolboxComponentInterface
}

// GalasaToolboxComponentInterface has methods to work with GalasaToolboxComponent resources.
type GalasaToolboxComponentInterface interface {
	Create(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.CreateOptions) (*v2alpha1.GalasaToolboxComponent, error)
	Update(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaToolboxComponent, error)
	UpdateStatus(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaToolboxComponent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.GalasaToolboxComponent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.GalasaToolboxComponentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaToolboxComponent, err error)
	GalasaToolboxComponentExpansion
}

// galasaToolboxComponents implements GalasaToolboxComponentInterface
type galasaToolboxComponents struct {
	client rest.Interface
	ns     string
}

// newGalasaToolboxComponents returns a GalasaToolboxComponents
func newGalasaToolboxComponents(c *GalasaV2alpha1Client, namespace string) *galasaToolboxComponents {
	return &galasaToolboxComponents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the galasaToolboxComponent, and returns the corresponding galasaToolboxComponent object, and an error if there is any.
func (c *galasaToolboxComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	result = &v2alpha1.GalasaToolboxComponent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GalasaToolboxComponents that match those selectors.
func (c *galasaToolboxComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaToolboxComponentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.GalasaToolboxComponentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested galasaToolboxComponents.
func (c *galasaToolboxComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a galasaToolboxComponent and creates it.  Returns the server's representation of the galasaToolboxComponent, and an error, if there is any.
func (c *galasaToolboxComponents) Create(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	result = &v2alpha1.GalasaToolboxComponent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaToolboxComponent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a galasaToolboxComponent and updates it. Returns the server's representation of the galasaToolboxComponent, and an error, if there is any.
func (c *galasaToolboxComponents) Update(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	result = &v2alpha1.GalasaToolboxComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		Name(galasaToolboxComponent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaToolboxComponent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *galasaToolboxComponents) UpdateStatus(ctx context.Context, galasaToolboxComponent *v2alpha1.GalasaToolboxComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaToolboxComponent, err error) {
	result = &v2alpha1.GalasaToolboxComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		Name(galasaToolboxComponent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(galasaToolboxComponent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the galasaToolboxComponent and deletes it. Returns an error if one occurs.
func (c *galasaToolboxComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *galasaToolboxComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched galasaToolboxComponent.
func (c *galasaToolboxComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaToolboxComponent, err error) {
	result = &v2alpha1.GalasaToolboxComponent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("galasatoolboxcomponents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}