/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1

import (
	time "time"

	aslangroup_io_v1 "github.com/aslanbekirov/personcrd/pkg/apis/aslangroup.io/v1"
	versioned "github.com/aslanbekirov/personcrd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/aslanbekirov/personcrd/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/aslanbekirov/personcrd/pkg/client/listers/aslangroup.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PersonInformer provides access to a shared informer and lister for
// Persons.
type PersonInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PersonLister
}

type personInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPersonInformer constructs a new informer for Person type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPersonInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPersonInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPersonInformer constructs a new informer for Person type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPersonInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AslangroupV1().Persons(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AslangroupV1().Persons(namespace).Watch(options)
			},
		},
		&aslangroup_io_v1.Person{},
		resyncPeriod,
		indexers,
	)
}

func (f *personInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPersonInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *personInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aslangroup_io_v1.Person{}, f.defaultInformer)
}

func (f *personInformer) Lister() v1.PersonLister {
	return v1.NewPersonLister(f.Informer().GetIndexer())
}
