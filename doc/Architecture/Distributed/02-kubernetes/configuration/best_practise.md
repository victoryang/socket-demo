# Best Practise

## General Configuration Tips

- When defining configurations, specify the latest stable API version
- Configuration files should be stored in version control before being pushed to the cluster. This allows you to quickly roll back a configuration change if necessary. It also aids cluster recreation and restoration.
- Write your configuration files using YAML rather than JSON. Though these formats can be used interchangeably in almost all scenarios , YAML tends to be more user-friendly.
- Group related objects into a single file whenever it makes sense. Once file is often easier to manage than serveral.
- Note also that many `kubectl` commands can be called on a directory. For example, you can call `kubectl apply` on a directory of config files.
- Don't specify default values unnecessarily: simple, minimal configuration will make errors less likely.
- Put object description in annotations, to allow better introspection.

## "Naked" Pods versus ReplicaSets, Deployment, and Jobs

Don't use naked Pods(that is, Pods not bound to a ReplicaSet or Deployment) if you can avoid it. Naked Pods will not be rescheduled in the event of a node failure.

A Deployment, which both creates a ReplicaSet to ensure that the desired number of Pods is always available, and specifies a strategy to replace Pods, is almost always preferable to creating Pods directly, except for some explicit `restartPolicy: Never` scenarios. A Job may also be appropriate.

## Services

Create a Service before its corresponding backend workloads(Deployments or ReplicaSets), and before any workloads that need to access it. When Kubernetes starts a container, it provides environment variables pointing to all the Services which were running when the container was started. For example, if a Service named `foo` exists, all containers will get the following variables in their initial environment

```bash
FOO_SERVICE_HOST=<the host the Service is running on>
FOO_SERVICE_PORT=<the port the Service is running on>
```

*This does imply an ordering requirement* - any `Service` that a `Pod` wants to access must be created before the `Pod` itself, or else the environment variable will not be populated. DNS does not have this restriction.

An optional(though strongly recommended) cluster add-on is a DNS server. The DNS server watches the Kubernetes API for new `Services` and creates a set of DNS records for each. If DNS has been enable throughout the cluster then all `Pods` should be able to do name resolution of `Service` automatically.

Don't specify a `hostPort` for a Pod unless it is absolutely necessary. When you bind a Pod to a `hostPort`, it limits the number of places the Pod can be scheduled, because each <`hostIP`,`hostPort`,`protocol`> combination must be unique. If you don't specify the `hostIP` and `protocol` explicitly, Kubernetes will use `0.0.0.0` as the default `hostIP` and `TCP` as the default `protocol`.

If you only need access to the port for defbugging purposes, you can use the apiserver proxy or kubectl port-forward.

If you explicitly need to expose a Pod's port on the node, consider using a NodePort Service before resorting to `hostPort`.

Avoid using `hostNetwork`, for the same reasons as `hostPort`.

Use headless Services(which have a `ClusterIP` of `None`) for easy service discovery when you don't need `kube-proxy` load balancing.

## Using Labels

Define and use labels that identify semantic attributes of your application or Deployment, such as {app: myapp, tier: frontend, phase: test, deployment: v3}. You can use these labels to select the appropiate Pods for other resources; for example, a Service that selects all `tier: frontend` Pods, or all `phase: test` components of `app:myapp`.

A Service can be made to span multiple Deployments by omitting release-specific labels from its selector. When you need to update a running service without downtime, use a Deployment.

A desired state of an object is described by a Deployment, and if changes to that spec are *applied*, the deployment controller changes the actual state to the desired state at a controlled rate.

- Use the Kubernetes common labels for common use cases. These standardized labels enrich the metadata in a way that allow tools, including `kubectl` and dashboard, to work in an interoperable way.

- You can manipulate labels for debugging. Because Kubernetes controllers (such as ReplicaSet) and Services match to Pods using selector lables, removing the relevant lables from a Pod will stop it from being considered by a controller or from being served traffic by a Service. If you remove the labels of an existing Pod, its controller will create a new Pod to take its place. This is a useful way to debug a previously "live" Pod in a "quarantine" environment. To interactively remove or add labels, use `kubectl lable`.

## Container Images

The `imagePullPolicy` and the tag of the image affect when the kubelet attempts to pull the specified image.

- `imagePullPolicy: IfNotPresent`: the image is pulled only if it is not already present locally.

- `imagePullPolicy: Always`: every time the kubelet launches a container, the kubelet queries the container image registry to resolve the name to an image digest. If the kubelet has a container image with that exact digest cached locally, the kubelet uses its cached image; otherwise, the kubelet downloads(pulls) the image with the resolved digest, and uses that image to launch the container.

- `imagePullPolicy` is omitted and either the image tag is `:latest` or it is omitted: `Always` is applied.

- `imagePullPolicy` is omitted and the image tag is present but not `:latest`: `IfNotPresent` is applied.

- `imagePullPolicy: Never`: the image is assumed to exist locally. No attempt is made to pull the image.

## Using kubectl

- Use `kubectl apply -f <directory>`. This looks for Kubernetes configuration in all `.yaml`, `.yml` and `.json` files in `<directory>` and passes it to `apply`.

- Use label selectors for `get` and `delete` operations instead of specific object names.

- Use `kubectl create deployment` and `kubectl expose` to quickly create single-container Deployments and Services.