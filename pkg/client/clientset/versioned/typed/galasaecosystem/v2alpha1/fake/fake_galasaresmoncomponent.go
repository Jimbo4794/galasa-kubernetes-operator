/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGalasaResmonComponents implements GalasaResmonComponentInterface
type FakeGalasaResmonComponents struct {
	Fake *FakeGalasaV2alpha1
	ns   string
}

var galasaresmoncomponentsResource = schema.GroupVersionResource{Group: "galasa.dev", Version: "v2alpha1", Resource: "galasaresmoncomponents"}

var galasaresmoncomponentsKind = schema.GroupVersionKind{Group: "galasa.dev", Version: "v2alpha1", Kind: "GalasaResmonComponent"}

// Get takes name of the galasaResmonComponent, and returns the corresponding galasaResmonComponent object, and an error if there is any.
func (c *FakeGalasaResmonComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaResmonComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(galasaresmoncomponentsResource, c.ns, name), &v2alpha1.GalasaResmonComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaResmonComponent), err
}

// List takes label and field selectors, and returns the list of GalasaResmonComponents that match those selectors.
func (c *FakeGalasaResmonComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaResmonComponentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(galasaresmoncomponentsResource, galasaresmoncomponentsKind, c.ns, opts), &v2alpha1.GalasaResmonComponentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.GalasaResmonComponentList{ListMeta: obj.(*v2alpha1.GalasaResmonComponentList).ListMeta}
	for _, item := range obj.(*v2alpha1.GalasaResmonComponentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested galasaResmonComponents.
func (c *FakeGalasaResmonComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(galasaresmoncomponentsResource, c.ns, opts))

}

// Create takes the representation of a galasaResmonComponent and creates it.  Returns the server's representation of the galasaResmonComponent, and an error, if there is any.
func (c *FakeGalasaResmonComponents) Create(ctx context.Context, galasaResmonComponent *v2alpha1.GalasaResmonComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaResmonComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(galasaresmoncomponentsResource, c.ns, galasaResmonComponent), &v2alpha1.GalasaResmonComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaResmonComponent), err
}

// Update takes the representation of a galasaResmonComponent and updates it. Returns the server's representation of the galasaResmonComponent, and an error, if there is any.
func (c *FakeGalasaResmonComponents) Update(ctx context.Context, galasaResmonComponent *v2alpha1.GalasaResmonComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaResmonComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(galasaresmoncomponentsResource, c.ns, galasaResmonComponent), &v2alpha1.GalasaResmonComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaResmonComponent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGalasaResmonComponents) UpdateStatus(ctx context.Context, galasaResmonComponent *v2alpha1.GalasaResmonComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaResmonComponent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(galasaresmoncomponentsResource, "status", c.ns, galasaResmonComponent), &v2alpha1.GalasaResmonComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaResmonComponent), err
}

// Delete takes name of the galasaResmonComponent and deletes it. Returns an error if one occurs.
func (c *FakeGalasaResmonComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(galasaresmoncomponentsResource, c.ns, name), &v2alpha1.GalasaResmonComponent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGalasaResmonComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(galasaresmoncomponentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.GalasaResmonComponentList{})
	return err
}

// Patch applies the patch and returns the patched galasaResmonComponent.
func (c *FakeGalasaResmonComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaResmonComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(galasaresmoncomponentsResource, c.ns, name, pt, data, subresources...), &v2alpha1.GalasaResmonComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaResmonComponent), err
}