/*
Copyright 2019 The Knative Authors
Copyright 2019 Scott Nichols

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/n3wscott/rssfeed/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RssFeedLister helps list RssFeeds.
type RssFeedLister interface {
	// List lists all RssFeeds in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RssFeed, err error)
	// RssFeeds returns an object that can list and get RssFeeds.
	RssFeeds(namespace string) RssFeedNamespaceLister
	RssFeedListerExpansion
}

// rssFeedLister implements the RssFeedLister interface.
type rssFeedLister struct {
	indexer cache.Indexer
}

// NewRssFeedLister returns a new RssFeedLister.
func NewRssFeedLister(indexer cache.Indexer) RssFeedLister {
	return &rssFeedLister{indexer: indexer}
}

// List lists all RssFeeds in the indexer.
func (s *rssFeedLister) List(selector labels.Selector) (ret []*v1alpha1.RssFeed, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RssFeed))
	})
	return ret, err
}

// RssFeeds returns an object that can list and get RssFeeds.
func (s *rssFeedLister) RssFeeds(namespace string) RssFeedNamespaceLister {
	return rssFeedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RssFeedNamespaceLister helps list and get RssFeeds.
type RssFeedNamespaceLister interface {
	// List lists all RssFeeds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RssFeed, err error)
	// Get retrieves the RssFeed from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RssFeed, error)
	RssFeedNamespaceListerExpansion
}

// rssFeedNamespaceLister implements the RssFeedNamespaceLister
// interface.
type rssFeedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RssFeeds in the indexer for a given namespace.
func (s rssFeedNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RssFeed, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RssFeed))
	})
	return ret, err
}

// Get retrieves the RssFeed from the indexer for a given namespace and name.
func (s rssFeedNamespaceLister) Get(name string) (*v1alpha1.RssFeed, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rssfeed"), name)
	}
	return obj.(*v1alpha1.RssFeed), nil
}
