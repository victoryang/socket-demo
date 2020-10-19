# CRD

[crd share part 1](https://blog.csdn.net/boling_cavalry/article/details/88917818)

[crd share part 2](https://blog.csdn.net/boling_cavalry/article/details/88924194)

[crd share part 3](https://blog.csdn.net/boling_cavalry/article/details/88934063)

[crd vs operator](https://zhuanlan.zhihu.com/p/141877334)

[java share](https://github.com/zq2599/blog_demos)

[kuberentes API Overview](https://kubernetes.io/docs/reference/using-api/api-overview/)

## CRD

third party resources

https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/thirdpartyresources.md

```
/apis/<group>/<version>/namespaces/<namespace-name>/<plural-resource-type>
```

## Custom Resources

*Custom resources* are extensions of the Kubernetes API. This page discusses when to add a custom resource to your Kubernetes cluster and when to use a standalone service. It describes the two methods for adding custom resources and how to choose between them.

### Custom resources

A *resource* is an endpoint in the Kubernetes API that stores a collection of API objects of a certain kind; for example, the built-in *pods* resource contains a collection of Pod objects.

A *custom resource* is an extension of the Kubernetes API that is not necessarily available in a default Kubernetes installation. It represents a customization of a particular Kubernetes installation. However, many core Kubernetes functions are now built using custom resources, making Kubernetes more modular.

Custom resources can appear and disappear in a running cluster through dynamic registration, and cluster admins can update custom resources independently of the cluster itself. Once a custom resource is installed, users can create and access its object using kubectl, just as they do for built-in resources like *Pods*.

### Custom controllers

On their own, custom resources simply let you store and retrieve structure data. When you combine a custom resource with a *custom controller*, custom resources provide a true
declarative *API*.

A declarative API allow you to declare or specify the desired state of your resource and tries to keep the current state of Kubernetes objects in sync with the desired state. The controller interprets the structure data as a record of the user's desired state, and continually maintains this state.

You can deploy and update a custom controller on a running cluster, independently of the cluster's lifecycle. Custom controllers can work with any kind of resource, but they are especially effective when combined with custom resources. The Operator pattern combines custom resources and custom controllers. You can use custom controllers to encode domain knowledge for specific applications into an extension of the Kubernetes API.

### Adding custom resources

Kubernetes provides two ways to add custom resources to your cluster:

- CRDs are simple and can be created without any programming.
- API Aggregation requires programming, but allow more control over API behaviors like how data is stored and conversion between API versions.

Kubernetes provides these two options to meet the needs of different users, so that neither ease of use nor flexibility is compromised.

Aggregated APIs are subordinate API servers that sit behind the primary API server, which acts as a proxy. This arrangement is called API Aggregation. To users, it simply appears that the Kubernetes API is extended.

CRDs allow users to create new types of resources without adding another API server. You don not need to understand API Aggregation to use CRDs.

Regardless of how they are installed, the new resources are referred to as Custom Resources to distinguish them from built-in Kubernetes resources(like pods).




