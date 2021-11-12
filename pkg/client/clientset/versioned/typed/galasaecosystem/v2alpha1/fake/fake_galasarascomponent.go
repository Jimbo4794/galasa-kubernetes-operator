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

// FakeGalasaRasComponents implements GalasaRasComponentInterface
type FakeGalasaRasComponents struct {
	Fake *FakeGalasaV2alpha1
	ns   string
}

var galasarascomponentsResource = schema.GroupVersionResource{Group: "galasa.dev", Version: "v2alpha1", Resource: "galasarascomponents"}

var galasarascomponentsKind = schema.GroupVersionKind{Group: "galasa.dev", Version: "v2alpha1", Kind: "GalasaRasComponent"}

// Get takes name of the galasaRasComponent, and returns the corresponding galasaRasComponent object, and an error if there is any.
func (c *FakeGalasaRasComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.GalasaRasComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(galasarascomponentsResource, c.ns, name), &v2alpha1.GalasaRasComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaRasComponent), err
}

// List takes label and field selectors, and returns the list of GalasaRasComponents that match those selectors.
func (c *FakeGalasaRasComponents) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.GalasaRasComponentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(galasarascomponentsResource, galasarascomponentsKind, c.ns, opts), &v2alpha1.GalasaRasComponentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.GalasaRasComponentList{ListMeta: obj.(*v2alpha1.GalasaRasComponentList).ListMeta}
	for _, item := range obj.(*v2alpha1.GalasaRasComponentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested galasaRasComponents.
func (c *FakeGalasaRasComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(galasarascomponentsResource, c.ns, opts))

}

// Create takes the representation of a galasaRasComponent and creates it.  Returns the server's representation of the galasaRasComponent, and an error, if there is any.
func (c *FakeGalasaRasComponents) Create(ctx context.Context, galasaRasComponent *v2alpha1.GalasaRasComponent, opts v1.CreateOptions) (result *v2alpha1.GalasaRasComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(galasarascomponentsResource, c.ns, galasaRasComponent), &v2alpha1.GalasaRasComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaRasComponent), err
}

// Update takes the representation of a galasaRasComponent and updates it. Returns the server's representation of the galasaRasComponent, and an error, if there is any.
func (c *FakeGalasaRasComponents) Update(ctx context.Context, galasaRasComponent *v2alpha1.GalasaRasComponent, opts v1.UpdateOptions) (result *v2alpha1.GalasaRasComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(galasarascomponentsResource, c.ns, galasaRasComponent), &v2alpha1.GalasaRasComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaRasComponent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGalasaRasComponents) UpdateStatus(ctx context.Context, galasaRasComponent *v2alpha1.GalasaRasComponent, opts v1.UpdateOptions) (*v2alpha1.GalasaRasComponent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(galasarascomponentsResource, "status", c.ns, galasaRasComponent), &v2alpha1.GalasaRasComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaRasComponent), err
}

// Delete takes name of the galasaRasComponent and deletes it. Returns an error if one occurs.
func (c *FakeGalasaRasComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(galasarascomponentsResource, c.ns, name), &v2alpha1.GalasaRasComponent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGalasaRasComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(galasarascomponentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.GalasaRasComponentList{})
	return err
}

// Patch applies the patch and returns the patched galasaRasComponent.
func (c *FakeGalasaRasComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.GalasaRasComponent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(galasarascomponentsResource, c.ns, name, pt, data, subresources...), &v2alpha1.GalasaRasComponent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.GalasaRasComponent), err
}
