/*
Copyright The Kubernetes Authors.

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
	"context"
	time "time"

	workflowcontrollerv1alpha1 "github.com/cloud-ark/kubeplus/platform-operator/pkg/apis/workflowcontroller/v1alpha1"
	versioned "github.com/cloud-ark/kubeplus/platform-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/cloud-ark/kubeplus/platform-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cloud-ark/kubeplus/platform-operator/pkg/generated/listers/workflowcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceMonitorInformer provides access to a shared informer and lister for
// ResourceMonitors.
type ResourceMonitorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceMonitorLister
}

type resourceMonitorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceMonitorInformer constructs a new informer for ResourceMonitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceMonitorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceMonitorInformer constructs a new informer for ResourceMonitor type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceMonitorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkflowsV1alpha1().ResourceMonitors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkflowsV1alpha1().ResourceMonitors(namespace).Watch(context.TODO(), options)
			},
		},
		&workflowcontrollerv1alpha1.ResourceMonitor{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceMonitorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceMonitorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceMonitorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workflowcontrollerv1alpha1.ResourceMonitor{}, f.defaultInformer)
}

func (f *resourceMonitorInformer) Lister() v1alpha1.ResourceMonitorLister {
	return v1alpha1.NewResourceMonitorLister(f.Informer().GetIndexer())
}