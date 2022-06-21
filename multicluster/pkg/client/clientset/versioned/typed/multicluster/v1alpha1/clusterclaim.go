// Copyright 2021 Antrea Authors
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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	scheme "antrea.io/antrea/multicluster/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterClaimsGetter has a method to return a ClusterClaimInterface.
// A group's client should implement this interface.
type ClusterClaimsGetter interface {
	ClusterClaims(namespace string) ClusterClaimInterface
}

// ClusterClaimInterface has methods to work with ClusterClaim resources.
type ClusterClaimInterface interface {
	Create(ctx context.Context, clusterClaim *v1alpha1.ClusterClaim, opts v1.CreateOptions) (*v1alpha1.ClusterClaim, error)
	Update(ctx context.Context, clusterClaim *v1alpha1.ClusterClaim, opts v1.UpdateOptions) (*v1alpha1.ClusterClaim, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterClaim, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterClaimList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterClaim, err error)
	ClusterClaimExpansion
}

// clusterClaims implements ClusterClaimInterface
type clusterClaims struct {
	client rest.Interface
	ns     string
}

// newClusterClaims returns a ClusterClaims
func newClusterClaims(c *MulticlusterV1alpha1Client, namespace string) *clusterClaims {
	return &clusterClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterClaim, and returns the corresponding clusterClaim object, and an error if there is any.
func (c *clusterClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterClaim, err error) {
	result = &v1alpha1.ClusterClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterClaims that match those selectors.
func (c *clusterClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterClaims.
func (c *clusterClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterClaim and creates it.  Returns the server's representation of the clusterClaim, and an error, if there is any.
func (c *clusterClaims) Create(ctx context.Context, clusterClaim *v1alpha1.ClusterClaim, opts v1.CreateOptions) (result *v1alpha1.ClusterClaim, err error) {
	result = &v1alpha1.ClusterClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterClaim and updates it. Returns the server's representation of the clusterClaim, and an error, if there is any.
func (c *clusterClaims) Update(ctx context.Context, clusterClaim *v1alpha1.ClusterClaim, opts v1.UpdateOptions) (result *v1alpha1.ClusterClaim, err error) {
	result = &v1alpha1.ClusterClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterclaims").
		Name(clusterClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterClaim and deletes it. Returns an error if one occurs.
func (c *clusterClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterClaim.
func (c *clusterClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterClaim, err error) {
	result = &v1alpha1.ClusterClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
