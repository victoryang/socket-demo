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

### kubeAPIServer



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