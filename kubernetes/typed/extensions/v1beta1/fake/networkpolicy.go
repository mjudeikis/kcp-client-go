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

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsextensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	extensionsv1beta1client "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	"k8s.io/client-go/testing"

	kcpextensionsv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/extensions/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var networkPoliciesResource = schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "networkpolicies"}
var networkPoliciesKind = schema.GroupVersionKind{Group: "extensions", Version: "v1beta1", Kind: "NetworkPolicy"}

type networkPoliciesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *networkPoliciesClusterClient) Cluster(cluster logicalcluster.Name) kcpextensionsv1beta1.NetworkPoliciesNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &networkPoliciesNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors across all clusters.
func (c *networkPoliciesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.NetworkPolicyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(networkPoliciesResource, networkPoliciesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &extensionsv1beta1.NetworkPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensionsv1beta1.NetworkPolicyList{ListMeta: obj.(*extensionsv1beta1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*extensionsv1beta1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested NetworkPolicies across all clusters.
func (c *networkPoliciesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(networkPoliciesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type networkPoliciesNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *networkPoliciesNamespacer) Namespace(namespace string) extensionsv1beta1client.NetworkPolicyInterface {
	return &networkPoliciesClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type networkPoliciesClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *networkPoliciesClient) Create(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicy, opts metav1.CreateOptions) (*extensionsv1beta1.NetworkPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(networkPoliciesResource, c.Cluster, c.Namespace, networkPolicy), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}

func (c *networkPoliciesClient) Update(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicy, opts metav1.UpdateOptions) (*extensionsv1beta1.NetworkPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(networkPoliciesResource, c.Cluster, c.Namespace, networkPolicy), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}

func (c *networkPoliciesClient) UpdateStatus(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicy, opts metav1.UpdateOptions) (*extensionsv1beta1.NetworkPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(networkPoliciesResource, c.Cluster, "status", c.Namespace, networkPolicy), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}

func (c *networkPoliciesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(networkPoliciesResource, c.Cluster, c.Namespace, name, opts), &extensionsv1beta1.NetworkPolicy{})
	return err
}

func (c *networkPoliciesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(networkPoliciesResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &extensionsv1beta1.NetworkPolicyList{})
	return err
}

func (c *networkPoliciesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*extensionsv1beta1.NetworkPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(networkPoliciesResource, c.Cluster, c.Namespace, name), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *networkPoliciesClient) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.NetworkPolicyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(networkPoliciesResource, networkPoliciesKind, c.Cluster, c.Namespace, opts), &extensionsv1beta1.NetworkPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensionsv1beta1.NetworkPolicyList{ListMeta: obj.(*extensionsv1beta1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*extensionsv1beta1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *networkPoliciesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(networkPoliciesResource, c.Cluster, c.Namespace, opts))
}

func (c *networkPoliciesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*extensionsv1beta1.NetworkPolicy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(networkPoliciesResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}

func (c *networkPoliciesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsextensionsv1beta1.NetworkPolicyApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.NetworkPolicy, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(networkPoliciesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}

func (c *networkPoliciesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsextensionsv1beta1.NetworkPolicyApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.NetworkPolicy, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(networkPoliciesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &extensionsv1beta1.NetworkPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.NetworkPolicy), err
}