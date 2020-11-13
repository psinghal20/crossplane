/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/crossplane/crossplane/apis/apiextensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCompositions implements CompositionInterface
type FakeCompositions struct {
	Fake *FakeApiextensionsV1beta1
}

var compositionsResource = schema.GroupVersionResource{Group: "apiextensions.crossplane.io", Version: "v1beta1", Resource: "compositions"}

var compositionsKind = schema.GroupVersionKind{Group: "apiextensions.crossplane.io", Version: "v1beta1", Kind: "Composition"}

// Get takes name of the composition, and returns the corresponding composition object, and an error if there is any.
func (c *FakeCompositions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Composition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(compositionsResource, name), &v1beta1.Composition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Composition), err
}

// List takes label and field selectors, and returns the list of Compositions that match those selectors.
func (c *FakeCompositions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CompositionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(compositionsResource, compositionsKind, opts), &v1beta1.CompositionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CompositionList{ListMeta: obj.(*v1beta1.CompositionList).ListMeta}
	for _, item := range obj.(*v1beta1.CompositionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested compositions.
func (c *FakeCompositions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(compositionsResource, opts))
}

// Create takes the representation of a composition and creates it.  Returns the server's representation of the composition, and an error, if there is any.
func (c *FakeCompositions) Create(ctx context.Context, composition *v1beta1.Composition, opts v1.CreateOptions) (result *v1beta1.Composition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(compositionsResource, composition), &v1beta1.Composition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Composition), err
}

// Update takes the representation of a composition and updates it. Returns the server's representation of the composition, and an error, if there is any.
func (c *FakeCompositions) Update(ctx context.Context, composition *v1beta1.Composition, opts v1.UpdateOptions) (result *v1beta1.Composition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(compositionsResource, composition), &v1beta1.Composition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Composition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCompositions) UpdateStatus(ctx context.Context, composition *v1beta1.Composition, opts v1.UpdateOptions) (*v1beta1.Composition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(compositionsResource, "status", composition), &v1beta1.Composition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Composition), err
}

// Delete takes name of the composition and deletes it. Returns an error if one occurs.
func (c *FakeCompositions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(compositionsResource, name), &v1beta1.Composition{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCompositions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(compositionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CompositionList{})
	return err
}

// Patch applies the patch and returns the patched composition.
func (c *FakeCompositions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Composition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(compositionsResource, name, pt, data, subresources...), &v1beta1.Composition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Composition), err
}