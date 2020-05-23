# Pod

[Pod](http://dockone.io/article/9065)

## Pod

A *Pod* is a basic execution unit of Kubernetes application - the smallest and simplest unit in the kubernetes object model that you create or deploy. A Pod represents processors running on your Cluster.

A Pod encapsulates an application's container(or, in some cases, multiple containers), storage resources, a unique network ip, and options govern how the container(s) should run. A Pod represents a unit of deployment: a single instance of an application in kubernetes, which might consist of either a single or a small number of containers that are tightly coupled and that share resources.

Pods in a Kubernetes cluster can be used in two main ways:

- **Pods that run a single container**. The "One-container-per-Pod" model is the most common kubernetes use case; in this case, you can think of a Pod as a wrapper around a single container, and Kubernetes manages Pods rather than the containers directly.

- **Pods that run multiple containers that need to work together**. A Pod might encapsulate an application composed of multiple co-located containers that tightly coupled and need to share resources. These co-located containers might form a single cohesive unit of service-one container serving files from a shared volume to the public, while a separate "sidecar" container refreshes or updates those files. The Pod wraps these containers and storage resources together as a single manageable entity.

Each Pod is meant to run a single instance of a given application. If you want to scale your application horizontally(e.g., run multiple instances), you should use multiple Pods, one for each instance. In Kubernetes, this is generally referred to as *replication*. Replicated Pods are usually created and managed as a group by an abstraction called a Controller.

<img src="pod.svg">

### How Pods manage multiple Containers

Pods are designed to support multiple cooperating processes (as containers) that form a cohesive unit of service. The containers in a Pod are automatically co-located and co-scheduled containers on the same physical or virtual machine in the cluster. The containers can share resources and dependencies, communicate with one another, and coordinate when and how they are terminated.

Note that grouping multiple co-located and co-managed containers in a single Pod is a relatively advanced use case. You should use this pattern only in specific instances in which your containers are tightly coupled. For example, you might have a container that acts as a web server for files in a shared volume, and a seperate "sidecar" container that updates those files from a remote source, as in the upper diagram.

Some Pods have init containers as well as app containers. Init containers run and complete before the app containers are started.

Pods provide two kinds of shared resources for their constituent containers: networking and storage.

#### Networking

Each Pod is assigned a unique IP address. Every container in a Pod shares the network namespace, including the IP address and network ports. Container inside a *Pod* can communicate with one another using `localhost`. When containers in a Pod communicate with entities *outside* the Pod, they must coordinate how they use the shared network resources(such as ports).

#### Storage

A pod can specify a set of shared storage volumes. All containers in a Pod can access the shared volumes, allowing those containers to share data. Volumes also allow persistent data in a Pod to survive in case one of the containers within need to be restarted.

## Working with Pods

You'll rarely create individual Pods directly un Kubernetes-even singleton Pods. When a Pod gets created (directly by you, or indirectly by a controller), it is scheduled to run on a Node in your cluster. The Pod remains on the node until the process is terminated, the Pod object is deleted, the Pod is evicted for lack of resources, or the node fails.

Noted that restarting a container in a Pod should not be confused with restarting a Pod. A Pod is not a process, but an environment for running a container. A pod persists until it is deleted.

Pods do not, by themselves, self-heal. If a Pod is scheduled to a Node that fails, or if the scheduling operation itself fails, the Pod is deleted; likewise, a Pod won't survive an eviction due to a lack of resources or Node maintenance.

Kubernetes uses a higher-level abstraction, called a controller, that handles the work of managing the relatively disposable Pod instances. Thus, while it is possible to use Pod directly, it's far more common in Kubernetes to manage your pods using a controller.

### Pods and controllers

You can use workload resources to create and manage multiple Pods for you. A controller for the resource handles replication and rollout and automatic healing in case of Pod failure. For example, if a Node fails, a controller notices that Pods on that Node have stopped working and create a replacement Pod. The scheduler places the replacement Pod onto a healthy Node.

Here are some examples of workload resources that manage one or more Pods:

- Deployment
- StatefulSet
- DaemonSet

## Pod templates

Controllers for workload resources create Pods from a pod template and manage those Pods on your behalf.

PodTemplates are specifications for creating Pods, and are included in workload resources such as Deployments, Jobs and DaemonSets

Each controller for a workload resource uses the PodTemplate inside the workload object to make actual Pods. The PodTemplate is part of the desired state of whatever workload resource you used to run your app.

Modifying the pod template or switching to a new pod template has no effect on the Pods that already exist. Pods do not receive template updates directly; instead, a new Pod is created to match the revised pod template.

For example, a Deployment controller ensures that the running Pods match the current pod template. If the template is updated, the controller has to remove the existing Pods and create new Pods based on the updated template. Each workload controller implements its own rules for handling changes to the Pod template.

On Nodes, the kubelet does not directly observe or manage any of the details around pod templates and updates; those details are abstracted away. That abstraction and separation of concerns simplifies system semantics, and makes it feasible to extend the cluster’s behavior without changing existing code.

## What is a Pod

Like individual application containers, Pods are considered to be relatively ephemeral (rather than durable) entities. As discussed in pod lifecycle, Pods are created, assigned a unique ID(UID), and scheduled to nodes where they remain until termination(according to restart policy) or deletion. If a node dies, the Pods scheduled to that node are scheduled for deletion, after a time period. A given Pod (as defined by a UID) is not "rescheduled" to a new node.;instead, it can be replaced by an identical Pod, with even the same name if desired, but with a new UID (see replication controller for more details).

When something is said to have the same lifecycle as a Pod, such as a volume, that means that it exists as long as that Pod (with that UID) exists. If the Pod is deleted for any reason, even if an identical replacement is created, the related thing is also destroyed and creat anew.

## Motivation for Pods

### Management

Pods are a model of the pattern of multiple cooperating processes which form a cohesive unit of service. They simplify application deployment and management by providing a high-level abstraction then the set of their constituent applications. Pods serve as unit of deployment, horizontal scaling, and replication. Colocation (co-scheduling), shared fate, coordinated replication, resource sharing, and dependency management are handled automatically for containers in a Pod.

### Resource sharing and communication

Pods enable data sharing and communication among their constituents.

The applications in a Pod all use the same network namespace (same IP and port space), and can thus “find” each other and communicate using localhost. Because of this, applications in a Pod must coordinate their usage of ports. Each Pod has an IP address in a flat shared networking space that has full communication with other physical computers and Pods across the network.

### Use of Pods

Pods can be used to host vertically integrated application stacks (e.g. LAMP), but their primary motivation is to support co-located, co-managed helper programs

Individual Pods are not intended to run multiple instances of the same application, in general.

For a longer explanation, see The Distributed System ToolKit: Patterns for Composite Containers.

### Durability of Pods

Pods aren’t intended to be treated as durable entities. They won’t survive scheduling failures, node failures, or other evictions, such as due to lack of resources, or in the case of node maintenance.

In general, users shouldn’t need to create Pods directly. They should almost always use controllers even for singletons, for example, Deployments. Controllers provide self-healing with a cluster scope, as well as replication and rollout management. Controllers like StatefulSet can also provide support to stateful Pods.

## Termination of Pods

Because Pods represents running processes on nodes in the cluster, it is important to allow those processes to gracefully terminate when they are no longer needed(vs being violently killed with a KILL signal and having no chance to clean up). Users should be able to request deletion and know when processes terminate, but also be able to ensure that deletes eventually complete. When user request deletion of a Pod, the system record intended grace period before the Pod is allowed to be forcefully killed, and a TERM signal is sent to the main process in each container. Once the grace period has expired, the KILL signal is sent to those processes, and the Pod is then deleted from the API server. If the Kubelet or the container manager is restarted while waiting for processes to terminate, the termination will be retried with the full grace period.

By default, all deletes are graceful within 30 seconds. The ```kubectl delete``` command supports the --graceful-period=\<second\> option which allows a user to override the default and specify their own value. The value 0 force deletes the Pod. You must specify an additional flag --force along with --graceful-period=0 in order to perform force deletions.

### Force deletion of Pods

Force deletion of a Pod is defined as deletion of a Pod from the cluster state and etcd immediately. When a force deletion is performed, the API server does not wait the confirmation from the kubelet that Pod has been terminated on the node it was running on. It removes the Pod in the API immediately so a new Pod can be created with the same name. On the node, Pods are set to terminate immediately will still be given a small grace period before being killed.

Force deletion can be potentially dangerous for some Pods and should be performed by caution.

## Privileged mode for pod containers

Any container in a Pod can enable privileged mode, using the privileged flag on the security context of the container spec. This is useful for containers that want to use Linux capabilities like manipulating the network stack and accessing devices. Processes within the container get almost the same privileges that are available to processes outside a container. With privileged mode, it should be easier to write network and volume plugins as separate Pods that don’t need to be compiled into the kubelet.