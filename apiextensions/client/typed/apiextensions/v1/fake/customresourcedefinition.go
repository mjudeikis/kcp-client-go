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

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	applyconfigurationsapiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/client/applyconfiguration/apiextensions/v1"
	apiextensionsv1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var customResourceDefinitionsResource = schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
var customResourceDefinitionsKind = schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"}

type customResourceDefinitionsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *customResourceDefinitionsClusterClient) Cluster(clusterPath logicalcluster.Path) apiextensionsv1client.CustomResourceDefinitionInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &customResourceDefinitionsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors across all clusters.
func (c *customResourceDefinitionsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*apiextensionsv1.CustomResourceDefinitionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(customResourceDefinitionsResource, customResourceDefinitionsKind, logicalcluster.Wildcard, opts), &apiextensionsv1.CustomResourceDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiextensionsv1.CustomResourceDefinitionList{ListMeta: obj.(*apiextensionsv1.CustomResourceDefinitionList).ListMeta}
	for _, item := range obj.(*apiextensionsv1.CustomResourceDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested CustomResourceDefinitions across all clusters.
func (c *customResourceDefinitionsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(customResourceDefinitionsResource, logicalcluster.Wildcard, opts))
}

type customResourceDefinitionsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *customResourceDefinitionsClient) Create(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts metav1.CreateOptions) (*apiextensionsv1.CustomResourceDefinition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(customResourceDefinitionsResource, c.ClusterPath, customResourceDefinition), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

func (c *customResourceDefinitionsClient) Update(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts metav1.UpdateOptions) (*apiextensionsv1.CustomResourceDefinition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(customResourceDefinitionsResource, c.ClusterPath, customResourceDefinition), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

func (c *customResourceDefinitionsClient) UpdateStatus(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts metav1.UpdateOptions) (*apiextensionsv1.CustomResourceDefinition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(customResourceDefinitionsResource, c.ClusterPath, "status", customResourceDefinition), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

func (c *customResourceDefinitionsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(customResourceDefinitionsResource, c.ClusterPath, name, opts), &apiextensionsv1.CustomResourceDefinition{})
	return err
}

func (c *customResourceDefinitionsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(customResourceDefinitionsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &apiextensionsv1.CustomResourceDefinitionList{})
	return err
}

func (c *customResourceDefinitionsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*apiextensionsv1.CustomResourceDefinition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(customResourceDefinitionsResource, c.ClusterPath, name), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

// List takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors.
func (c *customResourceDefinitionsClient) List(ctx context.Context, opts metav1.ListOptions) (*apiextensionsv1.CustomResourceDefinitionList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(customResourceDefinitionsResource, customResourceDefinitionsKind, c.ClusterPath, opts), &apiextensionsv1.CustomResourceDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiextensionsv1.CustomResourceDefinitionList{ListMeta: obj.(*apiextensionsv1.CustomResourceDefinitionList).ListMeta}
	for _, item := range obj.(*apiextensionsv1.CustomResourceDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *customResourceDefinitionsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(customResourceDefinitionsResource, c.ClusterPath, opts))
}

func (c *customResourceDefinitionsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*apiextensionsv1.CustomResourceDefinition, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(customResourceDefinitionsResource, c.ClusterPath, name, pt, data, subresources...), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

func (c *customResourceDefinitionsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsapiextensionsv1.CustomResourceDefinitionApplyConfiguration, opts metav1.ApplyOptions) (*apiextensionsv1.CustomResourceDefinition, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(customResourceDefinitionsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}

func (c *customResourceDefinitionsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsapiextensionsv1.CustomResourceDefinitionApplyConfiguration, opts metav1.ApplyOptions) (*apiextensionsv1.CustomResourceDefinition, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(customResourceDefinitionsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &apiextensionsv1.CustomResourceDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiextensionsv1.CustomResourceDefinition), err
}