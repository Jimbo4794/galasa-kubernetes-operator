package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/cps"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/enginecontroller"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/galasaecosystem"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/metrics"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/ras"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/resmon"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/reconciler/toolbox"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"
)

const (
	// ControllerLogKey is the name of the logger for the controller cmd
	ControllerLogKey = "galasa-ecosystem-operator"
)

var (
	namespace = flag.String("namespace", corev1.NamespaceAll, "Namespace to restrict informer to. Optional, defaults to all namespaces.")
)

func main() {
	cfg := injection.ParseAndGetRESTConfigOrDie()
	ctx := injection.WithNamespaceScope(signals.NewContext(), *namespace)

	// Set up liveness and readiness probes.
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/health", handler)
	mux.HandleFunc("/readiness", handler)

	port := os.Getenv("PROBES_PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		log.Printf("Readiness and health check server listening on port %s", port)
		log.Fatal(http.ListenAndServe(":"+port, mux))
	}()

	sharedmain.MainWithConfig(ctx, ControllerLogKey, cfg,
		galasaecosystem.NewController(*namespace),
		cps.NewController(*namespace),
		ras.NewController(*namespace),
		enginecontroller.NewController(*namespace),
		metrics.NewController(*namespace),
		resmon.NewController(*namespace),
		toolbox.NewController(*namespace),
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
