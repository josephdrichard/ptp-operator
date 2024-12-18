// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/josephdrichard/ptp-operator/pkg/client/clientset/versioned/typed/ptp/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePtpV1 struct {
	*testing.Fake
}

func (c *FakePtpV1) NodePtpDevices(namespace string) v1.NodePtpDeviceInterface {
	return &FakeNodePtpDevices{c, namespace}
}

func (c *FakePtpV1) PtpConfigs(namespace string) v1.PtpConfigInterface {
	return &FakePtpConfigs{c, namespace}
}

func (c *FakePtpV1) PtpOperatorConfigs(namespace string) v1.PtpOperatorConfigInterface {
	return &FakePtpOperatorConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePtpV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
