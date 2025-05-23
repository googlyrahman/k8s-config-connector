// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KMSKeyHandlesGetter has a method to return a KMSKeyHandleInterface.
// A group's client should implement this interface.
type KMSKeyHandlesGetter interface {
	KMSKeyHandles(namespace string) KMSKeyHandleInterface
}

// KMSKeyHandleInterface has methods to work with KMSKeyHandle resources.
type KMSKeyHandleInterface interface {
	Create(ctx context.Context, kMSKeyHandle *v1alpha1.KMSKeyHandle, opts v1.CreateOptions) (*v1alpha1.KMSKeyHandle, error)
	Update(ctx context.Context, kMSKeyHandle *v1alpha1.KMSKeyHandle, opts v1.UpdateOptions) (*v1alpha1.KMSKeyHandle, error)
	UpdateStatus(ctx context.Context, kMSKeyHandle *v1alpha1.KMSKeyHandle, opts v1.UpdateOptions) (*v1alpha1.KMSKeyHandle, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KMSKeyHandle, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KMSKeyHandleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KMSKeyHandle, err error)
	KMSKeyHandleExpansion
}

// kMSKeyHandles implements KMSKeyHandleInterface
type kMSKeyHandles struct {
	client rest.Interface
	ns     string
}

// newKMSKeyHandles returns a KMSKeyHandles
func newKMSKeyHandles(c *KmsV1alpha1Client, namespace string) *kMSKeyHandles {
	return &kMSKeyHandles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kMSKeyHandle, and returns the corresponding kMSKeyHandle object, and an error if there is any.
func (c *kMSKeyHandles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KMSKeyHandle, err error) {
	result = &v1alpha1.KMSKeyHandle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KMSKeyHandles that match those selectors.
func (c *kMSKeyHandles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KMSKeyHandleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KMSKeyHandleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kMSKeyHandles.
func (c *kMSKeyHandles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kMSKeyHandle and creates it.  Returns the server's representation of the kMSKeyHandle, and an error, if there is any.
func (c *kMSKeyHandles) Create(ctx context.Context, kMSKeyHandle *v1alpha1.KMSKeyHandle, opts v1.CreateOptions) (result *v1alpha1.KMSKeyHandle, err error) {
	result = &v1alpha1.KMSKeyHandle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kMSKeyHandle).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kMSKeyHandle and updates it. Returns the server's representation of the kMSKeyHandle, and an error, if there is any.
func (c *kMSKeyHandles) Update(ctx context.Context, kMSKeyHandle *v1alpha1.KMSKeyHandle, opts v1.UpdateOptions) (result *v1alpha1.KMSKeyHandle, err error) {
	result = &v1alpha1.KMSKeyHandle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		Name(kMSKeyHandle.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kMSKeyHandle).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kMSKeyHandles) UpdateStatus(ctx context.Context, kMSKeyHandle *v1alpha1.KMSKeyHandle, opts v1.UpdateOptions) (result *v1alpha1.KMSKeyHandle, err error) {
	result = &v1alpha1.KMSKeyHandle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		Name(kMSKeyHandle.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kMSKeyHandle).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kMSKeyHandle and deletes it. Returns an error if one occurs.
func (c *kMSKeyHandles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kMSKeyHandles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmskeyhandles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kMSKeyHandle.
func (c *kMSKeyHandles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KMSKeyHandle, err error) {
	result = &v1alpha1.KMSKeyHandle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmskeyhandles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
