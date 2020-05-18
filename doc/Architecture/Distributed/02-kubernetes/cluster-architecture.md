# Cluster Architecture

[Nodes](https://kubernetes.io/docs/concepts/architecture/nodes/)

## Nodes

Kubernetes runs your workload by placing containers into Pods to run on *Nodes*. A node may be a virtual or physical machine, depending on the cluster. Each node contains the service necessary to run Pods， managed by the control plane.

Typically you have several nodes in a cluster;in a learning or resource-limited environment,you might have just one. The components on a node include the kubelete, a container runtime, and the kube-proxy.

### Management

There are two main ways to have Nodes added to API server:

1. The kubelete on a node self-register to control plane
2. You, or another human user, manually add a Node object

#### Self-registration of Nodes

When the kubelet flag `--register-node` is true (the default), the kubelet will attempt to register itself with the API server. This is the preferred pattern, used by most distros.

For self-registeration, the kubelet is started with the following options:

- --kubeconfig - Path to credentials to authenticate itself to the API server
- --register-node - Automatically register with the API server
- --node-ip - IP address of the node
- --node-labels - Labels to add when registering node in the cluster
- --node-status-update-frequency - Specifies how often kubelet posts node status to master

When the Node Authorization mode and NodeRestriction admission plugin are enabled, kubelets are only authorized to create/modify their own Node resource.

#### Manual Node administration

You can create or modify Node objects using kubectl.

When you want to create Node objects manually, set the kubelet --register-node to false.

You can modify Node objects regardless of setting of --register-node flag. For example, you can set labels on an existing Node, or mark it unschedulable.

You can use labels on Nodes in conjunction with node selectors on Pods to control scheduling. For example, you can to constrain a Pod to only be eligible to run on a subset of available nodes.

Marking a node as unschedulable prevents the scheduler from placing new pods onto that Node, but does not affect existing Pods on the Node. This is useful as a preparatory step before a node reboot or other maintenance.

To mark a Node unschedulable, run:

kubectl cordon $NODENAME

`
Note: Pods that are part of a DaemonSet tolerate being run on an unschedulable Node. DaemonSets typically provide node-local services that should run on the Node even if it si being drained of workload applications.
`

### Node Status

A Node's status contains the following information:

- Addresses
- Conditions
- Capacity and Allocatable
- Info

You can use kubectl to view a Node's status and other details:

`kubectl decribe node <insert-node-name-here>`

Each section of the output is described below.

#### Addresses

