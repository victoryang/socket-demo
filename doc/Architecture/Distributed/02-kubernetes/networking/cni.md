# CNI

[cni.dev](https://www.cni.dev/docs/)

[CNI Compare](https://www.hwchiu.com/cni-compare.html)

[Flannel](https://www.kubernetes.org.cn/2270.html)

## CNI

CNI (Container Network Interface), a Cloud Native Computing Foundation project, consists of a specification and libraries for writing plugins to configure network interfaces in Linux containers, along with a number of supported plugins. CNI concerns itself only with network connectivity of containers and removing allocated resources when the container is deleted. Because of this focus, CNI has a wide range of support and the specification is simple to implement.

As well as the specification, this repository contains the Go source code of a library for integrating CNI into applications and an example command-line tool for executing CNI plugins. A separate repository contains reference plugins and a template for making new plugins.

## Spec

### Overview

This document proposes a generic plugin-based networking solution for application containers on Linux, the *Container Networking Interface*, or *CNI*

For the purpose of this proposal, we define three terms very specifically:

- *container* is a network isolation domain, though the actual isolation technology is not defined by the specification. This could be network namespace or a virtual machine, for example.
- *network* refers to a group of endpoints that are uniquely addressable that can communicate amongst each other. This could be either an individual container (as specified above), a machine, or some other network device (e.g. a router). Containers can be conceptually *added to* or *removed from* one or more networks.
- *runtime* is the program responsible for executing CNI plugins.
- *plugin* is a program that applies a specified network configuration.

### Summary

The CNI specification defines:

1. A format for administrators to define network configuration.
2. A protocol for container runtimes to make requests to network plugins.
3. A procedure for executing plugins based on a supplied configuration.
4. A procedure for plugins to delegate functionality to other plugins.
5. Data types for plugins to return their results to the runtime.

#### Network configuration format

CNI defines a network configuration format for administrators. It contains directives for both the container runtime as well as the plugins to consume. At plugin execution time, this configuration format is interpreted by the runtime and transformed in to a form to be passed to the plugins.

In general, the network configuration is intended to be static. It can conceptually be thought of as being "on disk", though the CNI specification does not actually require this.

#### Execution Protocol

The CNI protocol is based on execution of binaries invoked by the container runtime. CNI defines the protocol between the plugin binary and the runtime.

A CNI plugin is responsible for configuring a container's network interface in some manner. Plugins fall in to two broad categories:

- "interface" plugins, which create a network interface inside the container and ensure it has connectivity.
- "Chained" plugins, which adjust the configuration of an already-created interface (but may need to create more interfaces to do so).

The runtime passes parameters to the plugin via environment variables and configuration. It supplies configuration via stdin. The plugin returns a result on stdout on success, or an error on stderr if the operation fails. Configuration and results are encoded in JSON.

Parameters invocation-specific settings, whereas configuration is, with some exceptions, the same for any given network.

The runtime must execute the plugin in the runtime's networking domain. (For most cases, this means the root network namespace / dom0)

##### CNI operations

- ADD: add container to network, or apply modifications
    - create the interface defined by CNI_IFNAME inside the container at CNI_NETNS, or 
    - adjust the configuration of the interface defined by CNI_IFNAME inside the container at CNI_NETNS
- DEL: Remove container from network, or un-apply modifications
    - delete the interface defined by CNI_IFNAME inside the container at CNI_NETNS, or
    - undo any modifications applied in the plugin's ADD functionality
- CHECK: Check container's networking is as expected

#### Execution of Network Configurations

This section describe how a container runtime interprets a network configuration 

#### Plugin Delegation

There are some operation that, for whatever reason, cannot reasonably be implemented as a discrete chained plugin. Rather, a CNI plugin may wish to delegate some functionality to another plugin. One common example of this is IP address management.

As part of its operation, a CNI plugin is expected to assign (and maintain) an IP address to the interface and install any necessary routes relevant for that interface. This gives the CNI plugin great flexibility but also places a large burden on it. Many CNI plugins would need to have the same code to support serveral IP management schemes that users may desire (e.g. dhcp, host-local). A CNI plugin may choose to delegate IP management to another plugin.

To lessen the burden and make IP management strategy be orthogonal to the type of CNI plugin, we define a third of plugin - IP address Management Plugin (IPAM plugin), as well as a protocol for plugins to delegate functionality to other plugins.

It is however the responsibility of the CNI plugin, rather than the runtime, to invoke the IPAM plugin at the proper moment in its execution. The IPAM plugin must determine the interface IP/subnet, Gateway and Routes and return this information to the "main" plugin to apply. The IPAM plugin may obtain the information via a protocol(e.g. dhcp), data stored on a local filesystem, the "ipam" section of the Network Configuration file, etc.