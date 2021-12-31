# CRI

https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md

## What is CRI

CRI (Container Runtime Interface) consists of a specifications/requirements (to-be-added), protobuf API, and libraries for container runtimes to integrate with kubelet on a node. The CRI API is currently in Alpha, and the CRI-Docker integration is used by default as of Kubernetes 1.7+.

## Why develop CRI

Prior to the existence of CRI, container runtime (e.g., `docker`,`rkt`) were integrated with kubelet through implementing an internal, high-level interface in kubelet. The entrance barrier for runtimes was high because the integration required understanding the internals of kubelet and contributing to the main Kubernetes repository. More importantly, this would not scale because every new addition incurs a significant maintenance overhead in the main Kubernetes repository.

Kubernetes aims to be extensible. CRI is one small, yet important step to enable pluggable container runtimes and build a healthier ecosystem.

## How to use CRI?

https://kubernetes.io/docs/setup/production-environment/container-runtimes/

CRI is still young and we are actively incorporating feedback from developers to improve the API. Although we strive to maintain backward compatibility, developers should expect occasional API breaking changes.

## Specifications, design documents and proposals

### Motivation

Kubelet employs a declarative pod-level interface, which acts as the sole integration point for container runtimes (e.g.,`docker` and `rkt`). The high-level, declarative interface has caused higher integration and maintenamce cost, and also slowed down feature velocity for the following reasons.

1. **Not every container runtime supports the concept of pods natively**. When integrating with Kubernetes, a significant amount of work needs to go into implementing a shim of significant size to support all pod features. This also adds maintenance overhead (e.g.`docker`)
2. **High-level interface discourage code sharing and reuse among runtimes**. E.g., each runtime today implements an all-encompassing `SyncPod()` function, with the Pod Spec as the input argument. The runtime implements logic to determine how to achieve the desired state based on the current status, (re-)starts pods/containers and manages lifecycle hooks accordingly.
3. **Pod Spec is evolving rapidly**. New features are being added constantly. Any pod-level change or addition requires changing of all container runtime shims. E.g.,init containers and volume containers.

### Goals and Non-Goals

The goals of defining the interface are to
- **imporove extensibility**: Easier container runtime integration
- **improve feature velocity**
- **improve code maintainability**

The non-goals include
- proposing how to integrate with new runtimes, i.g., where the shim resides.
- versioning the new interface/API. We intend to provide API versioning to offer stability for runtime integrations, but the details are beyond the scope of this proposal.
- adding support to Windows containers. Windows container support is a parallel effort.
- re-defining Kubelet's internal interfaces.
- improving Kubelet's efficiency or performance

### Container Runtime Interface

The main idea of this proposal is to adopt an imperative container-level interface, which allows Kubelet to directly control the lifecycles of the containers.

Pod is composed of a group of containers in an isolated environment with resource constraints. In Kubernetes, pod is also the smallest schedulable unit. After a pod has been scheduled to the node, Kubelet will create the environment and the pod as a whole, we will call the pod environment **PodSandbox**.

The container runtime may interpret the PodSandBox concept differently based on how it operates internally. For runtimes relying on hypervisor, sandbox represents a virtual machine naturally. For others, it can be Linux Namespaces.

In short, a PodSandbox should have the following features.

- **Isolation:** E.g., Linux namespaces or a full virtual machine, or even support additional security features.
- **Compute resource specifications:** A PodSandbox should implement pod-level resource demands and restrictions.

*NOTE： The resource specification does not include externalized costs to container setup that are not currently trackable as Pod constraints, e.g., filesystem setup, container image pulling,etc.*

A container in a PodSandbox maps to an application in the Pod Spec. For Linux containers, they are expected to share at least network and IPC namespaces, with sharing more namespaces.

Below is an example of the proposed interfaces.

```go
// PodSandboxManager contains basic operations for sandbox.
type PodSandboxManager interface {
    Create(config *PodSandboxConfig) (string, error)
    Delete(id string) (string, error)
    List(filter PodSandboxFilter) []PodSandboxListItem
    Status(id string) PodSandboxStatus
}

// ContainerRuntime contains basic operations for containers.
type ContainerRuntime interface {
    Create(config *ContainerConfig, sandboxConfig *PodSandboxConfig, PodSandboxID string) (string, error)
    Start(id string) error
    Stop(id string, timeout int) error
    Remove(id string) error
    List(filter ContainerFilter) ([]ContainerListItem, error)
    Status(id string) (ContainerStatus, error)
    Exec(id string, cmd []string, streamOpts StreamOptions) error
}

// ImageService contains image-related operations.
type ImageService interface {
    List() ([]Image, error)
    Pull(image ImageSpec, auth AuthConfig) error
    Remove(image ImageSpec) error
    Status(image ImageSpec) (Image, error)
    Metrics(image ImageSpec) (ImageMetrics, error)
}

type ContainerMetricsGetter interface {
    ContainerMetrics(id string) (ContainerMetrics, error)
}
```

### Pod/Container Lifecycle

The PodSandbox's lifecycle is decoupled from the containers, i.e., a sandbox is created before any containers, and can exist after all containers in it have terminated.

Assume there is a pod with a single container C. To start a pod:
```
create sandbox Foo --> create container C --> start container C
```

To delete a pod:
```
 stop container C --> remove container C --> delete sandbox Foo
```

The container runtime must not apply any transition (such as starting a new container) unless explicitly instructed by Kubelet. It is Kubelet's responsibility to enforce garbage collection, restart policy, and otherwise react to changes in lifecycle.

The only transitions that are possible for a container are described below:
```
() -> Created        // A container can only transition to created from the
                     // empty, nonexistent state. The ContainerRuntime.Create
                     // method causes this transition.
Created -> Running   // The ContainerRuntime.Start method may be applied to a
                     // Created container to move it to Running
Running -> Exited    // The ContainerRuntime.Stop method may be applied to a running 
                     // container to move it to Exited.
                     // A container may also make this transition under its own volition 
Exited -> ()         // An exited container can be moved to the terminal empty
                     // state via a ContainerRuntime.Remove call.
```

Kubelet is also responsible for gracefully terminating all the containers in the sandbox before deleting the sandbox. If Kubelet chooses to delete the sandbox with running containers in it, those containers should be forcibly deleted.

Note that every PodSandbox/container lifecycle operation (create, start, stop, delete) should either return an error or block until the operation succeeds. A successful operation should include a state transition of the PodSandbox/container. E.g. if a `Create` call for a container does not return an error, the container state should be "created" when the runtime is queried.

### Updates to PodSandbox or Containers

Kubernetes support updates only to a very limited set of fields in the Pod Spec. These updates may require containers to be recreated by Kubelet. This can be achieved through the proposed, imperativing container-level interface. On the other hand, PodSandbox update currently is not required.

### Container Lifecycle Hooks

Kubernetes supports post-start and pre-stop lifecycle hooks, with ongoing discussion for supporting pre-start and post-stop hooks in #140.

These lifecycle hooks will be implemented by Kubelet via `Exec` call to the container runtime. This frees the runtimes from having support hooks natively.

Illustration of the container lifecycle and hooks:
```
            pre-start post-start    pre-stop post-stop
               |        |              |       |
              exec     exec           exec    exec
               |        |              |       |
 create --------> start ----------------> stop --------> remove
```

In order for the lifecycle hooks to function as expected, the `Exec` call will need access to the container's lifesystem(e.g. mount namespaces)

### Extensibility

There are several dimensions for container runtime extensibility.

- Host OS (e.g. Linux)
- PodSandbox isolation mechanism (e.g. namespaces or VM)
- PodSandbox OS (e.g. Linux)

As mentioned previously, this proposal will only address the Linux based PodSandbox and containers. All Linux-specific configuration willbe grouped into one field.A container runtime is required to enforce all configuration applicable to its platform, and should return an error otherwise.

## Container Runtime Interface (CRI) Networking Specifications

### Introduction

Container Runtime Interface (CRI) is an ongoing project to allow container runtimes to integrate with kubernetes via a newly-defined API. This document specifies the network requirements for container runtime interface (CRI). CRI networking requirements expand upon kubernetes pod networking requirements. This document does not specify requirements from upper layers of kubernetes network stack, such as Service.

### Requirements

1. Kubelet expects the runtime shim to manage pod's network life cycle. Pod networking should be handled accordingly along with pod sandbox operations.
    - `RunPodSandbox` must set up pod's network. This includes, but is not limited to allocating a pod IP, configuring the pod's network interfaces and default network route. Kubelet expects the pod sandbox to have an IP which is routable within the k8s cluster, if `RunPodSandbox` returns successfully. `RunPodSandbox` must return an error if it fails to set up the pod's network. If the pod's network has already been set up, `RunPodSandbox` must skip network setup and proceed.
    -`StopPodSandbox` must tear down the pod's network. The runtime shim must return error on network tear down failure. If pod's network has already been torn down, `StopPodSandbox` must skip network tear down and proceed.
    - `RemovePodSandbox` may tear down pod's network, if the networking has not been torn down already. `RemovePodSandbox` must return error on network tear down failure.
    - Response from `PodSandboxStatus` must include pod sandbox network status. The runtime shim must return an empty network status if it failed to construct a network status.

2. User supplied pod networking configurations, which are NOT directly exposed by the kubernetes API, should be handled directly by runtime shims. For instance, `hairpin-mode`, `cni-bin-dir`, `cni-conf-dir`, `network-plugin`, `network-plugin-mtu` and `non-masquerade-cidr`. Kubelet will no longer handle these configurations after the transition to CRI is complete.

3. Network configurations that are exposed through the kubernetes API are communicated to the runtime shim through `UpdateRuntimeConfig` interface, e.g. `podCIDR`. For each runtime and network implementation, some configs may not be applicable. The runtime shim may handle or ignore network configuration updates from `UpdateRuntimeConfig` interface.

### Extensibility

- Kubelet is oblivious to how the runtime shim manages networking, i.e. runtime shim is free to use CNI, CNM or any implementation as long as the CRI networking requirements and k8s networking requirements are satisfied.
- Runtime shims have full visibility into pod networking configurations
- As more network feature arrives, CRI will evolve.

## Container Metrics

## 