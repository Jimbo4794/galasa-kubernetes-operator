/*
 * Copyright contributors to the Galasa Project
 */
// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/Jimbo4794/galasa-kubernetes-operator/pkg/apis/galasaecosystem/v2alpha1"
	"github.com/Jimbo4794/galasa-kubernetes-operator/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type GalasaV2alpha1Interface interface {
	RESTClient() rest.Interface
	GalasaApiComponentsGetter
	GalasaCpsComponentsGetter
	GalasaEcosystemsGetter
	GalasaEngineControllerComponentsGetter
	GalasaMetricsComponentsGetter
	GalasaRasComponentsGetter
	GalasaResmonComponentsGetter
	GalasaToolboxComponentsGetter
}

// GalasaV2alpha1Client is used to interact with features provided by the galasa.dev group.
type GalasaV2alpha1Client struct {
	restClient rest.Interface
}

func (c *GalasaV2alpha1Client) GalasaApiComponents(namespace string) GalasaApiComponentInterface {
	return newGalasaApiComponents(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaCpsComponents(namespace string) GalasaCpsComponentInterface {
	return newGalasaCpsComponents(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaEcosystems(namespace string) GalasaEcosystemInterface {
	return newGalasaEcosystems(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaEngineControllerComponents(namespace string) GalasaEngineControllerComponentInterface {
	return newGalasaEngineControllerComponents(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaMetricsComponents(namespace string) GalasaMetricsComponentInterface {
	return newGalasaMetricsComponents(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaRasComponents(namespace string) GalasaRasComponentInterface {
	return newGalasaRasComponents(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaResmonComponents(namespace string) GalasaResmonComponentInterface {
	return newGalasaResmonComponents(c, namespace)
}

func (c *GalasaV2alpha1Client) GalasaToolboxComponents(namespace string) GalasaToolboxComponentInterface {
	return newGalasaToolboxComponents(c, namespace)
}

// NewForConfig creates a new GalasaV2alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*GalasaV2alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &GalasaV2alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new GalasaV2alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GalasaV2alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GalasaV2alpha1Client for the given RESTClient.
func New(c rest.Interface) *GalasaV2alpha1Client {
	return &GalasaV2alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GalasaV2alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
