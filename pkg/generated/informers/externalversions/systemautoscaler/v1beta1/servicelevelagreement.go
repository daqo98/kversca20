// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	systemautoscalerv1beta1 "github.com/lterrac/system-autoscaler/pkg/apis/systemautoscaler/v1beta1"
	versioned "github.com/lterrac/system-autoscaler/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/lterrac/system-autoscaler/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/lterrac/system-autoscaler/pkg/generated/listers/systemautoscaler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceLevelAgreementInformer provides access to a shared informer and lister for
// ServiceLevelAgreements.
type ServiceLevelAgreementInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ServiceLevelAgreementLister
}

type serviceLevelAgreementInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceLevelAgreementInformer constructs a new informer for ServiceLevelAgreement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceLevelAgreementInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceLevelAgreementInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceLevelAgreementInformer constructs a new informer for ServiceLevelAgreement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceLevelAgreementInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SystemautoscalerV1beta1().ServiceLevelAgreements(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SystemautoscalerV1beta1().ServiceLevelAgreements(namespace).Watch(context.TODO(), options)
			},
		},
		&systemautoscalerv1beta1.ServiceLevelAgreement{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceLevelAgreementInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceLevelAgreementInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceLevelAgreementInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&systemautoscalerv1beta1.ServiceLevelAgreement{}, f.defaultInformer)
}

func (f *serviceLevelAgreementInformer) Lister() v1beta1.ServiceLevelAgreementLister {
	return v1beta1.NewServiceLevelAgreementLister(f.Informer().GetIndexer())
}
