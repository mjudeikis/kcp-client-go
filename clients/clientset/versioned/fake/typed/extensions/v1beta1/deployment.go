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

	kcpextensionsv1beta1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/extensions/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var deploymentsResource = schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "deployments"}
var deploymentsKind = schema.GroupVersionKind{Group: "extensions", Version: "v1beta1", Kind: "Deployment"}

type deploymentsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *deploymentsClusterClient) Cluster(cluster logicalcluster.Name) kcpextensionsv1beta1.DeploymentsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &deploymentsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors across all clusters.
func (c *deploymentsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.DeploymentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(deploymentsResource, deploymentsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &extensionsv1beta1.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensionsv1beta1.DeploymentList{ListMeta: obj.(*extensionsv1beta1.DeploymentList).ListMeta}
	for _, item := range obj.(*extensionsv1beta1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Deployments across all clusters.
func (c *deploymentsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(deploymentsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type deploymentsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *deploymentsNamespacer) Namespace(namespace string) extensionsv1beta1client.DeploymentInterface {
	return &deploymentsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type deploymentsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *deploymentsClient) Create(ctx context.Context, deployment *extensionsv1beta1.Deployment, opts metav1.CreateOptions) (*extensionsv1beta1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(deploymentsResource, c.Cluster, c.Namespace, deployment), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

func (c *deploymentsClient) Update(ctx context.Context, deployment *extensionsv1beta1.Deployment, opts metav1.UpdateOptions) (*extensionsv1beta1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(deploymentsResource, c.Cluster, c.Namespace, deployment), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

func (c *deploymentsClient) UpdateStatus(ctx context.Context, deployment *extensionsv1beta1.Deployment, opts metav1.UpdateOptions) (*extensionsv1beta1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(deploymentsResource, c.Cluster, "status", c.Namespace, deployment), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

func (c *deploymentsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(deploymentsResource, c.Cluster, c.Namespace, name, opts), &extensionsv1beta1.Deployment{})
	return err
}

func (c *deploymentsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(deploymentsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &extensionsv1beta1.DeploymentList{})
	return err
}

func (c *deploymentsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*extensionsv1beta1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(deploymentsResource, c.Cluster, c.Namespace, name), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deploymentsClient) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.DeploymentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(deploymentsResource, deploymentsKind, c.Cluster, c.Namespace, opts), &extensionsv1beta1.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &extensionsv1beta1.DeploymentList{ListMeta: obj.(*extensionsv1beta1.DeploymentList).ListMeta}
	for _, item := range obj.(*extensionsv1beta1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *deploymentsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(deploymentsResource, c.Cluster, c.Namespace, opts))
}

func (c *deploymentsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*extensionsv1beta1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

func (c *deploymentsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsextensionsv1beta1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.Deployment, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

func (c *deploymentsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsextensionsv1beta1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.Deployment, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &extensionsv1beta1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Deployment), err
}

func (c *deploymentsClient) GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (*extensionsv1beta1.Scale, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetSubresourceAction(deploymentsResource, c.Cluster, "scale", c.Namespace, deploymentName), &extensionsv1beta1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Scale), err
}

func (c *deploymentsClient) UpdateScale(ctx context.Context, deploymentName string, scale *extensionsv1beta1.Scale, opts metav1.UpdateOptions) (*extensionsv1beta1.Scale, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(deploymentsResource, c.Cluster, "scale", c.Namespace, scale), &extensionsv1beta1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Scale), err
}

func (c *deploymentsClient) ApplyScale(ctx context.Context, deploymentName string, applyConfiguration *applyconfigurationsextensionsv1beta1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*extensionsv1beta1.Scale, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &extensionsv1beta1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*extensionsv1beta1.Scale), err
}
