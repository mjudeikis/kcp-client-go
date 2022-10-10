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

	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	"k8s.io/client-go/rest"

	kcpauthenticationv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/authentication/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpauthenticationv1.AuthenticationV1ClusterInterface = (*AuthenticationV1ClusterClient)(nil)

type AuthenticationV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AuthenticationV1ClusterClient) Cluster(cluster logicalcluster.Name) authenticationv1.AuthenticationV1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AuthenticationV1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *AuthenticationV1ClusterClient) TokenReviews() kcpauthenticationv1.TokenReviewClusterInterface {
	return &tokenReviewsClusterClient{Fake: c.Fake}
}

var _ authenticationv1.AuthenticationV1Interface = (*AuthenticationV1Client)(nil)

type AuthenticationV1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *AuthenticationV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AuthenticationV1Client) TokenReviews() authenticationv1.TokenReviewInterface {
	return &tokenReviewsClient{Fake: c.Fake, Cluster: c.Cluster}
}
