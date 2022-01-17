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

staging/k8s.io/apiserver/pkg/server/genericapiserver.go
m.InstallAPIGroups
```
for _,apiGroupInfo :=range apiGroupInfos {
    s.installAPIResources()
}
```

installAPIResources is a private method for installing the REST storage backing each api groupversionresource
```
apiGroupVersion := s.getAPIGroupVersion {
    s.newAPIGroupVersion(apiGroupInfo, groupVersion) {
        return &genericapi.APIGroupVersion{
            GroupVersion:     groupVersion,
            MetaGroupVersion: apiGroupInfo.MetaGroupVersion,
            ParameterCodec:        apiGroupInfo.ParameterCodec,
            Serializer:            apiGroupInfo.NegotiatedSerializer,
            Creater:               apiGroupInfo.Scheme,
            Convertor:             apiGroupInfo.Scheme,
            ConvertabilityChecker: apiGroupInfo.Scheme,
            UnsafeConvertor:       runtime.UnsafeObjectConvertor(apiGroupInfo.Scheme),
            Defaulter:             apiGroupInfo.Scheme,
            Typer:                 apiGroupInfo.Scheme,
            Linker:                runtime.SelfLinker(meta.NewAccessor()),

            EquivalentResourceRegistry: s.EquivalentResourceRegistry,

            Admit:             s.admissionControl,
            MinRequestTimeout: s.minRequestTimeout,
            Authorizer:        s.Authorizer,
        }
    }
}
r, err := apiGroupVersion.InstallREST(s.Handler.GoRestfulContainer)

resourceInfos = append(resourceInfos, r...)

s.StorageVersionManager.AddResourceInfo(resourceInfos...)
```

InstallREST registers the REST handlers (storage, watch, proxy and redirect) into a restful Container

```
prefix := path.Join(g.Root, g.GroupVersion.Group, g.GroupVersion.Version)
installer := &APIInstaller{
    group:             g,
    prefix:            prefix,
    minRequestTimeout: g.MinRequestTimeout,
}

apiResources, resourceInfos, ws, registrationErrors := installer.Install()
versionDiscoveryHandler := discovery.NewAPIVersionHandler(g.Serializer, g.GroupVersion, staticLister{apiResources})
versionDiscoveryHandler.AddToWebService(ws)
container.Add(ws)
```

----
**staging/src/k8s.io/apiserver/pkg/endpoints/installer.go**

installer.Install
```
// Install handlers for API resources
ws := a.newWebService()

for _,path :=range paths {
    apiResource, resourceInfo, err := a.registerResourceHandlers(path, a.group.Storage[path], ws)
}
```

**installer.registerResourceHandlers**
```
// register handler to restful webservice for all kinds of resources
```

e.g. create handler

```bash
case "POST": // Create a resource.
handler = restfulCreateResource(creater, reqScope, admit) {
    handlers.CreateResource(r, &scope, admit)(res.ResponseWriter, req.Request)
}
```

// createHandler
```bash
// enforce a timeout of at most requestTimeoutUpperBound (34s) or less if the user-provided
// timeout inside the parent context is lower than requestTimeoutUpperBound.
ctx, cancel := context.WithTimeout(req.Context(), requestTimeoutUpperBound)

gv := scope.Kind.GroupVersion()

decoder := scope.Serializer.DecoderToVersion(s.Serializer, scope.HubGroupVersion)

requestFunc := func() (runtime.Object, error) {
    return r.Create(
        ctx,
        name,
        obj,
        rest.AdmissionToValidateObjectFunc(admit, admissionAttributes, scope),
        options,
    )
}

result, err := requestFunc()
```
-----
**database operation**

staging/src/k8s.io/apiserver/pkg/registry/generic/store.go
```
(e *Store) Create
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

```
// PrepareRun prepares the aggregator to run, by setting up the OpenAPI spec and calling the generic PrepareRun
prepared, err := server.PrepareRun()

prepared.Run(stopCh)
```

staging/src/k8s.io/kube-aggregator/pkg/apiserver/apiserver.go
server.PrepareRun
```
prepared := s.GenericAPIServer.PrepareRun()

preparedAPIAggregator{APIAggregator: s, runnable: prepared}
```

staging/src/k8s.io/apiserver/pkg/server/genericapiserver.go
prepared.Run
```
s.NonBlockingRun
```
