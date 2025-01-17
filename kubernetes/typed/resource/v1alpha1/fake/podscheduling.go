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

	resourcev1alpha1 "k8s.io/api/resource/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsresourcev1alpha1 "k8s.io/client-go/applyconfigurations/resource/v1alpha1"
	resourcev1alpha1client "k8s.io/client-go/kubernetes/typed/resource/v1alpha1"
	"k8s.io/client-go/testing"

	kcpresourcev1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/resource/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var podSchedulingsResource = schema.GroupVersionResource{Group: "resource.k8s.io", Version: "v1alpha1", Resource: "podschedulings"}
var podSchedulingsKind = schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1alpha1", Kind: "PodScheduling"}

type podSchedulingsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *podSchedulingsClusterClient) Cluster(clusterPath logicalcluster.Path) kcpresourcev1alpha1.PodSchedulingsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podSchedulingsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PodSchedulings that match those selectors across all clusters.
func (c *podSchedulingsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha1.PodSchedulingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podSchedulingsResource, podSchedulingsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &resourcev1alpha1.PodSchedulingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha1.PodSchedulingList{ListMeta: obj.(*resourcev1alpha1.PodSchedulingList).ListMeta}
	for _, item := range obj.(*resourcev1alpha1.PodSchedulingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PodSchedulings across all clusters.
func (c *podSchedulingsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podSchedulingsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type podSchedulingsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *podSchedulingsNamespacer) Namespace(namespace string) resourcev1alpha1client.PodSchedulingInterface {
	return &podSchedulingsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type podSchedulingsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *podSchedulingsClient) Create(ctx context.Context, podScheduling *resourcev1alpha1.PodScheduling, opts metav1.CreateOptions) (*resourcev1alpha1.PodScheduling, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(podSchedulingsResource, c.ClusterPath, c.Namespace, podScheduling), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}

func (c *podSchedulingsClient) Update(ctx context.Context, podScheduling *resourcev1alpha1.PodScheduling, opts metav1.UpdateOptions) (*resourcev1alpha1.PodScheduling, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(podSchedulingsResource, c.ClusterPath, c.Namespace, podScheduling), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}

func (c *podSchedulingsClient) UpdateStatus(ctx context.Context, podScheduling *resourcev1alpha1.PodScheduling, opts metav1.UpdateOptions) (*resourcev1alpha1.PodScheduling, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(podSchedulingsResource, c.ClusterPath, "status", c.Namespace, podScheduling), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}

func (c *podSchedulingsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(podSchedulingsResource, c.ClusterPath, c.Namespace, name, opts), &resourcev1alpha1.PodScheduling{})
	return err
}

func (c *podSchedulingsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(podSchedulingsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &resourcev1alpha1.PodSchedulingList{})
	return err
}

func (c *podSchedulingsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*resourcev1alpha1.PodScheduling, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(podSchedulingsResource, c.ClusterPath, c.Namespace, name), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}

// List takes label and field selectors, and returns the list of PodSchedulings that match those selectors.
func (c *podSchedulingsClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha1.PodSchedulingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podSchedulingsResource, podSchedulingsKind, c.ClusterPath, c.Namespace, opts), &resourcev1alpha1.PodSchedulingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha1.PodSchedulingList{ListMeta: obj.(*resourcev1alpha1.PodSchedulingList).ListMeta}
	for _, item := range obj.(*resourcev1alpha1.PodSchedulingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *podSchedulingsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podSchedulingsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *podSchedulingsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*resourcev1alpha1.PodScheduling, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podSchedulingsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}

func (c *podSchedulingsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha1.PodSchedulingApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha1.PodScheduling, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podSchedulingsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}

func (c *podSchedulingsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha1.PodSchedulingApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha1.PodScheduling, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podSchedulingsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &resourcev1alpha1.PodScheduling{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha1.PodScheduling), err
}
