/*
Copyright The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	kops "k8s.io/kops/pkg/apis/kops"
)

// FakeKeysets implements KeysetInterface
type FakeKeysets struct {
	Fake *FakeKops
	ns   string
}

var keysetsResource = kops.SchemeGroupVersion.WithResource("keysets")

var keysetsKind = kops.SchemeGroupVersion.WithKind("Keyset")

// Get takes name of the keyset, and returns the corresponding keyset object, and an error if there is any.
func (c *FakeKeysets) Get(ctx context.Context, name string, options v1.GetOptions) (result *kops.Keyset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keysetsResource, c.ns, name), &kops.Keyset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.Keyset), err
}

// List takes label and field selectors, and returns the list of Keysets that match those selectors.
func (c *FakeKeysets) List(ctx context.Context, opts v1.ListOptions) (result *kops.KeysetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keysetsResource, keysetsKind, c.ns, opts), &kops.KeysetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kops.KeysetList{ListMeta: obj.(*kops.KeysetList).ListMeta}
	for _, item := range obj.(*kops.KeysetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keysets.
func (c *FakeKeysets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keysetsResource, c.ns, opts))

}

// Create takes the representation of a keyset and creates it.  Returns the server's representation of the keyset, and an error, if there is any.
func (c *FakeKeysets) Create(ctx context.Context, keyset *kops.Keyset, opts v1.CreateOptions) (result *kops.Keyset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keysetsResource, c.ns, keyset), &kops.Keyset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.Keyset), err
}

// Update takes the representation of a keyset and updates it. Returns the server's representation of the keyset, and an error, if there is any.
func (c *FakeKeysets) Update(ctx context.Context, keyset *kops.Keyset, opts v1.UpdateOptions) (result *kops.Keyset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keysetsResource, c.ns, keyset), &kops.Keyset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.Keyset), err
}

// Delete takes name of the keyset and deletes it. Returns an error if one occurs.
func (c *FakeKeysets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(keysetsResource, c.ns, name, opts), &kops.Keyset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeysets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keysetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kops.KeysetList{})
	return err
}

// Patch applies the patch and returns the patched keyset.
func (c *FakeKeysets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kops.Keyset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keysetsResource, c.ns, name, pt, data, subresources...), &kops.Keyset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kops.Keyset), err
}
