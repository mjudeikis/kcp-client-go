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

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	networkingv1beta1client "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// IPAddressesClusterGetter has a method to return a IPAddressClusterInterface.
// A group's cluster client should implement this interface.
type IPAddressesClusterGetter interface {
	IPAddresses() IPAddressClusterInterface
}

// IPAddressClusterInterface can operate on IPAddresses across all clusters,
// or scope down to one cluster and return a networkingv1beta1client.IPAddressInterface.
type IPAddressClusterInterface interface {
	Cluster(logicalcluster.Path) networkingv1beta1client.IPAddressInterface
	List(ctx context.Context, opts metav1.ListOptions) (*networkingv1beta1.IPAddressList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type iPAddressesClusterInterface struct {
	clientCache kcpclient.Cache[*networkingv1beta1client.NetworkingV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *iPAddressesClusterInterface) Cluster(clusterPath logicalcluster.Path) networkingv1beta1client.IPAddressInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).IPAddresses()
}

// List returns the entire collection of all IPAddresses across all clusters.
func (c *iPAddressesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*networkingv1beta1.IPAddressList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).IPAddresses().List(ctx, opts)
}

// Watch begins to watch all IPAddresses across all clusters.
func (c *iPAddressesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).IPAddresses().Watch(ctx, opts)
}