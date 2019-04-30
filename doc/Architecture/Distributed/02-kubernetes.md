# Kubernetes

[k8s官方文档](https://kubernetes.io/docs/home/)

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