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

