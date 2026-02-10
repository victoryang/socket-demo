# Flannel

[flannel](https://github.com/flannel-io/flannel)

## How it works

Flannel runs a small, single binary agent called `flanneld` on each host, and is responsible for allocating a subnet lease to each host out of a larger, preconfigured address space. Flannel uses either the Kubernetes API or etcd directly to store the network configuration, the allocated subnets, and any auxiliary data (such as the host's public IP). Packets are forward using one of serveral backend mechanisms including VXLAN and various cloud integrations.

## Networking detais

Platforms like Kubernetes assume that each container (pod) has a unique, routable IP inside the cluster. The advantage of this model is that it removes the port mapping complexities that come from sharing a single host IP.

Flannel is responsible for providing a layer 3 IPv4 network between multiple nodes in a cluster. Flannel does not control how containers are networked to the host, only how the traffic is transported between hosts. However, flannel does provide a CNI plugin for Kubernetes and a guidance on integrating with Docker.

Flannel is focused on networking. For network policy, other projects such as Calico can be used.

## Getting started on Kubernetes

The easiest way to deploy flannel with Kubernetes is to use one of several deployment tools and disributions that network clusters with flannel by default.

Though not required, it's recommended that flannel uses the Kubernetes API as its backing store which avoids the need to deploy a discrete `etcd` cluster for `flannel`. This `flannel` mode is known as the *kube subnet manager*.

## Getting started on Docker

flannel is also widely used outside of kubernetes. When deployed outside of kubernetes, etcd is always used as the datastore.

- Y: qnx和linux 分析调研
- T: 
- B: 无