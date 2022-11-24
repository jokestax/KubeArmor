// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeArmor

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubearmor/KubeArmor/pkg/KubeArmorController/api/security.kubearmor.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubeArmorHostPolicyLister helps list KubeArmorHostPolicies.
// All objects returned here must be treated as read-only.
type KubeArmorHostPolicyLister interface {
	// List lists all KubeArmorHostPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KubeArmorHostPolicy, err error)
	// KubeArmorHostPolicies returns an object that can list and get KubeArmorHostPolicies.
	KubeArmorHostPolicies(namespace string) KubeArmorHostPolicyNamespaceLister
	KubeArmorHostPolicyListerExpansion
}

// kubeArmorHostPolicyLister implements the KubeArmorHostPolicyLister interface.
type kubeArmorHostPolicyLister struct {
	indexer cache.Indexer
}

// NewKubeArmorHostPolicyLister returns a new KubeArmorHostPolicyLister.
func NewKubeArmorHostPolicyLister(indexer cache.Indexer) KubeArmorHostPolicyLister {
	return &kubeArmorHostPolicyLister{indexer: indexer}
}

// List lists all KubeArmorHostPolicies in the indexer.
func (s *kubeArmorHostPolicyLister) List(selector labels.Selector) (ret []*v1.KubeArmorHostPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KubeArmorHostPolicy))
	})
	return ret, err
}

// KubeArmorHostPolicies returns an object that can list and get KubeArmorHostPolicies.
func (s *kubeArmorHostPolicyLister) KubeArmorHostPolicies(namespace string) KubeArmorHostPolicyNamespaceLister {
	return kubeArmorHostPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubeArmorHostPolicyNamespaceLister helps list and get KubeArmorHostPolicies.
// All objects returned here must be treated as read-only.
type KubeArmorHostPolicyNamespaceLister interface {
	// List lists all KubeArmorHostPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KubeArmorHostPolicy, err error)
	// Get retrieves the KubeArmorHostPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.KubeArmorHostPolicy, error)
	KubeArmorHostPolicyNamespaceListerExpansion
}

// kubeArmorHostPolicyNamespaceLister implements the KubeArmorHostPolicyNamespaceLister
// interface.
type kubeArmorHostPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubeArmorHostPolicies in the indexer for a given namespace.
func (s kubeArmorHostPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.KubeArmorHostPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KubeArmorHostPolicy))
	})
	return ret, err
}

// Get retrieves the KubeArmorHostPolicy from the indexer for a given namespace and name.
func (s kubeArmorHostPolicyNamespaceLister) Get(name string) (*v1.KubeArmorHostPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kubearmorhostpolicy"), name)
	}
	return obj.(*v1.KubeArmorHostPolicy), nil
}
