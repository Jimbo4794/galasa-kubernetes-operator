package galasaecosystem

import (
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1/api"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1/cps"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1/enginecontroller"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1/metrics"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1/ras"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1/resmon"
)

func Cps(cpsCrd *v2alpha1.GalasaCpsComponent) v2alpha1.ComponentInterface {
	return cps.New(cpsCrd)
}

func Ras(rasCrd *v2alpha1.GalasaRasComponent) v2alpha1.ComponentInterface {
	return ras.New(rasCrd)
}

func Api(apiCrd *v2alpha1.GalasaApiComponent) v2alpha1.ComponentInterface {
	return api.New(apiCrd)
}

func Metrics(metricsCrd *v2alpha1.GalasaMetricsComponent) v2alpha1.ComponentInterface {
	return metrics.New(metricsCrd)
}

func Resmon(resmonCrd *v2alpha1.GalasaResmonComponent) v2alpha1.ComponentInterface {
	return resmon.New(resmonCrd)
}

func EngineController(engineControllerCrd *v2alpha1.GalasaEngineControllerComponent) v2alpha1.ComponentInterface {
	return enginecontroller.New(engineControllerCrd)
}

// func Toolbox(cpsCrd *v2alpha1.GalasaCpsComponent) v2alpha1.CpsInterface {
// 	return cps.New(cpsCrd)
// }
