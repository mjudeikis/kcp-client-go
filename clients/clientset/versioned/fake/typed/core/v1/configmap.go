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

package v1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"

	kcpcorev1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var configMapsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"}
var configMapsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}

type configMapsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *configMapsClusterClient) Cluster(cluster logicalcluster.Name) kcpcorev1.ConfigMapsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &configMapsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ConfigMaps that match those selectors across all clusters.
func (c *configMapsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(configMapsResource, configMapsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &corev1.ConfigMapList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ConfigMapList{ListMeta: obj.(*corev1.ConfigMapList).ListMeta}
	for _, item := range obj.(*corev1.ConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ConfigMaps across all clusters.
func (c *configMapsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(configMapsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type configMapsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *configMapsNamespacer) Namespace(namespace string) corev1client.ConfigMapInterface {
	return &configMapsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type configMapsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *configMapsClient) Create(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.CreateOptions) (*corev1.ConfigMap, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(configMapsResource, c.Cluster, c.Namespace, configMap), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

func (c *configMapsClient) Update(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.UpdateOptions) (*corev1.ConfigMap, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(configMapsResource, c.Cluster, c.Namespace, configMap), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

func (c *configMapsClient) UpdateStatus(ctx context.Context, configMap *corev1.ConfigMap, opts metav1.UpdateOptions) (*corev1.ConfigMap, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(configMapsResource, c.Cluster, "status", c.Namespace, configMap), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

func (c *configMapsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(configMapsResource, c.Cluster, c.Namespace, name, opts), &corev1.ConfigMap{})
	return err
}

func (c *configMapsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(configMapsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.ConfigMapList{})
	return err
}

func (c *configMapsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1.ConfigMap, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(configMapsResource, c.Cluster, c.Namespace, name), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

// List takes label and field selectors, and returns the list of ConfigMaps that match those selectors.
func (c *configMapsClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ConfigMapList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(configMapsResource, configMapsKind, c.Cluster, c.Namespace, opts), &corev1.ConfigMapList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.ConfigMapList{ListMeta: obj.(*corev1.ConfigMapList).ListMeta}
	for _, item := range obj.(*corev1.ConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *configMapsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(configMapsResource, c.Cluster, c.Namespace, opts))
}

func (c *configMapsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.ConfigMap, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(configMapsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

func (c *configMapsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ConfigMap, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(configMapsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}

func (c *configMapsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (*corev1.ConfigMap, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(configMapsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &corev1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.ConfigMap), err
}