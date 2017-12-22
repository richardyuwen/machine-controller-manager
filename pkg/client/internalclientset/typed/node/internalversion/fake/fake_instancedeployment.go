package fake

import (
	node "code.sapcloud.io/kubernetes/node-controller-manager/pkg/apis/node"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstanceDeployments implements InstanceDeploymentInterface
type FakeInstanceDeployments struct {
	Fake *FakeNode
}

var instancedeploymentsResource = schema.GroupVersionResource{Group: "node.sapcloud.io", Version: "", Resource: "instancedeployments"}

var instancedeploymentsKind = schema.GroupVersionKind{Group: "node.sapcloud.io", Version: "", Kind: "InstanceDeployment"}

// Get takes name of the instanceDeployment, and returns the corresponding instanceDeployment object, and an error if there is any.
func (c *FakeInstanceDeployments) Get(name string, options v1.GetOptions) (result *node.InstanceDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(instancedeploymentsResource, name), &node.InstanceDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.InstanceDeployment), err
}

// List takes label and field selectors, and returns the list of InstanceDeployments that match those selectors.
func (c *FakeInstanceDeployments) List(opts v1.ListOptions) (result *node.InstanceDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(instancedeploymentsResource, instancedeploymentsKind, opts), &node.InstanceDeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &node.InstanceDeploymentList{}
	for _, item := range obj.(*node.InstanceDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instanceDeployments.
func (c *FakeInstanceDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(instancedeploymentsResource, opts))
}

// Create takes the representation of a instanceDeployment and creates it.  Returns the server's representation of the instanceDeployment, and an error, if there is any.
func (c *FakeInstanceDeployments) Create(instanceDeployment *node.InstanceDeployment) (result *node.InstanceDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(instancedeploymentsResource, instanceDeployment), &node.InstanceDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.InstanceDeployment), err
}

// Update takes the representation of a instanceDeployment and updates it. Returns the server's representation of the instanceDeployment, and an error, if there is any.
func (c *FakeInstanceDeployments) Update(instanceDeployment *node.InstanceDeployment) (result *node.InstanceDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(instancedeploymentsResource, instanceDeployment), &node.InstanceDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.InstanceDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstanceDeployments) UpdateStatus(instanceDeployment *node.InstanceDeployment) (*node.InstanceDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(instancedeploymentsResource, "status", instanceDeployment), &node.InstanceDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.InstanceDeployment), err
}

// Delete takes name of the instanceDeployment and deletes it. Returns an error if one occurs.
func (c *FakeInstanceDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(instancedeploymentsResource, name), &node.InstanceDeployment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstanceDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(instancedeploymentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &node.InstanceDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched instanceDeployment.
func (c *FakeInstanceDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *node.InstanceDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(instancedeploymentsResource, name, data, subresources...), &node.InstanceDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.InstanceDeployment), err
}

// GetScale takes name of the instanceDeployment, and returns the corresponding scale object, and an error if there is any.
func (c *FakeInstanceDeployments) GetScale(instanceDeploymentName string, options v1.GetOptions) (result *node.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(instancedeploymentsResource, instanceDeploymentName), &node.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeInstanceDeployments) UpdateScale(instanceDeploymentName string, scale *node.Scale) (result *node.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(instancedeploymentsResource, instanceDeployment), &node.InstanceDeployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*node.Scale), err
}