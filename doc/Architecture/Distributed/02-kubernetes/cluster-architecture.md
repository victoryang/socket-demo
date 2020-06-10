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

The usage of these fields varies depending on your cloud provider or bare metal configuration.

- Hostname: The hostname as reported by the node's kernel
- ExternalIP
- InternalIP

#### Conditions

The ```conditions``` field describe the status of all ```Running``` nodes. Examples of conditions include:

|Node Condition|Description|
|-|-|
|Ready|```True``` if the node is healthy and ready to accept pods, ```false``` if the node is no healthy and is not accepting pods, and ```Unknown``` if the node controller has not heard from the node in the last ```node-monitor-grace-period```(default is 40s)|
|DiskPressure|```True``` if pressure exists on the disk size - that is, if the disk capacity is low; otherwise ```false```|
|MemoryPressure|```True``` if pressure exists on the node memory - that is, if the node memory is low; otherwise ```false```|
|PIDPressure|```True``` if pressure exists on the processes - that is, if there is too many processes; otherwise ```false```|
|NetworkUnavailable|```True``` if the network for the node is not correctly configured, otherwise ```false```|

Noted: If you use command-line tools to print details of a cordoned Node, the Condition includes ```SchedulingDisabled```. ```SchedulingDisabled``` is not a Condition in the Kubernetes API; instead, cordoned nodes are marked Unschedulable in their spec.

If the Status of the Ready condition remains ```Unknown``` or ```False``` for longer than the pod-eviction-timeout (an arguement passed to the kube-controller-manager), all the Pods on the node are scheduled for deteletion by the node controller. The default eviction timeout duration is **five minutes**. In some cases when the node is unreachable, the API server is unable to communicate with kubelet on the node. The decision to delete the pods cannot be communicated to the kubelet until comminucation with the API server is re-established. In the meantime, the pods that are scheduled for deletion may continue to run on the partitioned node.

The node controller does not force delete pods until it is confirmed that they have stopped running in the cluster. You can see the pods that might be running on an unreachable node as being in ```Terminating``` or ```Unknown``` state. In cases where Kubernetes cannot deduce from the underlying infrastructure if a node has permanetly left a cluster, the cluster administrator may need to delete the node object by hand. Deleting the node object from Kubernetes causes all the Pod objects running on the node to be deleted from the API server, and frees up their names.

The node lifecycle controller automatically creates taints that represents conditions. The scheduler takes the Node's taints into consideration when assigning a Pod to a Node. Pods can also have tolerations which let them tolerate a Node's taints.

#### Capacity and Allocatable

Describes the resources available on the node: CPU, memory and the maximum number of Pods that can be scheduled onto the node.

The fields in the capacity block indicate the total amount of resources that a Node has. The allocatable block indicates the amount of resources on a Node that is available to be consumed by normal Pods.

#### Info

Describes general information about the node, such as kernel version, Kubernetes version(kubelet and kube-proxy version), Docker version, and OS name.

#### Node controller

The node controller is a Kubernetes control plane component that messages various aspects of nodes.

The node controller has multiple roles in a node's life.

The first is assigning a CIDR block to the node when it is registered (if CIDR assignment is turned on).

The second is keeping the node controller’s internal list of nodes up to date with the cloud provider’s list of available machines.

The third is monitoring the nodes’ health. 

##### Heartbeats



## Control Plane - Node Communication

This document catalogs the communication path between the control plane(really the apiserver) and the Kubenetes cluster. The intent is to allow users to customize their installation to harden the network configuration such that the cluster can be run on an untrusted network(or on fully public IPs on a cloud provider).

