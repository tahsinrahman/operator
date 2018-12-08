/*
Copyright 2018 The KubeDB Authors.

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

package v1alpha1

import (
	time "time"

	kubedbv1alpha1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	versioned "github.com/kubedb/apimachinery/client/clientset/versioned"
	internalinterfaces "github.com/kubedb/apimachinery/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubedb/apimachinery/client/listers/kubedb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PostgresInformer provides access to a shared informer and lister for
// Postgreses.
type PostgresInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PostgresLister
}

type postgresInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPostgresInformer constructs a new informer for Postgres type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPostgresInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPostgresInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPostgresInformer constructs a new informer for Postgres type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPostgresInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedbV1alpha1().Postgreses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedbV1alpha1().Postgreses(namespace).Watch(options)
			},
		},
		&kubedbv1alpha1.Postgres{},
		resyncPeriod,
		indexers,
	)
}

func (f *postgresInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPostgresInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *postgresInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubedbv1alpha1.Postgres{}, f.defaultInformer)
}

func (f *postgresInformer) Lister() v1alpha1.PostgresLister {
	return v1alpha1.NewPostgresLister(f.Informer().GetIndexer())
}