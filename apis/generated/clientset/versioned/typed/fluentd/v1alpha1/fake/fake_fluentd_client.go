/*
Copyright 2022.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/fluent/fluent-operator/v2/apis/generated/clientset/versioned/typed/fluentd/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeFluentdV1alpha1 struct {
	*testing.Fake
}

func (c *FakeFluentdV1alpha1) ClusterFilters() v1alpha1.ClusterFilterInterface {
	return &FakeClusterFilters{c}
}

func (c *FakeFluentdV1alpha1) ClusterFluentdConfigs() v1alpha1.ClusterFluentdConfigInterface {
	return &FakeClusterFluentdConfigs{c}
}

func (c *FakeFluentdV1alpha1) ClusterInputs() v1alpha1.ClusterInputInterface {
	return &FakeClusterInputs{c}
}

func (c *FakeFluentdV1alpha1) ClusterOutputs() v1alpha1.ClusterOutputInterface {
	return &FakeClusterOutputs{c}
}

func (c *FakeFluentdV1alpha1) Filters(namespace string) v1alpha1.FilterInterface {
	return &FakeFilters{c, namespace}
}

func (c *FakeFluentdV1alpha1) Fluentds(namespace string) v1alpha1.FluentdInterface {
	return &FakeFluentds{c, namespace}
}

func (c *FakeFluentdV1alpha1) FluentdConfigs(namespace string) v1alpha1.FluentdConfigInterface {
	return &FakeFluentdConfigs{c, namespace}
}

func (c *FakeFluentdV1alpha1) Inputs(namespace string) v1alpha1.InputInterface {
	return &FakeInputs{c, namespace}
}

func (c *FakeFluentdV1alpha1) Outputs(namespace string) v1alpha1.OutputInterface {
	return &FakeOutputs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeFluentdV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
