package galasaecosystem

import (
	"context"
	"fmt"
	"time"

	ecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"

	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	galasaecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned"
	galasaecosystemlisters "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/listers/galasaecosystem/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
)

type Reconciler struct {
	KubeClientSet kubernetes.Interface

	GalasaEcosystemClientSet     galasaecosystem.Interface
	GalasaEcosystemLister        galasaecosystemlisters.GalasaEcosystemLister
	GalasaCPSLister              galasaecosystemlisters.GalasaCpsComponentLister
	GalasaRASLister              galasaecosystemlisters.GalasaRasComponentLister
	GalasaAPILister              galasaecosystemlisters.GalasaApiComponentLister
	GalasaResmonLister           galasaecosystemlisters.GalasaResmonComponentLister
	GalasaEngineControllerLister galasaecosystemlisters.GalasaEngineControllerComponentLister
	GalasaMetricsLister          galasaecosystemlisters.GalasaMetricsComponentLister
	GalasaToolboxLister          galasaecosystemlisters.GalasaToolboxComponentLister

	Cps              *v2alpha1.GalasaCpsComponent
	Ras              *v2alpha1.GalasaRasComponent
	Api              *v2alpha1.GalasaApiComponent
	Metrics          *v2alpha1.GalasaMetricsComponent
	Resmon           *v2alpha1.GalasaResmonComponent
	EngineController *v2alpha1.GalasaEngineControllerComponent
	Toolbox          *v2alpha1.GalasaToolboxComponent
}

func (c *Reconciler) ReconcileKind(ctx context.Context, p *v2alpha1.GalasaEcosystem) pkgreconciler.Event {
	// p.validate - Needs this to check all components are created and in p
	logger := logging.FromContext(ctx)
	selector := labels.NewSelector().Add(mustNewRequirement("galasa-ecosystem-name", selection.Equals, []string{p.Name}))

	logger.Info("Managing CPS")
	err := c.ManageCps(ctx, p, selector)
	if err != nil {
		return err
	}

	logger.Info("Managing RAS")
	err = c.ManageRas(ctx, p, selector)
	if err != nil {
		return err
	}

	logger.Info("Managing API")
	err = c.ManageApi(ctx, p, selector)
	if err != nil {
		return err
	}

	logger.Info("Managing Metrics")
	err = c.ManageMetrics(ctx, p, selector)
	if err != nil {
		return err
	}

	logger.Info("Managing EC")
	err = c.ManageEngineController(ctx, p, selector)
	if err != nil {
		return err
	}

	logger.Info("Managing Resmon")
	err = c.ManageResmon(ctx, p, selector)
	if err != nil {
		return err
	}

	return nil
}

func (c *Reconciler) ManageCps(ctx context.Context, p *v2alpha1.GalasaEcosystem, selector labels.Selector) error {
	l := logging.FromContext(ctx)
	cpslist, err := c.GalasaCPSLister.List(selector)
	if err != nil {
		return controller.NewPermanentError(fmt.Errorf("failed to retrieve cps: %v", err))
	}

	cpsSpec := p.Spec.ComponentsSpec["cps"]
	if len(cpslist) == 0 {
		l.Infof("No CPS detected, creating CRD")
		// Create CPS CRD
		t := true
		cps := &v2alpha1.GalasaCpsComponent{
			ObjectMeta: v1.ObjectMeta{
				Name:      "cps-" + p.Name,
				Namespace: p.Namespace,
				Labels: map[string]string{
					"galasa-ecosystem-name": p.Name,
				},
				OwnerReferences: []v1.OwnerReference{
					{
						APIVersion:         "GalasaEcosystem",
						Kind:               "galasa.dev/v2alpha1",
						Name:               p.Name,
						UID:                p.GetUID(),
						Controller:         &t,
						BlockOwnerDeletion: &t,
					},
				},
			},
			Spec: v2alpha1.ComponentSpec{
				Image:            cpsSpec.Image,
				ImagePullPolicy:  cpsSpec.ImagePullPolicy,
				Storage:          cpsSpec.Storage,
				StorageClassName: cpsSpec.StorageClassName,
				NodeSelector:     cpsSpec.NodeSelector,
				ComponentParms: map[string]string{
					"hostname": p.Spec.Hostname,
				},
			},
		}
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaCpsComponents(p.Namespace).Create(ctx, cps, v1.CreateOptions{})
		c.Cps = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to create cps: %v", err))
		}
		return controller.NewRequeueAfter(5 * time.Second)
	}
	l.Infof("CPS detected, checking state")
	if len(cpslist) > 1 {
		return controller.NewPermanentError(fmt.Errorf("too many cps's defined!"))
	}
	cps := ecosystem.Cps(cpslist[0], c.GalasaEcosystemClientSet)
	if !cps.IsReady(ctx) {
		l.Infof("CPS not ready, waiting, %s", cpslist[0].Status)
		return controller.NewRequeueAfter(time.Second * 5)
	}
	if cps.HasChanged(cpsSpec) {
		l.Infof("CPS changes detected")
		cpsUpdate := cpslist[0]
		cpsUpdate.Spec.Image = cpsSpec.Image
		cpsUpdate.Spec.ImagePullPolicy = cpsSpec.ImagePullPolicy
		cpsUpdate.Spec.Storage = cpsSpec.Storage
		cpsUpdate.Spec.StorageClassName = cpsSpec.StorageClassName
		cpsUpdate.Spec.NodeSelector = cpsSpec.NodeSelector
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaCpsComponents(p.Namespace).Update(ctx, cpsUpdate, v1.UpdateOptions{})
		c.Cps = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to update cps: %v", err))
		}
	}
	l.Infof("CPS finished, ending")
	return nil
}

func (c *Reconciler) ManageRas(ctx context.Context, p *v2alpha1.GalasaEcosystem, selector labels.Selector) error {
	l := logging.FromContext(ctx)
	raslist, err := c.GalasaRASLister.List(selector)
	if err != nil {
		return controller.NewPermanentError(fmt.Errorf("failed to retrieve ras: %v", err))
	}
	l.Infof("RasList: %v", raslist)
	rasSpec := p.Spec.ComponentsSpec["ras"]
	if len(raslist) == 0 {
		// Create RAS CRD
		t := true
		ras := &v2alpha1.GalasaRasComponent{
			ObjectMeta: v1.ObjectMeta{
				Name:      "ras-" + p.Name,
				Namespace: p.Namespace,
				Labels: map[string]string{
					"galasa-ecosystem-name": p.Name,
				},
				OwnerReferences: []v1.OwnerReference{
					{
						APIVersion:         "GalasaEcosystem",
						Kind:               "galasa.dev/v2alpha1",
						Name:               p.Name,
						UID:                p.GetUID(),
						Controller:         &t,
						BlockOwnerDeletion: &t,
					},
				},
			},
			Spec: v2alpha1.ComponentSpec{
				Image:            rasSpec.Image,
				ImagePullPolicy:  rasSpec.ImagePullPolicy,
				Storage:          rasSpec.Storage,
				StorageClassName: rasSpec.StorageClassName,
				NodeSelector:     rasSpec.NodeSelector,
			},
		}
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaRasComponents(p.Namespace).Create(ctx, ras, v1.CreateOptions{})
		c.Ras = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to create ras: %v", err))
		}
		return controller.NewRequeueAfter(5 * time.Second)
	}
	// Check changes, ready, requeue
	// Coming back to the changes from here
	if len(raslist) > 1 {
		return controller.NewPermanentError(fmt.Errorf("too many ras's defined!"))
	}
	ras := ecosystem.Ras(raslist[0], c.GalasaEcosystemClientSet)
	if !ras.IsReady(ctx) {
		return controller.NewRequeueAfter(time.Second * 5)
	}
	if ras.HasChanged(rasSpec) {
		rasUpdate := raslist[0]
		rasUpdate.Spec.Image = rasSpec.Image
		rasUpdate.Spec.ImagePullPolicy = rasSpec.ImagePullPolicy
		rasUpdate.Spec.Storage = rasSpec.Storage
		rasUpdate.Spec.StorageClassName = rasSpec.StorageClassName
		rasUpdate.Spec.NodeSelector = rasSpec.NodeSelector
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaRasComponents(p.Namespace).Update(ctx, rasUpdate, v1.UpdateOptions{})
		c.Ras = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to update ras: %v", err))
		}
	}
	return nil
}

func (c *Reconciler) ManageApi(ctx context.Context, p *v2alpha1.GalasaEcosystem, selector labels.Selector) error {
	logger := logging.FromContext(ctx)
	apilist, err := c.GalasaAPILister.List(selector)
	if err != nil {
		return controller.NewPermanentError(fmt.Errorf("failed to retrieve api: %v", err))
	}

	apiSpec := p.Spec.ComponentsSpec["api"]
	if len(apilist) == 0 {
		logger.Info("Creating API")
		// Create API CRD
		t := true
		api := &v2alpha1.GalasaApiComponent{
			ObjectMeta: v1.ObjectMeta{
				Name:      "api-" + p.Name,
				Namespace: p.Namespace,
				Labels: map[string]string{
					"galasa-ecosystem-name": p.Name,
				},
				OwnerReferences: []v1.OwnerReference{
					{
						APIVersion:         "GalasaEcosystem",
						Kind:               "galasa.dev/v2alpha1",
						Name:               p.Name,
						UID:                p.GetUID(),
						Controller:         &t,
						BlockOwnerDeletion: &t,
					},
				},
			},
			Spec: v2alpha1.ComponentSpec{
				Image:            apiSpec.Image,
				ImagePullPolicy:  apiSpec.ImagePullPolicy,
				Storage:          apiSpec.Storage,
				StorageClassName: apiSpec.StorageClassName,
				NodeSelector:     apiSpec.NodeSelector,
				ComponentParms: map[string]string{
					"busyboxImage": p.Spec.BusyboxImage,
					"hostname":     p.Spec.Hostname,
					"cpsuri":       c.Cps.Status.StatusParms["cpsuri"],
				},
			},
		}
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaApiComponents(p.Namespace).Create(ctx, api, v1.CreateOptions{})
		c.Api = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to create api: %v", err))
		}
		return controller.NewRequeueAfter(time.Second * 5)
	}
	// Check changes, ready, requeue
	// Coming back to the changes from here
	if len(apilist) > 1 {
		return controller.NewPermanentError(fmt.Errorf("too many api's defined!"))
	}
	api := ecosystem.Api(apilist[0], c.GalasaEcosystemClientSet)
	if !api.IsReady(ctx) {
		return controller.NewRequeueAfter(time.Second * 5)
	}
	if api.HasChanged(apiSpec) {
		apiUpdate := apilist[0]
		apiUpdate.Spec.Image = apiSpec.Image
		apiUpdate.Spec.ImagePullPolicy = apiSpec.ImagePullPolicy
		apiUpdate.Spec.Storage = apiSpec.Storage
		apiUpdate.Spec.StorageClassName = apiSpec.StorageClassName
		apiUpdate.Spec.NodeSelector = apiSpec.NodeSelector
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaApiComponents(p.Namespace).Update(ctx, apiUpdate, v1.UpdateOptions{})
		c.Api = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to update api: %v", err))
		}
	}
	return nil
}

func (c *Reconciler) ManageMetrics(ctx context.Context, p *v2alpha1.GalasaEcosystem, selector labels.Selector) error {
	logger := logging.FromContext(ctx)
	metricslist, err := c.GalasaMetricsLister.List(selector)
	if err != nil {
		return controller.NewPermanentError(fmt.Errorf("failed to retrieve metrics: %v", err))
	}

	metricsSpec := p.Spec.ComponentsSpec["metrics"]
	logger.Infof("spec: %v", p.Spec)
	if len(metricslist) == 0 {
		// Create Metrics CRD
		t := true
		metrics := &v2alpha1.GalasaMetricsComponent{
			ObjectMeta: v1.ObjectMeta{
				Name:      "metrics-" + p.Name,
				Namespace: p.Namespace,
				Labels: map[string]string{
					"galasa-ecosystem-name": p.Name,
				},
				OwnerReferences: []v1.OwnerReference{
					{
						APIVersion:         "GalasaEcosystem",
						Kind:               "galasa.dev/v2alpha1",
						Name:               p.Name,
						UID:                p.GetUID(),
						Controller:         &t,
						BlockOwnerDeletion: &t,
					},
				},
			},
			Spec: v2alpha1.ComponentSpec{
				Image:           metricsSpec.Image,
				ImagePullPolicy: metricsSpec.ImagePullPolicy,
				NodeSelector:    metricsSpec.NodeSelector,
			},
		}
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaMetricsComponents(p.Namespace).Create(ctx, metrics, v1.CreateOptions{})
		c.Metrics = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to create metrics: %v", err))
		}
		return nil
	}
	// Check changes, ready, requeue
	// Coming back to the changes from here
	if len(metricslist) > 1 {
		return controller.NewPermanentError(fmt.Errorf("too many metrics's defined!"))
	}
	metrics := ecosystem.Metrics(metricslist[0], c.GalasaEcosystemClientSet)
	if !metrics.IsReady(ctx) {
		return nil
	}
	if metrics.HasChanged(metricsSpec) {
		metricsUpdate := metricslist[0]
		metricsUpdate.Spec.Image = metricsSpec.Image
		metricsUpdate.Spec.ImagePullPolicy = metricsSpec.ImagePullPolicy
		metricsUpdate.Spec.NodeSelector = metricsSpec.NodeSelector
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaMetricsComponents(p.Namespace).Update(ctx, metricsUpdate, v1.UpdateOptions{})
		c.Metrics = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to update metrics: %v", err))
		}
	}
	return nil
}

func (c *Reconciler) ManageResmon(ctx context.Context, p *v2alpha1.GalasaEcosystem, selector labels.Selector) error {
	logger := logging.FromContext(ctx)
	resmonlist, err := c.GalasaResmonLister.List(selector)
	if err != nil {
		return controller.NewPermanentError(fmt.Errorf("failed to retrieve resmon: %v", err))
	}

	resmonSpec := p.Spec.ComponentsSpec["resmon"]
	logger.Infof("spec: %v", p.Spec)
	if len(resmonlist) == 0 {
		// Create Resmon CRD
		t := true
		resmon := &v2alpha1.GalasaResmonComponent{
			ObjectMeta: v1.ObjectMeta{
				Name:      "resmon-" + p.Name,
				Namespace: p.Namespace,
				Labels: map[string]string{
					"galasa-ecosystem-name": p.Name,
				},
				OwnerReferences: []v1.OwnerReference{
					{
						APIVersion:         "GalasaEcosystem",
						Kind:               "galasa.dev/v2alpha1",
						Name:               p.Name,
						UID:                p.GetUID(),
						Controller:         &t,
						BlockOwnerDeletion: &t,
					},
				},
			},
			Spec: v2alpha1.ComponentSpec{
				Image:           resmonSpec.Image,
				ImagePullPolicy: resmonSpec.ImagePullPolicy,
				NodeSelector:    resmonSpec.NodeSelector,
			},
		}
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaResmonComponents(p.Namespace).Create(ctx, resmon, v1.CreateOptions{})
		c.Resmon = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to create resmon: %v", err))
		}
		return nil
	}
	// Check changes, ready, requeue
	// Coming back to the changes from here
	if len(resmonlist) > 1 {
		return controller.NewPermanentError(fmt.Errorf("too many resmon's defined!"))
	}
	resmon := ecosystem.Resmon(resmonlist[0], c.GalasaEcosystemClientSet)
	if !resmon.IsReady(ctx) {
		return nil
	}
	if resmon.HasChanged(resmonSpec) {
		resmonUpdate := resmonlist[0]
		resmonUpdate.Spec.Image = resmonSpec.Image
		resmonUpdate.Spec.ImagePullPolicy = resmonSpec.ImagePullPolicy
		resmonUpdate.Spec.NodeSelector = resmonSpec.NodeSelector
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaResmonComponents(p.Namespace).Update(ctx, resmonUpdate, v1.UpdateOptions{})
		c.Resmon = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to update resmon: %v", err))
		}
	}
	return nil
}

func (c *Reconciler) ManageEngineController(ctx context.Context, p *v2alpha1.GalasaEcosystem, selector labels.Selector) error {
	logger := logging.FromContext(ctx)
	enginecontrollerlist, err := c.GalasaEngineControllerLister.List(selector)
	if err != nil {
		return controller.NewPermanentError(fmt.Errorf("failed to retrieve enginecontroller: %v", err))
	}

	enginecontrollerSpec := p.Spec.ComponentsSpec["enginecontroller"]
	logger.Infof("spec: %v", p.Spec)
	if len(enginecontrollerlist) == 0 {
		// Create EngineController CRD
		t := true
		enginecontroller := &v2alpha1.GalasaEngineControllerComponent{
			ObjectMeta: v1.ObjectMeta{
				Name:      "enginecontroller-" + p.Name,
				Namespace: p.Namespace,
				Labels: map[string]string{
					"galasa-ecosystem-name": p.Name,
				},
				OwnerReferences: []v1.OwnerReference{
					{
						APIVersion:         "GalasaEcosystem",
						Kind:               "galasa.dev/v2alpha1",
						Name:               p.Name,
						UID:                p.GetUID(),
						Controller:         &t,
						BlockOwnerDeletion: &t,
					},
				},
			},
			Spec: v2alpha1.ComponentSpec{
				Image:           enginecontrollerSpec.Image,
				ImagePullPolicy: enginecontrollerSpec.ImagePullPolicy,
				NodeSelector:    enginecontrollerSpec.NodeSelector,
			},
		}
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaEngineControllerComponents(p.Namespace).Create(ctx, enginecontroller, v1.CreateOptions{})
		c.EngineController = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to create enginecontroller: %v", err))
		}
		return nil
	}
	// Check changes, ready, requeue
	// Coming back to the changes from here
	if len(enginecontrollerlist) > 1 {
		return controller.NewPermanentError(fmt.Errorf("too many enginecontroller's defined!"))
	}
	enginecontroller := ecosystem.EngineController(enginecontrollerlist[0], c.GalasaEcosystemClientSet)
	if !enginecontroller.IsReady(ctx) {
		return nil
	}
	if enginecontroller.HasChanged(enginecontrollerSpec) {
		enginecontrollerUpdate := enginecontrollerlist[0]
		enginecontrollerUpdate.Spec.Image = enginecontrollerSpec.Image
		enginecontrollerUpdate.Spec.ImagePullPolicy = enginecontrollerSpec.ImagePullPolicy
		enginecontrollerUpdate.Spec.NodeSelector = enginecontrollerSpec.NodeSelector
		i, err := c.GalasaEcosystemClientSet.GalasaV2alpha1().GalasaEngineControllerComponents(p.Namespace).Update(ctx, enginecontrollerUpdate, v1.UpdateOptions{})
		c.EngineController = i
		if err != nil {
			return controller.NewPermanentError(fmt.Errorf("failed to update enginecontroller: %v", err))
		}
	}
	return nil
}

func mustNewRequirement(key string, op selection.Operator, vals []string) labels.Requirement {
	r, err := labels.NewRequirement(key, op, vals)
	if err != nil {
		panic(fmt.Sprintf("mustNewRequirement(%v, %v, %v) = %v", key, op, vals, err))
	}
	return *r
}
