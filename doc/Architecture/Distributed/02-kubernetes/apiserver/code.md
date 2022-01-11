# Code

## cmd/kube-apiserver/app/server.go

Run
```
server, err := CreateServerChain(completeOptions, stopCh)

prepared, err := server.PrepareRun()

prepared.Run(stopCh)
```

CreateServerChain
- apiExtensionServer
- apiServer
- aggregatorServer

```
kubeAPIServerConfig, serviceResolver, pluginInitializer, err := CreateKubeAPIServerConfig(completedOptions)

apiExtensionsConfig, err := createAPIExtensionsConfig

apiExtensionsServer, err := createAPIExtensionsServer

kubeAPIServer, err := CreateKubeAPIServer

aggregatorConfig, err := createAggregatorConfig
aggregatorServer, err := createAggregatorServer
```

### apiExtensionServer

createAPIExtensionsConfig

```
// override gernericConfig.AdmissionControl with apiextensions' scheme,
commandOptions.Admission.ApplyTo()

// copy the etcd options so we don't mutate originals.
etcdOptions := *commandOptions.Etcd

// override MergedResourceConfig with apiextensions defaults and registry
commandOptions.APIEnablement.ApplyTo()

apiextensionsConfig := &apiextensionsapiserver.Config{
    GenericConfig: &genericapiserver.RecommendedConfig{
        Config:                genericConfig,
        SharedInformerFactory: externalInformers,
    },
    ExtraConfig: apiextensionsapiserver.ExtraConfig{
        CRDRESTOptionsGetter: apiextensionsoptions.NewCRDRESTOptionsGetter(etcdOptions),
        MasterCount:          masterCount,
        AuthResolverWrapper:  authResolverWrapper,
        ServiceResolver:      serviceResolver,
    },
}
```

k8s.io/apiextensions-apiserver/pkg/apiserver/apiserver.go

```
// createAPIExtensionsServer
apiextensionsConfig.Complete().New(delegateAPIServer)
```

New()
```
s := &CustomResourceDefinitions{
    GenericAPIServer: genericServer,
}

apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo

customResourceDefinitionStorage, err := customresourcedefinition.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter)

crdHandler, err := NewCustomResourceDefinitionHandler

// add hook
// start-apiextensions-informers
// start-apiextensions-controllers
// crd-informer-synced

s.GenericAPIServer.InstallAPIGroup
```

### kubeAPIServer

```
// CreateKubeAPIServer
kubeAPIServer, err := kubeAPIServerConfig.Complete().New(delegateAPIServer)
```

pkg/controlplane/instance.go

config.Complete
```
serviceIPRange, apiServerServiceIP, err := ServiceIPRange(cfg.ExtraConfig.ServiceIPRange)

// Instance contains state for a Kubernetes cluster api server instance.
type Instance struct {
	GenericAPIServer *genericapiserver.GenericAPIServer

	ClusterAuthenticationInfo clusterauthenticationtrust.ClusterAuthenticationInfo
}

```

config.New
```
// New returns a new instance of Master from the given config.
// Certain config fields will be set to a default value if unset
// Certain config fields must be specified, including: KubeletClientConfig

// apiserver/pkg/server/config.go
// return &GenericAPIServer{}
s, err := c.GenericConfig.New("kube-apiserver", delegationTarget)

md, err := serviceaccount.NewOpenIDMetadata(
    c.ExtraConfig.ServiceAccountIssuerURL,
    c.ExtraConfig.ServiceAccountJWKSURI,
    c.GenericConfig.ExternalAddress,
    c.ExtraConfig.ServiceAccountPublicKeys,
)

routes.NewOpenIDMetadataServer(md.ConfigJSON, md.PublicKeysetJSON).
			Install(s.Handler.GoRestfulContainer)
```

RESTStorageProvider
```
// RESTStorageProvider is a factory type for REST storage.
type RESTStorageProvider interface {
	GroupName() string
	NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool, error)
}
```

```
// The order here is preserved in discovery.
// If resources with identical names exist in more than one of these groups (e.g. "deployments.apps"" and "deployments.extensions"),
// the order of this list determines which group an unqualified resource name (e.g. "deployments") should prefer.
// This priority order is used for local discovery, but it ends up aggregated in `k8s.io/kubernetes/cmd/kube-apiserver/app/aggregator.go
// with specific priorities.
// TODO: describe the priority all the way down in the RESTStorageProviders and plumb it back through the various discovery
// handlers that we have.
restStorageProviders := []RESTStorageProvider{
    apiserverinternalrest.StorageProvider{},
    authenticationrest.RESTStorageProvider{Authenticator: c.GenericConfig.Authentication.Authenticator, APIAudiences: c.GenericConfig.Authentication.APIAudiences},
    authorizationrest.RESTStorageProvider{Authorizer: c.GenericConfig.Authorization.Authorizer, RuleResolver: c.GenericConfig.RuleResolver},
    autoscalingrest.RESTStorageProvider{},
    batchrest.RESTStorageProvider{},
    certificatesrest.RESTStorageProvider{},
    coordinationrest.RESTStorageProvider{},
    discoveryrest.StorageProvider{},
    networkingrest.RESTStorageProvider{},
    noderest.RESTStorageProvider{},
    policyrest.RESTStorageProvider{},
    rbacrest.RESTStorageProvider{Authorizer: c.GenericConfig.Authorization.Authorizer},
    schedulingrest.RESTStorageProvider{},
    storagerest.RESTStorageProvider{},
    flowcontrolrest.RESTStorageProvider{},
    // keep apps after extensions so legacy clients resolve the extensions versions of shared resource names.
    // See https://github.com/kubernetes/kubernetes/issues/42392
    appsrest.StorageProvider{},
    admissionregistrationrest.RESTStorageProvider{},
    eventsrest.RESTStorageProvider{TTL: c.ExtraConfig.EventTTL},
}

m.InstallAPIs(restStorageProviders...)

// add hook
// start-cluster-authentication-info-controller 
// start-kube-apiserver-identity-lease-controller 
// start-kube-apiserver-identity-lease-garbage-collector

m.GenericAPIServer.AddPostStartHookOrDie
```

m.InstallAPIs
```
// Exposes given api groups in the API.
m.GenericAPIServer.InstallAPIGroups(apiGroupsInfo...)
```

### aggregatorServer

```
commandOptions.Admission.ApplyTo()

etcdOptions := *commandOptions.Etcd

commandOptions.APIEnablement.ApplyTo()

aggregatorConfig := &aggregatorapiserver.Config{
    GenericConfig: &genericapiserver.RecommendedConfig{
        Config:                genericConfig,
        SharedInformerFactory: externalInformers,
    },
    ExtraConfig: aggregatorapiserver.ExtraConfig{
        ProxyClientCertFile: commandOptions.ProxyClientCertFile,
        ProxyClientKeyFile:  commandOptions.ProxyClientKeyFile,
        ServiceResolver:     serviceResolver,
        ProxyTransport:      proxyTransport,
    },
}
```

