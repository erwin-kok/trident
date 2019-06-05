// Copyright 2019 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTridentVolumeTransactions implements TridentVolumeTransactionInterface
type FakeTridentVolumeTransactions struct {
	Fake *FakeTridentV1
	ns   string
}

var tridentvolumetransactionsResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentvolumetransactions"}

var tridentvolumetransactionsKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentTransaction"}

// Get takes name of the tridentVolumeTransaction, and returns the corresponding tridentVolumeTransaction object, and an error if there is any.
func (c *FakeTridentVolumeTransactions) Get(name string, options v1.GetOptions) (result *netappv1.TridentTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tridentvolumetransactionsResource, c.ns, name), &netappv1.TridentTransaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentTransaction), err
}

// List takes label and field selectors, and returns the list of TridentVolumeTransactions that match those selectors.
func (c *FakeTridentVolumeTransactions) List(opts v1.ListOptions) (result *netappv1.TridentTransactionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tridentvolumetransactionsResource, tridentvolumetransactionsKind, c.ns, opts), &netappv1.TridentTransactionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentTransactionList{ListMeta: obj.(*netappv1.TridentTransactionList).ListMeta}
	for _, item := range obj.(*netappv1.TridentTransactionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentVolumeTransactions.
func (c *FakeTridentVolumeTransactions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tridentvolumetransactionsResource, c.ns, opts))

}

// Create takes the representation of a tridentVolumeTransaction and creates it.  Returns the server's representation of the tridentVolumeTransaction, and an error, if there is any.
func (c *FakeTridentVolumeTransactions) Create(tridentVolumeTransaction *netappv1.TridentTransaction) (result *netappv1.TridentTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tridentvolumetransactionsResource, c.ns, tridentVolumeTransaction), &netappv1.TridentTransaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentTransaction), err
}

// Update takes the representation of a tridentVolumeTransaction and updates it. Returns the server's representation of the tridentVolumeTransaction, and an error, if there is any.
func (c *FakeTridentVolumeTransactions) Update(tridentVolumeTransaction *netappv1.TridentTransaction) (result *netappv1.TridentTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tridentvolumetransactionsResource, c.ns, tridentVolumeTransaction), &netappv1.TridentTransaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentTransaction), err
}

// Delete takes name of the tridentVolumeTransaction and deletes it. Returns an error if one occurs.
func (c *FakeTridentVolumeTransactions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tridentvolumetransactionsResource, c.ns, name), &netappv1.TridentTransaction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentVolumeTransactions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tridentvolumetransactionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &netappv1.TridentTransactionList{})
	return err
}

// Patch applies the patch and returns the patched tridentVolumeTransaction.
func (c *FakeTridentVolumeTransactions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *netappv1.TridentTransaction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tridentvolumetransactionsResource, c.ns, name, pt, data, subresources...), &netappv1.TridentTransaction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentTransaction), err
}
