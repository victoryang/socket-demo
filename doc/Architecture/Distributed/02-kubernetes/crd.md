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

*Custom resources* are extensions of the Kubernetes API. This page discusses when to add a custom resource to your Kubernetes cluster and when to use a standalone service. It describes thw two methods for adding custom resources and how to choose between them.

### Custom resources

A *resource* is an endpoint in the Kubernetes API that stores a collection of API objects of a certain kind; for example, the built-in *pods* resource contains a collection of Pod objects.

A *custom resource* is an extension of the Kubernetes API that is not necessarily available in a default Kubernetes installation. It represents a customization of a particular Kubernetes installation. However, many core Kubernetes functions are now built using custom resources, making Kubernetes more modular.

Custom resources can appear and disappear in a running cluster through dynamic registration, and cluster admins can update custom resources independently of the cluster itself. Once a custom resource is installed, users can create and access its object using kubectl, just as they do for built-in resources like *Pods*.

### Custom controllers

On their own, custom resources simply let you store and retrieve structure data. When you combine a custom resource with a *custom controller*, custom resources provide a true

## Extend the Kubernetes API with CustomResourceDefinitions

