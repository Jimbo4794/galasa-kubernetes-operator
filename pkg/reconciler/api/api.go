package api

import (
	"context"

	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	galasaecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	galasaecosystemlisters "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/listers/galasaecosystem/v2alpha1"
	"k8s.io/client-go/kubernetes"
)

type Reconciler struct {
	KubeClientSet kubernetes.Interface

	GalasaEcosystemClientSet galasaecosystem.Interface
	GalasaEcosystemLister    galasaecosystemlisters.GalasaEcosystemLister
}

func (c *Reconciler) ReconcileKind(ctx context.Context, p *v2alpha1.GalasaApi) pkgreconciler.Event {
	logger := logging.FromContext(ctx)
	logger.Infof("Hello World")
	return nil
}
