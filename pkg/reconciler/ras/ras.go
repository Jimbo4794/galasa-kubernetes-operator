package ras

import (
	"context"
	"fmt"
	"time"

	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	ecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem"
	galasaecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	galasaecosystemlisters "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/listers/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

type Reconciler struct {
	KubeClientSet kubernetes.Interface

	GalasaEcosystemClientSet galasaecosystem.Interface
	GalasaRasLister          galasaecosystemlisters.GalasaRasComponentLister
}

func (c *Reconciler) ReconcileKind(ctx context.Context, p *v2alpha1.GalasaRasComponent) pkgreconciler.Event {
	logger := logging.FromContext(ctx)
	ras := ecosystem.Ras(p, c.GalasaEcosystemClientSet)
	objects := ras.GetObjects()

	for _, obj := range objects {
		switch obj.GetObjectKind().GroupVersionKind() {
		case schema.FromAPIVersionAndKind("v1", "Service"):
			logger.Infof("Found service: %s", obj.(*corev1.Service).Name)
			service := obj.(*corev1.Service)
			s, _ := c.KubeClientSet.CoreV1().Services(p.Namespace).Get(ctx, service.Name, v1.GetOptions{})
			if s.Name == "" {
				logger.Infof("Create service")
				_, err := c.KubeClientSet.CoreV1().Services(p.Namespace).Create(ctx, service, v1.CreateOptions{})
				if err != nil {
					return controller.NewPermanentError(fmt.Errorf("failed to create service resources: %v", err))
				}
			} else {
				logger.Infof("Service pre-existing, please manually remove service %s to apply new changes: %v", service.Name, s.Name)
			}

		case schema.FromAPIVersionAndKind("apps/v1", "StatefulSet"):
			logger.Infof("Found Stateful set: %s", obj.(*appsv1.StatefulSet).Name)
			ss := obj.(*appsv1.StatefulSet)
			s, _ := c.KubeClientSet.AppsV1().StatefulSets(p.Namespace).Get(ctx, ss.Name, v1.GetOptions{})
			if s.Name == "" {
				logger.Infof("Create Stateful set")
				_, err := c.KubeClientSet.AppsV1().StatefulSets(p.Namespace).Create(ctx, ss, v1.CreateOptions{})
				if err != nil {
					return controller.NewPermanentError(fmt.Errorf("failed to create statefulset resources: %v", err))
				}
			} else {
				logger.Infof("Updating StatefulSet with new configuration")
				_, err := c.KubeClientSet.AppsV1().StatefulSets(p.Namespace).Update(ctx, ss, v1.UpdateOptions{})
				if err != nil {
					return controller.NewPermanentError(fmt.Errorf("failed to create service resources: %v", err))
				}
			}

		case schema.FromAPIVersionAndKind("v1", "PersistentVolumeClaim"):
			logger.Infof("Found pvc: %s", obj.(*corev1.PersistentVolumeClaim).Name)
			pvc := obj.(*corev1.PersistentVolumeClaim)
			pvcG, _ := c.KubeClientSet.CoreV1().PersistentVolumeClaims(p.Namespace).Get(ctx, pvc.Name, v1.GetOptions{})
			if pvcG.Name == "" {
				logger.Infof("Create pvc: %s", pvc.Name)
				_, err := c.KubeClientSet.CoreV1().PersistentVolumeClaims(p.Namespace).Create(ctx, pvc, v1.CreateOptions{})
				if err != nil {
					return controller.NewPermanentError(fmt.Errorf("failed to create pvc resources: %v", err))
				}
			} else {
				logger.Infof("PVC found, skipping creation")
			}

		default:
			logger.Infof("Type %s was unexpected", obj.GetObjectKind().GroupVersionKind())
			return controller.NewPermanentError(fmt.Errorf("unexpected type"))
		}
	}
	statefulset, err := c.KubeClientSet.AppsV1().StatefulSets(p.Namespace).Get(ctx, p.Name, v1.GetOptions{})
	if err != nil {
		return err
	}
	if statefulset.Status.ReadyReplicas == 1 {
		p.Status = v2alpha1.ComponentStatus{
			Ready: true,
		}
		c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaRasComponents(p.Namespace).UpdateStatus(ctx, p, v1.UpdateOptions{})
		return nil
	} else {
		p.Status = v2alpha1.ComponentStatus{
			Ready: false,
		}
		c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaRasComponents(p.Namespace).UpdateStatus(ctx, p, v1.UpdateOptions{})
		return controller.NewRequeueAfter(time.Second * 3)
	}
}
