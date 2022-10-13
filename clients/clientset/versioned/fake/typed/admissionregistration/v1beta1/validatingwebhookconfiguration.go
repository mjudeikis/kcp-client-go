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

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsadmissionregistrationv1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	admissionregistrationv1beta1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var validatingWebhookConfigurationsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "v1beta1", Resource: "validatingwebhookconfigurations"}
var validatingWebhookConfigurationsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "v1beta1", Kind: "ValidatingWebhookConfiguration"}

type validatingWebhookConfigurationsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *validatingWebhookConfigurationsClusterClient) Cluster(cluster logicalcluster.Name) admissionregistrationv1beta1client.ValidatingWebhookConfigurationInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &validatingWebhookConfigurationsClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ValidatingWebhookConfigurations that match those selectors across all clusters.
func (c *validatingWebhookConfigurationsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfigurationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(validatingWebhookConfigurationsResource, validatingWebhookConfigurationsKind, logicalcluster.Wildcard, opts), &admissionregistrationv1beta1.ValidatingWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1beta1.ValidatingWebhookConfigurationList{ListMeta: obj.(*admissionregistrationv1beta1.ValidatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1beta1.ValidatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ValidatingWebhookConfigurations across all clusters.
func (c *validatingWebhookConfigurationsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(validatingWebhookConfigurationsResource, logicalcluster.Wildcard, opts))
}

type validatingWebhookConfigurationsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *validatingWebhookConfigurationsClient) Create(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1beta1.ValidatingWebhookConfiguration, opts metav1.CreateOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(validatingWebhookConfigurationsResource, c.Cluster, validatingWebhookConfiguration), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}

func (c *validatingWebhookConfigurationsClient) Update(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1beta1.ValidatingWebhookConfiguration, opts metav1.UpdateOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(validatingWebhookConfigurationsResource, c.Cluster, validatingWebhookConfiguration), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}

func (c *validatingWebhookConfigurationsClient) UpdateStatus(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1beta1.ValidatingWebhookConfiguration, opts metav1.UpdateOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(validatingWebhookConfigurationsResource, c.Cluster, "status", validatingWebhookConfiguration), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}

func (c *validatingWebhookConfigurationsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(validatingWebhookConfigurationsResource, c.Cluster, name, opts), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	return err
}

func (c *validatingWebhookConfigurationsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(validatingWebhookConfigurationsResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &admissionregistrationv1beta1.ValidatingWebhookConfigurationList{})
	return err
}

func (c *validatingWebhookConfigurationsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(validatingWebhookConfigurationsResource, c.Cluster, name), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}

// List takes label and field selectors, and returns the list of ValidatingWebhookConfigurations that match those selectors.
func (c *validatingWebhookConfigurationsClient) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfigurationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(validatingWebhookConfigurationsResource, validatingWebhookConfigurationsKind, c.Cluster, opts), &admissionregistrationv1beta1.ValidatingWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1beta1.ValidatingWebhookConfigurationList{ListMeta: obj.(*admissionregistrationv1beta1.ValidatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1beta1.ValidatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *validatingWebhookConfigurationsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(validatingWebhookConfigurationsResource, c.Cluster, opts))
}

func (c *validatingWebhookConfigurationsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingWebhookConfigurationsResource, c.Cluster, name, pt, data, subresources...), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}

func (c *validatingWebhookConfigurationsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsadmissionregistrationv1beta1.ValidatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingWebhookConfigurationsResource, c.Cluster, *name, types.ApplyPatchType, data), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}

func (c *validatingWebhookConfigurationsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsadmissionregistrationv1beta1.ValidatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1beta1.ValidatingWebhookConfiguration, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingWebhookConfigurationsResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &admissionregistrationv1beta1.ValidatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingWebhookConfiguration), err
}
