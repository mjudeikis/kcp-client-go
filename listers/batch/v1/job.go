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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	batchv1listers "k8s.io/client-go/listers/batch/v1"
	"k8s.io/client-go/tools/cache"
)

// JobClusterLister can list Jobs across all workspaces, or scope down to a JobLister for one workspace.
// All objects returned here must be treated as read-only.
type JobClusterLister interface {
	// List lists all Jobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*batchv1.Job, err error)
	// Cluster returns a lister that can list and get Jobs in one workspace.
	Cluster(cluster logicalcluster.Name) batchv1listers.JobLister
	JobClusterListerExpansion
}

type jobClusterLister struct {
	indexer cache.Indexer
}

// NewJobClusterLister returns a new JobClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewJobClusterLister(indexer cache.Indexer) *jobClusterLister {
	return &jobClusterLister{indexer: indexer}
}

// List lists all Jobs in the indexer across all workspaces.
func (s *jobClusterLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*batchv1.Job))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Jobs.
func (s *jobClusterLister) Cluster(cluster logicalcluster.Name) batchv1listers.JobLister {
	return &jobLister{indexer: s.indexer, cluster: cluster}
}

// jobLister implements the batchv1listers.JobLister interface.
type jobLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all Jobs in the indexer for a workspace.
func (s *jobLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.Job))
	})
	return ret, err
}

// Jobs returns an object that can list and get Jobs in one namespace.
func (s *jobLister) Jobs(namespace string) batchv1listers.JobNamespaceLister {
	return &jobNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// jobNamespaceLister implements the batchv1listers.JobNamespaceLister interface.
type jobNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all Jobs in the indexer for a given workspace and namespace.
func (s *jobNamespaceLister) List(selector labels.Selector) (ret []*batchv1.Job, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*batchv1.Job))
	})
	return ret, err
}

// Get retrieves the Job from the indexer for a given workspace, namespace and name.
func (s *jobNamespaceLister) Get(name string) (*batchv1.Job, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batchv1.Resource("Job"), name)
	}
	return obj.(*batchv1.Job), nil
}
