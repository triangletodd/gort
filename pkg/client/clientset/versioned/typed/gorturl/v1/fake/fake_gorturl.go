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
	gorturlv1 "github.com/triangletodd/gort/pkg/apis/gorturl/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGortURLs implements GortURLInterface
type FakeGortURLs struct {
	Fake *FakeMtnV1
	ns   string
}

var gorturlsResource = schema.GroupVersionResource{Group: "mtn.cc", Version: "v1", Resource: "gorturls"}

var gorturlsKind = schema.GroupVersionKind{Group: "mtn.cc", Version: "v1", Kind: "GortURL"}

// Get takes name of the gortURL, and returns the corresponding gortURL object, and an error if there is any.
func (c *FakeGortURLs) Get(name string, options v1.GetOptions) (result *gorturlv1.GortURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gorturlsResource, c.ns, name), &gorturlv1.GortURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gorturlv1.GortURL), err
}

// List takes label and field selectors, and returns the list of GortURLs that match those selectors.
func (c *FakeGortURLs) List(opts v1.ListOptions) (result *gorturlv1.GortURLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gorturlsResource, gorturlsKind, c.ns, opts), &gorturlv1.GortURLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gorturlv1.GortURLList{ListMeta: obj.(*gorturlv1.GortURLList).ListMeta}
	for _, item := range obj.(*gorturlv1.GortURLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gortURLs.
func (c *FakeGortURLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gorturlsResource, c.ns, opts))

}

// Create takes the representation of a gortURL and creates it.  Returns the server's representation of the gortURL, and an error, if there is any.
func (c *FakeGortURLs) Create(gortURL *gorturlv1.GortURL) (result *gorturlv1.GortURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gorturlsResource, c.ns, gortURL), &gorturlv1.GortURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gorturlv1.GortURL), err
}

// Update takes the representation of a gortURL and updates it. Returns the server's representation of the gortURL, and an error, if there is any.
func (c *FakeGortURLs) Update(gortURL *gorturlv1.GortURL) (result *gorturlv1.GortURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gorturlsResource, c.ns, gortURL), &gorturlv1.GortURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gorturlv1.GortURL), err
}

// Delete takes name of the gortURL and deletes it. Returns an error if one occurs.
func (c *FakeGortURLs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gorturlsResource, c.ns, name), &gorturlv1.GortURL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGortURLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gorturlsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &gorturlv1.GortURLList{})
	return err
}

// Patch applies the patch and returns the patched gortURL.
func (c *FakeGortURLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *gorturlv1.GortURL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gorturlsResource, c.ns, name, pt, data, subresources...), &gorturlv1.GortURL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gorturlv1.GortURL), err
}
