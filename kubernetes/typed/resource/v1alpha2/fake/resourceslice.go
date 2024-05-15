//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha2 "k8s.io/api/resource/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsresourcev1alpha2 "k8s.io/client-go/applyconfigurations/resource/v1alpha2"
	resourcev1alpha2client "k8s.io/client-go/kubernetes/typed/resource/v1alpha2"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var resourceSlicesResource = schema.GroupVersionResource{Group: "resource.k8s.io", Version: "v1alpha2", Resource: "resourceslices"}
var resourceSlicesKind = schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1alpha2", Kind: "ResourceSlice"}

type resourceSlicesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *resourceSlicesClusterClient) Cluster(clusterPath logicalcluster.Path) resourcev1alpha2client.ResourceSliceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &resourceSlicesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ResourceSlices that match those selectors across all clusters.
func (c *resourceSlicesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha2.ResourceSliceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(resourceSlicesResource, resourceSlicesKind, logicalcluster.Wildcard, opts), &resourcev1alpha2.ResourceSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha2.ResourceSliceList{ListMeta: obj.(*resourcev1alpha2.ResourceSliceList).ListMeta}
	for _, item := range obj.(*resourcev1alpha2.ResourceSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ResourceSlices across all clusters.
func (c *resourceSlicesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(resourceSlicesResource, logicalcluster.Wildcard, opts))
}

type resourceSlicesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *resourceSlicesClient) Create(ctx context.Context, resourceSlice *resourcev1alpha2.ResourceSlice, opts metav1.CreateOptions) (*resourcev1alpha2.ResourceSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(resourceSlicesResource, c.ClusterPath, resourceSlice), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}

func (c *resourceSlicesClient) Update(ctx context.Context, resourceSlice *resourcev1alpha2.ResourceSlice, opts metav1.UpdateOptions) (*resourcev1alpha2.ResourceSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(resourceSlicesResource, c.ClusterPath, resourceSlice), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}

func (c *resourceSlicesClient) UpdateStatus(ctx context.Context, resourceSlice *resourcev1alpha2.ResourceSlice, opts metav1.UpdateOptions) (*resourcev1alpha2.ResourceSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(resourceSlicesResource, c.ClusterPath, "status", resourceSlice), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}

func (c *resourceSlicesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(resourceSlicesResource, c.ClusterPath, name, opts), &resourcev1alpha2.ResourceSlice{})
	return err
}

func (c *resourceSlicesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(resourceSlicesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &resourcev1alpha2.ResourceSliceList{})
	return err
}

func (c *resourceSlicesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*resourcev1alpha2.ResourceSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(resourceSlicesResource, c.ClusterPath, name), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}

// List takes label and field selectors, and returns the list of ResourceSlices that match those selectors.
func (c *resourceSlicesClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha2.ResourceSliceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(resourceSlicesResource, resourceSlicesKind, c.ClusterPath, opts), &resourcev1alpha2.ResourceSliceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha2.ResourceSliceList{ListMeta: obj.(*resourcev1alpha2.ResourceSliceList).ListMeta}
	for _, item := range obj.(*resourcev1alpha2.ResourceSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *resourceSlicesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(resourceSlicesResource, c.ClusterPath, opts))
}

func (c *resourceSlicesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*resourcev1alpha2.ResourceSlice, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(resourceSlicesResource, c.ClusterPath, name, pt, data, subresources...), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}

func (c *resourceSlicesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha2.ResourceSliceApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha2.ResourceSlice, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(resourceSlicesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}

func (c *resourceSlicesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha2.ResourceSliceApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha2.ResourceSlice, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(resourceSlicesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &resourcev1alpha2.ResourceSlice{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceSlice), err
}