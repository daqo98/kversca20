// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/lterrac/system-autoscaler/pkg/apis/systemautoscaler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceLevelAgreements implements ServiceLevelAgreementInterface
type FakeServiceLevelAgreements struct {
	Fake *FakeSystemautoscalerV1beta1
	ns   string
}

var servicelevelagreementsResource = schema.GroupVersionResource{Group: "systemautoscaler.polimi.it", Version: "v1beta1", Resource: "servicelevelagreements"}

var servicelevelagreementsKind = schema.GroupVersionKind{Group: "systemautoscaler.polimi.it", Version: "v1beta1", Kind: "ServiceLevelAgreement"}

// Get takes name of the serviceLevelAgreement, and returns the corresponding serviceLevelAgreement object, and an error if there is any.
func (c *FakeServiceLevelAgreements) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ServiceLevelAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicelevelagreementsResource, c.ns, name), &v1beta1.ServiceLevelAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceLevelAgreement), err
}

// List takes label and field selectors, and returns the list of ServiceLevelAgreements that match those selectors.
func (c *FakeServiceLevelAgreements) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ServiceLevelAgreementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicelevelagreementsResource, servicelevelagreementsKind, c.ns, opts), &v1beta1.ServiceLevelAgreementList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceLevelAgreementList{ListMeta: obj.(*v1beta1.ServiceLevelAgreementList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceLevelAgreementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceLevelAgreements.
func (c *FakeServiceLevelAgreements) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicelevelagreementsResource, c.ns, opts))

}

// Create takes the representation of a serviceLevelAgreement and creates it.  Returns the server's representation of the serviceLevelAgreement, and an error, if there is any.
func (c *FakeServiceLevelAgreements) Create(ctx context.Context, serviceLevelAgreement *v1beta1.ServiceLevelAgreement, opts v1.CreateOptions) (result *v1beta1.ServiceLevelAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicelevelagreementsResource, c.ns, serviceLevelAgreement), &v1beta1.ServiceLevelAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceLevelAgreement), err
}

// Update takes the representation of a serviceLevelAgreement and updates it. Returns the server's representation of the serviceLevelAgreement, and an error, if there is any.
func (c *FakeServiceLevelAgreements) Update(ctx context.Context, serviceLevelAgreement *v1beta1.ServiceLevelAgreement, opts v1.UpdateOptions) (result *v1beta1.ServiceLevelAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicelevelagreementsResource, c.ns, serviceLevelAgreement), &v1beta1.ServiceLevelAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceLevelAgreement), err
}

// Delete takes name of the serviceLevelAgreement and deletes it. Returns an error if one occurs.
func (c *FakeServiceLevelAgreements) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicelevelagreementsResource, c.ns, name), &v1beta1.ServiceLevelAgreement{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceLevelAgreements) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicelevelagreementsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ServiceLevelAgreementList{})
	return err
}

// Patch applies the patch and returns the patched serviceLevelAgreement.
func (c *FakeServiceLevelAgreements) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ServiceLevelAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicelevelagreementsResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceLevelAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceLevelAgreement), err
}
