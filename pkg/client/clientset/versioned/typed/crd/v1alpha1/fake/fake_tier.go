// Copyright 2022 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "antrea.io/antrea/pkg/apis/crd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTiers implements TierInterface
type FakeTiers struct {
	Fake *FakeCrdV1alpha1
}

var tiersResource = schema.GroupVersionResource{Group: "crd.antrea.io", Version: "v1alpha1", Resource: "tiers"}

var tiersKind = schema.GroupVersionKind{Group: "crd.antrea.io", Version: "v1alpha1", Kind: "Tier"}

// Get takes name of the tier, and returns the corresponding tier object, and an error if there is any.
func (c *FakeTiers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Tier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tiersResource, name), &v1alpha1.Tier{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tier), err
}

// List takes label and field selectors, and returns the list of Tiers that match those selectors.
func (c *FakeTiers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TierList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tiersResource, tiersKind, opts), &v1alpha1.TierList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TierList{ListMeta: obj.(*v1alpha1.TierList).ListMeta}
	for _, item := range obj.(*v1alpha1.TierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tiers.
func (c *FakeTiers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tiersResource, opts))
}

// Create takes the representation of a tier and creates it.  Returns the server's representation of the tier, and an error, if there is any.
func (c *FakeTiers) Create(ctx context.Context, tier *v1alpha1.Tier, opts v1.CreateOptions) (result *v1alpha1.Tier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tiersResource, tier), &v1alpha1.Tier{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tier), err
}

// Update takes the representation of a tier and updates it. Returns the server's representation of the tier, and an error, if there is any.
func (c *FakeTiers) Update(ctx context.Context, tier *v1alpha1.Tier, opts v1.UpdateOptions) (result *v1alpha1.Tier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tiersResource, tier), &v1alpha1.Tier{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tier), err
}

// Delete takes name of the tier and deletes it. Returns an error if one occurs.
func (c *FakeTiers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(tiersResource, name, opts), &v1alpha1.Tier{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTiers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tiersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TierList{})
	return err
}

// Patch applies the patch and returns the patched tier.
func (c *FakeTiers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Tier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tiersResource, name, pt, data, subresources...), &v1alpha1.Tier{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tier), err
}
