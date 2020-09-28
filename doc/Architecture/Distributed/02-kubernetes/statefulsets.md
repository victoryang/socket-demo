# StatefulSets

StatefulSet is the workload API object used to manage stateful applications.

Manages the deployment and scaling of a set of Pods, *and provides guarantees about the ordering and uniqueness* of these Pods.

Like a Deployment, a StatefulSet manages Pods that are based on an identical container spec.
Unlike a Deployment, a StatefulSet maintains a sticky identity for each of their Pods. These pods are created from the same spec, but are not interchangeable: each has a persistent identifier that it maintains across any rescheduling.

If you want to use storage volumes to provide persistence for your workload, you can use a StatefulSet as part of the solution. Although individual Pods in a StatefulSet are susceptible to failure, the persistent Pod identifiers make it easier to match existing volumes to the new Pods that replace any that have failed.

## Using StatfulSets

StatfulSets are valuable for applications that require one or more of the following.

- Stable, unique network identifies.
- Stable, persistent storage
- Ordered, graceful deployment and scaling
- Ordered, automated rolling updates.

In the above, stable is synonymous with persistence across Pod (re)scheduling. If an application doesn't require any stable identifiers or ordered deployment, deletion, or scaling, you should deploy your application using a workload object that provides a set of stateless replicas. Deployment or ReplicaSet may be better suited to your statless needs.

## Limitations

- The storage for a given Pod must either be provisioned by a PersistentVolume Provisioner based on the requested `storage class`, or pre-provisioned by an admin.
- Deleting and/or scaling a StatefulSet down will not delete the volumes associated with the StatefulSet. This is done to ensure data safety, which is generally more valuable than an automatic purge of all related StatefulSet resources.
- StatefulSets currently require a Headless Service to be responsible for the network identity of the Pods. You are responsible for creating this service.
- StatefulSets do not provide any guarantees on the termination of pods when a StatefulSet is deleted. To achieve ordered and graceful termination of the pods in the StatefulSet, it is possible to scale the StatefulSet down to 0 prior to deletion.
- When using Rolling Updates with the default Pod Management Policy(`OrderedReady`), it's possible to get into a broken state that requires manual intervention to repair.

## Pod Identity

StatefulSet Pods have a unique identity that is comprised of an ordinal, a stable network identity, and stable storage. The identity sticks to the Pod, regardless of which node it's (re)scheduled on.

### Ordinal Index

For a StatefulSet with N replicas, each Pod in the StatefulSet will be assigned an integer ordinal, from 0 up through N-1, that is unique over the set.

### Stable Network ID

Each Pod in a StatefulSet derives its hostname from the name of the StatefulSet and the ordinal of the Pod. The pattern for the constructed hostname is `$(statefulsetname)-$(ordinal)`. The domain managed by this service takes the form: `$(service name).$(namespace).svc.cluster.local`, where "cluster.local" is the cluster domain. As each Pod is created, it gets a matching DNS subdomain, taking the form: `$(podname).$(governing service domain)`, where the governing service is defined by the `serviceName` field on the StatefulSet.

Depending on how DNS is configured in your cluster, you may not be able to look up the DNS name for a newly-run Pod immediately. This behavior can occur when other clients in the cluster have already sent queries for the hostname of the Pod before it was created. Negative caching(normal in DNS) means that the results of previous failed lookups are remembered and reused, event after the Pod is running, for at least a few seconds.

if you need to discover Pods promptly after they are created. you have a few options:

- Query the Kubernetes API directly (for example, using a watch) rather than replying on DNS lookups.
- Decrease the time of caching in your Kubernetes DNS provider (typically this means editing the config map for CoreDNS, which currently caches for 30 seconds).

As mentioned in the limitations section, you are responsible for creating the Headless Service responsible for the network identity of the pods.

Here are some examples of choices for Cluster Domain, Service name, StatefulSet name, and how that affects the DNS names for the StatefulSet's Pods.

|Cluster Domain|Service(ns/name)|StatefulSet(ns/name)|StatefulSet Domain|Pod DNS|Pod Hostname|
|-|-|-|-|-|-|
|cluster.local|default/nginx|default/web|nginx.default.svc.cluster.local|web-{0..N-1}.nginx.default.svc.cluster.local|web-{0..N-1}|
|cluster.local|foo/nginx|foo/web|nginx.foo.svc.cluster.local|web-{0..N-1}.nginx.foo.svc.cluster.local|web-{0..N-1}|
|kube.local|foo/nginx|foo/web|nginx.foo.svc.kube.local|web-{0..N-1}.nginx.foo.svc.kube.local|web-{0..N-1}|

### Stable Storage

Kubernetes creates one PersistentVolume for each VolumeClaimTemplate. If no StorageClass is specified, then the default StorageClass will be used. When a Pod is (re)scheduled onto a node, its `volumeMounts` mount the PersistentVolumes associated with its PersistentVolume Claims. Note that, the PersistentVolume associated with the Pod's PersistentVolume Claim are not deleted when the Pod, or StatefulSet are deleted. This must be done manually.

### Pod Name Label

When the StatefulSet Controller creates a Pod, it adds a label, `statefulset.kubernetes.io/pod-name`, that is set to the name of the Pod. This label allows you to attach a Service to a specific Pod in the StatefulSet.

## Deployment and Scaling Guarantees

- For a StatefulSet with N replicas, when Pods are being deployed, they are created sequentially, in order from {0..N-1}
- When Pods are being deleted, they are terminated in reverse order, from {N-1..1}
- Before a scaling operation is applied to a Pod, all of its predencessors must be Running and Ready.
- Before a Pod is terminated, all of its successors must be completely shutdown.

## Update Strategies

### On Delete

The `OnDelete` update strategy implements the legacy behavior.