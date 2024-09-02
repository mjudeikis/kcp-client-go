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

package v1alpha3

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	resourcev1alpha3 "k8s.io/api/resource/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	resourcev1alpha3listers "k8s.io/client-go/listers/resource/v1alpha3"
	"k8s.io/client-go/tools/cache"
)

// DeviceClassClusterLister can list DeviceClasses across all workspaces, or scope down to a DeviceClassLister for one workspace.
// All objects returned here must be treated as read-only.
type DeviceClassClusterLister interface {
	// List lists all DeviceClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*resourcev1alpha3.DeviceClass, err error)
	// Cluster returns a lister that can list and get DeviceClasses in one workspace.
	Cluster(clusterName logicalcluster.Name) resourcev1alpha3listers.DeviceClassLister
	DeviceClassClusterListerExpansion
}

type deviceClassClusterLister struct {
	indexer cache.Indexer
}

// NewDeviceClassClusterLister returns a new DeviceClassClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewDeviceClassClusterLister(indexer cache.Indexer) *deviceClassClusterLister {
	return &deviceClassClusterLister{indexer: indexer}
}

// List lists all DeviceClasses in the indexer across all workspaces.
func (s *deviceClassClusterLister) List(selector labels.Selector) (ret []*resourcev1alpha3.DeviceClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*resourcev1alpha3.DeviceClass))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get DeviceClasses.
func (s *deviceClassClusterLister) Cluster(clusterName logicalcluster.Name) resourcev1alpha3listers.DeviceClassLister {
	return &deviceClassLister{indexer: s.indexer, clusterName: clusterName}
}

// deviceClassLister implements the resourcev1alpha3listers.DeviceClassLister interface.
type deviceClassLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all DeviceClasses in the indexer for a workspace.
func (s *deviceClassLister) List(selector labels.Selector) (ret []*resourcev1alpha3.DeviceClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*resourcev1alpha3.DeviceClass))
	})
	return ret, err
}

// Get retrieves the DeviceClass from the indexer for a given workspace and name.
func (s *deviceClassLister) Get(name string) (*resourcev1alpha3.DeviceClass, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(resourcev1alpha3.Resource("deviceclasses"), name)
	}
	return obj.(*resourcev1alpha3.DeviceClass), nil
}