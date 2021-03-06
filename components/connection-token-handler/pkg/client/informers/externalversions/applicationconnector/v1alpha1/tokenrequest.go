/*
.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	applicationconnectorv1alpha1 "github.com/kyma-project/kyma/components/connection-token-handler/pkg/apis/applicationconnector/v1alpha1"
	versioned "github.com/kyma-project/kyma/components/connection-token-handler/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyma-project/kyma/components/connection-token-handler/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kyma-project/kyma/components/connection-token-handler/pkg/client/listers/applicationconnector/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TokenRequestInformer provides access to a shared informer and lister for
// TokenRequests.
type TokenRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TokenRequestLister
}

type tokenRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTokenRequestInformer constructs a new informer for TokenRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTokenRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTokenRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTokenRequestInformer constructs a new informer for TokenRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTokenRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApplicationconnectorV1alpha1().TokenRequests(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApplicationconnectorV1alpha1().TokenRequests(namespace).Watch(options)
			},
		},
		&applicationconnectorv1alpha1.TokenRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *tokenRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTokenRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tokenRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&applicationconnectorv1alpha1.TokenRequest{}, f.defaultInformer)
}

func (f *tokenRequestInformer) Lister() v1alpha1.TokenRequestLister {
	return v1alpha1.NewTokenRequestLister(f.Informer().GetIndexer())
}
