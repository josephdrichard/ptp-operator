// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/josephdrichard/ptp-operator/api/v1"
	scheme "github.com/josephdrichard/ptp-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PtpConfigsGetter has a method to return a PtpConfigInterface.
// A group's client should implement this interface.
type PtpConfigsGetter interface {
	PtpConfigs(namespace string) PtpConfigInterface
}

// PtpConfigInterface has methods to work with PtpConfig resources.
type PtpConfigInterface interface {
	Create(ctx context.Context, ptpConfig *v1.PtpConfig, opts metav1.CreateOptions) (*v1.PtpConfig, error)
	Update(ctx context.Context, ptpConfig *v1.PtpConfig, opts metav1.UpdateOptions) (*v1.PtpConfig, error)
	UpdateStatus(ctx context.Context, ptpConfig *v1.PtpConfig, opts metav1.UpdateOptions) (*v1.PtpConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PtpConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PtpConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PtpConfig, err error)
	PtpConfigExpansion
}

// ptpConfigs implements PtpConfigInterface
type ptpConfigs struct {
	client rest.Interface
	ns     string
}

// newPtpConfigs returns a PtpConfigs
func newPtpConfigs(c *PtpV1Client, namespace string) *ptpConfigs {
	return &ptpConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ptpConfig, and returns the corresponding ptpConfig object, and an error if there is any.
func (c *ptpConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PtpConfig, err error) {
	result = &v1.PtpConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ptpconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PtpConfigs that match those selectors.
func (c *ptpConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PtpConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PtpConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ptpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ptpConfigs.
func (c *ptpConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ptpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a ptpConfig and creates it.  Returns the server's representation of the ptpConfig, and an error, if there is any.
func (c *ptpConfigs) Create(ctx context.Context, ptpConfig *v1.PtpConfig, opts metav1.CreateOptions) (result *v1.PtpConfig, err error) {
	result = &v1.PtpConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ptpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ptpConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a ptpConfig and updates it. Returns the server's representation of the ptpConfig, and an error, if there is any.
func (c *ptpConfigs) Update(ctx context.Context, ptpConfig *v1.PtpConfig, opts metav1.UpdateOptions) (result *v1.PtpConfig, err error) {
	result = &v1.PtpConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ptpconfigs").
		Name(ptpConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ptpConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *ptpConfigs) UpdateStatus(ctx context.Context, ptpConfig *v1.PtpConfig, opts metav1.UpdateOptions) (result *v1.PtpConfig, err error) {
	result = &v1.PtpConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ptpconfigs").
		Name(ptpConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ptpConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the ptpConfig and deletes it. Returns an error if one occurs.
func (c *ptpConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ptpconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ptpConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ptpconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched ptpConfig.
func (c *ptpConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PtpConfig, err error) {
	result = &v1.PtpConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ptpconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
