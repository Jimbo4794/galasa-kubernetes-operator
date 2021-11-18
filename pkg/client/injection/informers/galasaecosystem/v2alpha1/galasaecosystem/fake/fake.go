/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	fake "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/injection/informers/factory/fake"
	galasaecosystem "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/injection/informers/galasaecosystem/v2alpha1/galasaecosystem"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = galasaecosystem.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Galasa().V2alpha1().GalasaEcosystems()
	return context.WithValue(ctx, galasaecosystem.Key{}, inf), inf.Informer()
}
