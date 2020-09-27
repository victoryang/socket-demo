# Kubernetes

[k8s官方文档](https://kubernetes.io/docs/home/)

[kubernetes guide](https://k.i4t.com/)

[kubernetes CN](http://docs.kubernetes.org.cn/)

## Containers
    Containers are small and fast, one applcation can be packed in each container image. The one to one application to image relationship unlocks the full benefits of containers. With containers, immutable container images can be created/release time rather than deployment time, since each application doesn't need to be composed with reset of the application stack, nor married to the production infrastructure environment. Generating container images at build/release time enabls a consistent environment to be carried from development into production.

    - Agile application creation and deployment
    - Continuous development, integration and deployment
    - Dev and Ops seperation of concerns
    - Obervability
    - Environment consistency across development, testing and the production
    - Cloud and OS distribution portability
    - Application-centric management
    - Resource isolation
    - Resource utilization

## Kuernetes
    Kubernetes is an open-source system for automating development, scaling, and management of containerized applications.

    - Master component
    - Node component
    - Addons

### Master Component
    Master Component provides the cluster's control plane. Master component make global decisions about the cluster, and detecting and responding to cluster events.

    Master Components can be run on any machine in the cluster. However, for simplicity, set up scripts typically start all the master components on the same machine, and do not run user containers on this machine.

    - kube-apiserver
    - etcd
    - kube-scheduler
    - kube-controller-manager
    - cloud-controller-manager

### Node Component
    Node Components run on every node, maintaining running pods and providing the Kubenetes runtime environment

    - kubelet
    - kube-proxy
    - container runtime

### Addons
    features

    - DNS
    - WEB UI
    - Container resource monitoring
    - Cluster-level logging

## Kubernetes cluster
### 1. Use MiniKube to create a Cluster
Kubernetes coordinates a highly available cluster of computer that are connected to work as a single unit.

### 2. Deployment
    The deployment instructs Kubernetes how to create and update instances of your application. 
    - Once you've created a Deployment, the Kubertes master schedules methioned application instances onto individual Nodes in the cluster. 
    - Once the applcation instances are created, a Kubernetes Deployment Controller continuously monitors those instances. 
    - If the Node hosting an instance goes down or is deleted, the Deployment controller replaces the instance with an instance on another Node in the cluster. 
    
    This provides a self-healing mechanism to address machine failure or maintainance.

    In a pre-orchestration world, installation scripts would often be used to start applications, but they did not allow recovery from machine failure. By both creating your application instances and keeping them running across Nodes, Kubernetes Deployments provide a fundamentally different approach to application management.

### 3. Kubectl 
    A Kubernetes Pod is a group of one or more Containers, tied together for the purposes of administration and networking.

    A kuberntes Deployment checks on the health of your Pod and restarts the Pod's Container if it terminates.

    Depolyments are the recommended way to manage the creation and scaling of Pods.

#### 3.1 kubectl create deployment
    To create a Deployment that manages a Pod.

#### 3.2 kubectl get deployments
    View the deployments

#### 3.3 kubectl get pods
    View the Pod

#### 3.4 kubectl get events
    View cluster events

#### 3.5 kubectl expose
    To make container accessible from outside the Kubernetes virtual network, you have to expose the Pod as a Kubernetes Service.

#### 3.6 kubectl get services
    View the services

#### 3.7 Addons
    minikube addons list
    To see the currently supported addons

### 4. Explore your App
#### 4.1 Kubernetes Pod
    A Pod is a Kubernetes abstraction that represents a group of one or more application containers, and some shared resources for those containers. Those resources include: 
    - Shared storage, as Volumes
    - Networking, as a unique cluster IP address
    - Information about how to run each container, such as the container image version or specific ports to use

    A Pod models an application-specific "logical host" and can contain different application containers which are relatively tightly coupled.

    Pods are the atomic unit on the kubernetes platform.

#### 4.1.1 Nodes
    A Pod always run on a Node. A Node is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the Master. A Node can have multiple pods, and the Kubernetes master automatically handles scheduling the pods across the Nodes in the cluster. The Master's automatic scheduling takes into account the available resources on each Node.

    Every Kubernetes Node runs at least:
    - Kubelet, aprocess responsible for communication between the Kubernetes Master and the Node; It manages the Pods and the containers running on a machine.
    - A container runtime resiponsible for pulling the container image from a registry, unpacking the container , and running the application

#### 4.2 Node
    kubectl command:
    - kubectl get, list resources
    - kubectl describe, show details information about a resource
    - kubectl logs, print the logs from a container in a pod
    - kubectl exec, execute a command on a container in a pod


### Expose your App Publicly
#### Kubernetes Services
