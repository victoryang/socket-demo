# Authentication

```
completedOptions, err := Complete(s) {
    s.Authentication.ApplyAuthorization(s.Authorization)
}
```

Run
```
CreateServerChain

CreateKubeAPIServerConfig {
    s.Authentication.ApplyTo

    BuildAuthorizer

    s.Admission.ApplyTo
}

CreateKubeAPIServer {
    kubeAPIServerConfig.Complete().New(delegateAPIServer)
}

kubeAPIServerConfig.Complete().New {
    m := &Instance{
        GenericAPIServer:          s,
        ClusterAuthenticationInfo: c.ExtraConfig.ClusterAuthenticationInfo,
    }

    restStorageProviders := []RESTStorageProvider {
        ...
        authenticationrest.RESTStorageProvider,
        authorizationrest.RESTStorageProvider,
        ...
        rbacrest.RESTStorageProvider,
        ...
    }
}
```

---
m.GenericAPIServer.AddPostStartHookOrDie

start-cluster-authentication-info-controller
```
// NewClusterAuthenticationTrustController returns a controller that will maintain the kube-system configmap/extension-apiserver-authentication
// that holds information about how to aggregated apiservers are recommended (but not required) to configure themselves. 
clusterauthenticationtrust.NewClusterAuthenticationTrustController
```

---

NewServerRunOptions
```
Admission:               kubeoptions.NewAdmissionOptions(),
Authentication:          kubeoptions.NewBuiltInAuthenticationOptions().WithAll(),
Authorization:           kubeoptions.NewBuiltInAuthorizationOptions(),
```

NewBuiltInAuthenticationOptions().WithAll()
```
// WithAll set default value for every build-in authentication option
func (o *BuiltInAuthenticationOptions) WithAll() *BuiltInAuthenticationOptions {
    return o.
        WithAnonymous().
        WithBootstrapToken().
        WithClientCert().
        WithOIDC().
        WithRequestHeader().
        WithServiceAccounts().
        WithTokenFile().
        WithWebHook()
}
```
