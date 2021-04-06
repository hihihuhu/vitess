/*
Copyright 2020 The Vitess Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	topov1beta1 "vitess.io/vitess/go/vt/topo/k8stopo/apis/topo/v1beta1"
	versioned "vitess.io/vitess/go/vt/topo/k8stopo/client/clientset/versioned"
	internalinterfaces "vitess.io/vitess/go/vt/topo/k8stopo/client/informers/externalversions/internalinterfaces"
	v1beta1 "vitess.io/vitess/go/vt/topo/k8stopo/client/listers/topo/v1beta1"
)

// VitessTopoNodeInformer provides access to a shared informer and lister for
// VitessTopoNodes.
type VitessTopoNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.VitessTopoNodeLister
}

type vitessTopoNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVitessTopoNodeInformer constructs a new informer for VitessTopoNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVitessTopoNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVitessTopoNodeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVitessTopoNodeInformer constructs a new informer for VitessTopoNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVitessTopoNodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopoV1beta1().VitessTopoNodes(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TopoV1beta1().VitessTopoNodes(namespace).Watch(options)
			},
		},
		&topov1beta1.VitessTopoNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *vitessTopoNodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVitessTopoNodeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vitessTopoNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&topov1beta1.VitessTopoNode{}, f.defaultInformer)
}

func (f *vitessTopoNodeInformer) Lister() v1beta1.VitessTopoNodeLister {
	return v1beta1.NewVitessTopoNodeLister(f.Informer().GetIndexer())
}
