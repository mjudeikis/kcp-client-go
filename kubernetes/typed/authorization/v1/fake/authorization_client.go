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
	"github.com/kcp-dev/logicalcluster/v2"

	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/client-go/rest"

	kcpauthorizationv1 "github.com/kcp-dev/client-go/kubernetes/typed/authorization/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpauthorizationv1.AuthorizationV1ClusterInterface = (*AuthorizationV1ClusterClient)(nil)

type AuthorizationV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AuthorizationV1ClusterClient) Cluster(cluster logicalcluster.Name) authorizationv1.AuthorizationV1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AuthorizationV1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *AuthorizationV1ClusterClient) SubjectAccessReviews() kcpauthorizationv1.SubjectAccessReviewClusterInterface {
	return &subjectAccessReviewsClusterClient{Fake: c.Fake}
}

func (c *AuthorizationV1ClusterClient) SelfSubjectAccessReviews() kcpauthorizationv1.SelfSubjectAccessReviewClusterInterface {
	return &selfSubjectAccessReviewsClusterClient{Fake: c.Fake}
}

func (c *AuthorizationV1ClusterClient) LocalSubjectAccessReviews() kcpauthorizationv1.LocalSubjectAccessReviewClusterInterface {
	return &localSubjectAccessReviewsClusterClient{Fake: c.Fake}
}

func (c *AuthorizationV1ClusterClient) SelfSubjectRulesReviews() kcpauthorizationv1.SelfSubjectRulesReviewClusterInterface {
	return &selfSubjectRulesReviewsClusterClient{Fake: c.Fake}
}

var _ authorizationv1.AuthorizationV1Interface = (*AuthorizationV1Client)(nil)

type AuthorizationV1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *AuthorizationV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AuthorizationV1Client) SubjectAccessReviews() authorizationv1.SubjectAccessReviewInterface {
	return &subjectAccessReviewsClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *AuthorizationV1Client) SelfSubjectAccessReviews() authorizationv1.SelfSubjectAccessReviewInterface {
	return &selfSubjectAccessReviewsClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *AuthorizationV1Client) LocalSubjectAccessReviews(namespace string) authorizationv1.LocalSubjectAccessReviewInterface {
	return &localSubjectAccessReviewsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *AuthorizationV1Client) SelfSubjectRulesReviews() authorizationv1.SelfSubjectRulesReviewInterface {
	return &selfSubjectRulesReviewsClient{Fake: c.Fake, Cluster: c.Cluster}
}