/*
Copyright 2019 The Knative Authors

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

package rssfeed

import (
	"context"

	svcinformer "github.com/knative/pkg/injection/informers/kubeinformers/corev1/service"
	asclient "github.com/n3wscott/rssfeed/pkg/client/injection/client"
	rssfeedinformer "github.com/n3wscott/rssfeed/pkg/client/injection/informers/sources/v1alpha1/rssfeed"

	"github.com/knative/pkg/configmap"
	"github.com/knative/pkg/controller"
	"github.com/knative/pkg/logging"
	"github.com/knative/pkg/tracker"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
)

const (
	controllerAgentName = "rssfeed-controller"
)

// NewController returns a new HPA reconcile controller.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	rssfeedInformer := rssfeedinformer.Get(ctx)
	svcInformer := svcinformer.Get(ctx)

	c := &Reconciler{
		Client:        asclient.Get(ctx),
		Lister:        rssfeedInformer.Lister(),
		ServiceLister: svcInformer.Lister(),
		Recorder: record.NewBroadcaster().NewRecorder(
			scheme.Scheme, corev1.EventSource{Component: controllerAgentName}),
	}
	impl := controller.NewImpl(c, logger, "RssFeed")

	logger.Info("Setting up event handlers")

	rssfeedInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	c.Tracker = tracker.New(impl.EnqueueKey, controller.GetTrackerLease(ctx))
	svcInformer.Informer().AddEventHandler(controller.HandleAll(
		// Call the tracker's OnChanged method, but we've seen the objects
		// coming through this path missing TypeMeta, so ensure it is properly
		// populated.
		controller.EnsureTypeMeta(
			c.Tracker.OnChanged,
			corev1.SchemeGroupVersion.WithKind("Service"),
		),
	))
	return impl
}
