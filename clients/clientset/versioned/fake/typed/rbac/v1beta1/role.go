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

package v1beta1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsrbacv1beta1 "k8s.io/client-go/applyconfigurations/rbac/v1beta1"
	rbacv1beta1client "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	"k8s.io/client-go/testing"

	kcprbacv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/rbac/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var rolesResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1beta1", Resource: "roles"}
var rolesKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1beta1", Kind: "Role"}

type rolesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *rolesClusterClient) Cluster(cluster logicalcluster.Name) kcprbacv1beta1.RolesNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &rolesNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of Roles that match those selectors across all clusters.
func (c *rolesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1beta1.RoleList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(rolesResource, rolesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &rbacv1beta1.RoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1beta1.RoleList{ListMeta: obj.(*rbacv1beta1.RoleList).ListMeta}
	for _, item := range obj.(*rbacv1beta1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Roles across all clusters.
func (c *rolesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(rolesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type rolesNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *rolesNamespacer) Namespace(namespace string) rbacv1beta1client.RoleInterface {
	return &rolesClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type rolesClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *rolesClient) Create(ctx context.Context, role *rbacv1beta1.Role, opts metav1.CreateOptions) (*rbacv1beta1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(rolesResource, c.Cluster, c.Namespace, role), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}

func (c *rolesClient) Update(ctx context.Context, role *rbacv1beta1.Role, opts metav1.UpdateOptions) (*rbacv1beta1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(rolesResource, c.Cluster, c.Namespace, role), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}

func (c *rolesClient) UpdateStatus(ctx context.Context, role *rbacv1beta1.Role, opts metav1.UpdateOptions) (*rbacv1beta1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(rolesResource, c.Cluster, "status", c.Namespace, role), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}

func (c *rolesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(rolesResource, c.Cluster, c.Namespace, name, opts), &rbacv1beta1.Role{})
	return err
}

func (c *rolesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(rolesResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &rbacv1beta1.RoleList{})
	return err
}

func (c *rolesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*rbacv1beta1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(rolesResource, c.Cluster, c.Namespace, name), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}

// List takes label and field selectors, and returns the list of Roles that match those selectors.
func (c *rolesClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1beta1.RoleList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(rolesResource, rolesKind, c.Cluster, c.Namespace, opts), &rbacv1beta1.RoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1beta1.RoleList{ListMeta: obj.(*rbacv1beta1.RoleList).ListMeta}
	for _, item := range obj.(*rbacv1beta1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *rolesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(rolesResource, c.Cluster, c.Namespace, opts))
}

func (c *rolesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1beta1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolesResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}

func (c *rolesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1beta1.RoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1beta1.Role, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}

func (c *rolesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1beta1.RoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1beta1.Role, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &rbacv1beta1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.Role), err
}
