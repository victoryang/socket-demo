# Best Practise

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