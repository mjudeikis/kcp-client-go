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

package v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsrbacv1alpha1 "k8s.io/client-go/applyconfigurations/rbac/v1alpha1"
	rbacv1alpha1client "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var clusterRoleBindingsResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1alpha1", Resource: "clusterrolebindings"}
var clusterRoleBindingsKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1alpha1", Kind: "ClusterRoleBinding"}

type clusterRoleBindingsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterRoleBindingsClusterClient) Cluster(cluster logicalcluster.Name) rbacv1alpha1client.ClusterRoleBindingInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterRoleBindingsClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ClusterRoleBindings that match those selectors across all clusters.
func (c *clusterRoleBindingsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1alpha1.ClusterRoleBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterRoleBindingsResource, clusterRoleBindingsKind, logicalcluster.Wildcard, opts), &rbacv1alpha1.ClusterRoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1alpha1.ClusterRoleBindingList{ListMeta: obj.(*rbacv1alpha1.ClusterRoleBindingList).ListMeta}
	for _, item := range obj.(*rbacv1alpha1.ClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterRoleBindings across all clusters.
func (c *clusterRoleBindingsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterRoleBindingsResource, logicalcluster.Wildcard, opts))
}

type clusterRoleBindingsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *clusterRoleBindingsClient) Create(ctx context.Context, clusterRoleBinding *rbacv1alpha1.ClusterRoleBinding, opts metav1.CreateOptions) (*rbacv1alpha1.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterRoleBindingsResource, c.Cluster, clusterRoleBinding), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}

func (c *clusterRoleBindingsClient) Update(ctx context.Context, clusterRoleBinding *rbacv1alpha1.ClusterRoleBinding, opts metav1.UpdateOptions) (*rbacv1alpha1.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterRoleBindingsResource, c.Cluster, clusterRoleBinding), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}

func (c *clusterRoleBindingsClient) UpdateStatus(ctx context.Context, clusterRoleBinding *rbacv1alpha1.ClusterRoleBinding, opts metav1.UpdateOptions) (*rbacv1alpha1.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterRoleBindingsResource, c.Cluster, "status", clusterRoleBinding), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}

func (c *clusterRoleBindingsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterRoleBindingsResource, c.Cluster, name, opts), &rbacv1alpha1.ClusterRoleBinding{})
	return err
}

func (c *clusterRoleBindingsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterRoleBindingsResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &rbacv1alpha1.ClusterRoleBindingList{})
	return err
}

func (c *clusterRoleBindingsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*rbacv1alpha1.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterRoleBindingsResource, c.Cluster, name), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}

// List takes label and field selectors, and returns the list of ClusterRoleBindings that match those selectors.
func (c *clusterRoleBindingsClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1alpha1.ClusterRoleBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterRoleBindingsResource, clusterRoleBindingsKind, c.Cluster, opts), &rbacv1alpha1.ClusterRoleBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1alpha1.ClusterRoleBindingList{ListMeta: obj.(*rbacv1alpha1.ClusterRoleBindingList).ListMeta}
	for _, item := range obj.(*rbacv1alpha1.ClusterRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterRoleBindingsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterRoleBindingsResource, c.Cluster, opts))
}

func (c *clusterRoleBindingsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1alpha1.ClusterRoleBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterRoleBindingsResource, c.Cluster, name, pt, data, subresources...), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}

func (c *clusterRoleBindingsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1alpha1.ClusterRoleBindingApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1alpha1.ClusterRoleBinding, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterRoleBindingsResource, c.Cluster, *name, types.ApplyPatchType, data), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}

func (c *clusterRoleBindingsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1alpha1.ClusterRoleBindingApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1alpha1.ClusterRoleBinding, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterRoleBindingsResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &rbacv1alpha1.ClusterRoleBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1alpha1.ClusterRoleBinding), err
}
