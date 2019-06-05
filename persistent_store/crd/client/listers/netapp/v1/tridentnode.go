// Copyright 2019 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentNodeLister helps list TridentNodes.
type TridentNodeLister interface {
	// List lists all TridentNodes in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentNode, err error)
	// TridentNodes returns an object that can list and get TridentNodes.
	TridentNodes(namespace string) TridentNodeNamespaceLister
	TridentNodeListerExpansion
}

// tridentNodeLister implements the TridentNodeLister interface.
type tridentNodeLister struct {
	indexer cache.Indexer
}

// NewTridentNodeLister returns a new TridentNodeLister.
func NewTridentNodeLister(indexer cache.Indexer) TridentNodeLister {
	return &tridentNodeLister{indexer: indexer}
}

// List lists all TridentNodes in the indexer.
func (s *tridentNodeLister) List(selector labels.Selector) (ret []*v1.TridentNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentNode))
	})
	return ret, err
}

// TridentNodes returns an object that can list and get TridentNodes.
func (s *tridentNodeLister) TridentNodes(namespace string) TridentNodeNamespaceLister {
	return tridentNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentNodeNamespaceLister helps list and get TridentNodes.
type TridentNodeNamespaceLister interface {
	// List lists all TridentNodes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentNode, err error)
	// Get retrieves the TridentNode from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentNode, error)
	TridentNodeNamespaceListerExpansion
}

// tridentNodeNamespaceLister implements the TridentNodeNamespaceLister
// interface.
type tridentNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentNodes in the indexer for a given namespace.
func (s tridentNodeNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentNode))
	})
	return ret, err
}

// Get retrieves the TridentNode from the indexer for a given namespace and name.
func (s tridentNodeNamespaceLister) Get(name string) (*v1.TridentNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentnode"), name)
	}
	return obj.(*v1.TridentNode), nil
}
